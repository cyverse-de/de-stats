package cron

import (
	"database/sql"
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
	rows, err := db.Query(query, amount, startDate, endDate)

	if err != nil {
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

		apps = append(apps, app)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return apps, nil
}





