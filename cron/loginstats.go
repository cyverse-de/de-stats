package cron

import "database/sql"

type Login struct {
	Count int
}

// /logins/distinct
func GetDistinctLoginCount(db *sql.DB, startDate string, endDate string) (Login, error){
	var count int

	query := `SELECT count(distinct user_id) FROM logins WHERE login_time >= ($1 :: DATE )
		AND login_time <= ($2 :: DATE);`

	rows, err := db.Query(query, startDate, endDate)

	if err != nil {
		return Login{-1}, err
	}

	defer rows.Close()
	var logins Login
	logins = Login{0}
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return Login{-1}, err
		}
		logins.Count += count
	}

	err = rows.Err()
	if err != nil {
		return Login{-1}, err
	}

	return logins, nil
}

// /login
func GetLoginCount(db *sql.DB, startDate string, endDate string) (Login, error){
	var count int

	query := `select count(user_id) from logins where login_time >= ($1 :: DATE) 
		AND login_time <= ($2 :: DATE);`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return Login{-1}, err
	}

	defer rows.Close()
	var logins Login
	logins = Login{0}
	for rows.Next(){
		err := rows.Scan(&count)
		if err != nil {
			return Login{-1}, err
		}
		logins.Count += count
	}

	err = rows.Err()
	if err != nil {
		return Login{-1}, err
	}

	return logins, nil
}
