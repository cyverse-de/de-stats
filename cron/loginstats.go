package cron

import (
	"database/sql"
)

type LoginCount struct {
	Count int
	DistinctCount int
}

// /logins
func GetLoginCount(db *sql.DB, startDate string, endDate string) (LoginCount, error){
	var count int
	var distinctCount int
	query := `select count(user_id), count(distinct user_id) from logins where login_time >= ($1 :: DATE) 
		AND login_time <= ($2 :: DATE) + INTERVAL  '1 day';`

	row := db.QueryRow(query, startDate, endDate)

	var logins LoginCount
	logins = LoginCount{0, 0}
	err := row.Scan(&count, &distinctCount)

	if err != nil {
		return LoginCount{-1, -1}, err
	}
	logins.Count = count
	logins.DistinctCount = distinctCount
	return logins, nil
}
