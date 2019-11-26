package cron

import (
	"database/sql"
	"fmt"
)

type Job struct {
	Category	string
	Status		string
	Count		int
}

//jobs/submitted
func getSubmittedJobCounts(db *sql.DB, startDay string, endDay string)([]Job, error) {
	var jobType *string
	var count int
	//skip curl wrapper jobs
	query := `SELECT b.job_type, count(b.*) AS count
            FROM (
              SELECT a.id,
                CASE WHEN array_length(a.types, 1) = 1 THEN a.types[1]
              ELSE 'Mixed'
           END AS job_type
           FROM (
               SELECT j.id, array_agg(DISTINCT t.name) AS types FROM jobs j
               JOIN job_steps s ON j.id = s.job_id
               JOIN job_types t ON s.job_type_id = t.id
               WHERE j.start_date >= ($2 :: DATE) AND j.start_date <= ($3 :: DATE)
               AND j.app_id != '1e8f719b-0452-4d39-a2f3-8714793ee3e6'
               GROUP BY j.id
           ) AS a
          ) AS b
         GROUP BY b.job_type
         ORDER BY b.job_type;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobs := make([]Job, 3)

	for i := 0; rows.Next(); i++ {
		err := rows.Scan(&jobType, &count)
		if err != nil {
			return nil, err
		}
		jobs[i] = Job{getStringValue(jobType), "Submitted", count}
		output := fmt.Sprintf("Total no.of %[1]v jobs Submitted in last 24 hours: %[2]v", getStringValue(jobType), count)
		fmt.Println(output)

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}
