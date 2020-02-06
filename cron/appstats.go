package cron

import (
	"database/sql"
	"github.com/cyverse-de/de-stats/logs"
	_ "github.com/lib/pq"
)
type App struct {
	Name	string
	ID		string
	Count	int
}

func GetTopApps(db *sql.DB, amount int, startDate string, endDate string) ([]App, error){
	query := `SELECT app_name, app_id, count(*) AS job_count FROM jobs
           WHERE start_date >= ($2 :: DATE)
           AND start_date <= ($3 :: DATE) + INTERVAL '1 day'
           AND app_id NOT IN (
           		SELECT app_steps.app_id::TEXT FROM app_steps
    			JOIN tasks ON app_steps.task_id = tasks.id
   				JOIN tools ON tasks.tool_id = tools.id
    			JOIN tool_types ON tools.tool_type_id = tool_types.id
    			WHERE tool_types.name = 'internal'
           )
           GROUP BY app_name, app_id
           ORDER BY job_count DESC
           LIMIT $1`
	logs.Logger.Debug("Parameters passed to /apps: amount - %s, startDate - %s, endDate - %s", amount, startDate, endDate)
	rows, err := db.Query(query, amount, startDate, endDate)
	if err != nil {
		logs.Logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	var apps []App
	for rows.Next(){
		var app App
		err := rows.Scan(&app.Name, &app.ID, &app.Count)
		if err != nil {
			return nil, err
		}
		logs.Logger.Debug("App name: %s, ID: %s, count: %s", app.Name, app.ID, app.Count)
		apps = append(apps, app)

	}
	err = rows.Err()
	if err != nil {
		logs.Logger.Error(err)
		return nil, err
	}

	return apps, nil
}





