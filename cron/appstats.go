package cron

import (
	"database/sql"
	"fmt"
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
           AND app_id != '1e8f719b-0452-4d39-a2f3-8714793ee3e6'
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
		output := fmt.Sprintf("App name %[1]v App ID %[2]v App Count %[3]v", app.Name, app.ID, app.Count)
		fmt.Println(output)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return apps, nil
}





