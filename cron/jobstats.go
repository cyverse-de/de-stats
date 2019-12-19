package cron

import (
	"database/sql"
	"fmt"
)

type JobStats struct {
	Category	string
	Status		string
	Count		int
}

//jobs/counts
func GetJobCounts(db *sql.DB, startDate string, endDate string)([]JobStats, error){

	query := `WITH internal_app_ids AS (
				SELECT app_steps.app_id::TEXT FROM app_steps
				JOIN tasks ON app_steps.task_id = tasks.id
				JOIN tools ON tasks.tool_id = tools.id
				JOIN tool_types ON tools.tool_type_id = tool_types.id
				WHERE tool_types.name = 'internal'
			)
			SELECT c.job_type, c.status, c.count FROM (
				SELECT b.job_type, 'Submitted' AS status, count(b.*) AS count
				FROM (
					SELECT
						a.id,
						CASE
							WHEN array_length(a.types, 1) = 1 THEN a.types[1]
							ELSE 'Mixed'
						END AS job_type
					FROM (
						SELECT j.id, array_agg(DISTINCT t.name) AS types FROM jobs j
						JOIN job_steps s ON j.id = s.job_id
						JOIN job_types t ON s.job_type_id = t.id
						WHERE j.start_date >= ($1 :: DATE) AND j.start_date <= ($2 :: DATE) + INTERVAL '1 day'
						AND j.app_id NOT IN (SELECT app_id FROM internal_app_ids)
						GROUP BY j.id
					) AS a
				) AS b
				GROUP BY b.job_type
				UNION
				SELECT b.job_type, b.status, count(b.*) AS count
				FROM (
					SELECT
						a.id,a.status,
						CASE
							WHEN array_length(a.types, 1) = 1 THEN a.types[1]
							ELSE 'Mixed'
						END AS job_type
					FROM (
						SELECT j.id, j.status, array_agg(t.name) AS types FROM jobs j
						JOIN job_steps s ON j.id = s.job_id
						JOIN job_types t ON s.job_type_id = t.id
						WHERE j.end_date >= ($1 :: DATE) AND j.end_date <= ($2 :: DATE) + INTERVAL '1 day'
						AND j.app_id NOT IN (SELECT app_id FROM internal_app_ids)
						AND j.status in ('Completed', 'Failed', 'Canceled')
						GROUP BY j.id
					) AS a
				) AS b
				GROUP BY b.job_type, b.status
			) AS c
			ORDER BY c.job_type, c.status`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err;
	}

	defer rows.Close()

	var jobs []JobStats

	for rows.Next() {
		var job JobStats
		err := rows.Scan(&job.Category, &job.Status, &job.Count)
		if err != nil {
			return nil, err;
		}
		jobs = append(jobs, job)
		output := fmt.Sprintf("Total no.of %[1]v jobs %[2]v: %[3]v", job.Category, job.Status, job.Count)
		fmt.Println(output)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return jobs, nil


}