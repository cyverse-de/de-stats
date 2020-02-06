package cron

import (
	"database/sql"
	"github.com/cyverse-de/de-stats/logs"
	_ "github.com/lib/pq"
)

type User struct {
	Name string
	Count int
}

func GetTopUsers(db *sql.DB, amount int, startDate string, endDate string) ([]User, error){

	query := `SELECT regexp_replace(u.username, '@.*', '') AS username, count(*) AS count  FROM jobs j
           JOIN users u ON j.user_id = u.id
           WHERE j.start_date >= ($2 :: DATE) AND j.start_date <= ($3 :: DATE) + INTERVAL '1 day'
           GROUP BY u.username
           ORDER BY count DESC
		   LIMIT $1;`
	logs.Logger.Debug("Parameters passed to /users: amount - %s, startDate - %s, endDate - %s", amount, startDate, endDate)
	rows, err := db.Query(query, amount, startDate, endDate)

	if err != nil {
		logs.Logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next(){
		var user User
		err := rows.Scan(&user.Name, &user.Count)
		if err != nil {
			return nil, err
		}
		logs.Logger.Debug("username: %s, count: %s", user.Name, user.Count)
		users = append(users, user)

	}

	err = rows.Err()
	if err != nil {
		logs.Logger.Error(err)
		return nil, err
	}

	return users, nil

}