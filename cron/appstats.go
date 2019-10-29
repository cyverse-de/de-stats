package cron

import (
	//"bufio"
	//"bytes"
	"database/sql"
	"fmt"
	//"io"
	//"os"
	//"path/filepath"
	//"strings"
	//"text/tabwriter"
	//"time"

	_ "github.com/lib/pq"
)
type App struct {
	Name	string
	ID		string
	Count	int
}



func GetTopApps(db *sql.DB, amount int, days int) []App{
	var name *string
	var appID string
	var appCount int

	query := `SELECT app_name, app_id, count(*) AS job_count FROM jobs
           WHERE start_date >= (now() - ($2 || ' DAY')::INTERVAL )
           AND app_id != '1e8f719b-0452-4d39-a2f3-8714793ee3e6'
           GROUP BY app_name, app_id
           ORDER BY job_count DESC
           LIMIT $1`
	rows, err := db.Query(query, amount, days)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	apps := make([]App, amount)
	for i := 0; rows.Next(); i++{
		err := rows.Scan(&name, &appID, &appCount)
		if err != nil {
			panic(err)
			break
		}

		apps[i] = App{getStringValue(name), appID, appCount}
		output := fmt.Sprintf("App name %[1]v App ID %[2]v App Count %[3]v", getStringValue(name), appID, appCount)
		fmt.Println(output)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return apps
}

func getStringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}




