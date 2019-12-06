package cron
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	Name string
	Count int
}

func GetTopUsers(db *sql.DB, amount int, startDate string, endDate string) ([]User, error){
	var username *string
	var count int

	query := `SELECT regexp_replace(u.username, '@.*', '') AS username, count(*) AS count  FROM jobs j
           JOIN users u ON j.user_id = u.id
           WHERE j.start_date >= ($2 :: DATE) AND j.start_date <= ($3 :: DATE) + INTERVAL '1 day'
           GROUP BY u.username
           ORDER BY count DESC
		   LIMIT $1;`

	rows, err := db.Query(query, amount, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for i := 0; rows.Next(); i++{
		err := rows.Scan(&username, &count)
		if err != nil {
			return nil, err
		}
		users = append(users, User{getStringValue(username), count})
		output := fmt.Sprintf("Username %[1]v Count %[2]v", getStringValue(username), count)
		fmt.Println(output)

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil

}