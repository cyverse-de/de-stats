package cron

import (
	"database/sql"
)

type LoginCount struct {
	Count int
}

// /logins/distinct
func GetDistinctLoginCount(db *sql.DB, startDate string, endDate string) (LoginCount, error){
	var count int

	query := `SELECT count(distinct user_id) FROM logins WHERE login_time >= ($1 :: DATE )
		AND login_time <= ($2 :: DATE) + INTERVAL '1 day';`

	row := db.QueryRow(query, startDate, endDate)

	var logins LoginCount
	logins = LoginCount{0}
	err := row.Scan(&count)

	if err != nil {
		return LoginCount{-1}, err
	}
	logins.Count = count

	return logins, nil
}

// /login
func GetLoginCount(db *sql.DB, startDate string, endDate string) (LoginCount, error){
	var count int

	query := `select count(user_id) from logins where login_time >= ($1 :: DATE) 
		AND login_time <= ($2 :: DATE) + INTERVAL  '1 day';`

	row := db.QueryRow(query, startDate, endDate)

	var logins LoginCount
	logins = LoginCount{0}
	err := row.Scan(&count)

	if err != nil {
		return LoginCount{-1}, err
	}
	logins.Count = count


	return logins, nil
}
