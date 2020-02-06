package cron

import (
	"database/sql"
	"github.com/cyverse-de/de-stats/logs"
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

	logs.Logger.Debug("Parameters passed to /logins: startDate - %s, endDate - %s", startDate, endDate)
	row := db.QueryRow(query, startDate, endDate)

	var logins LoginCount
	logins = LoginCount{0, 0}
	err := row.Scan(&count, &distinctCount)

	if err != nil {
		logs.Logger.Error(err)
		return LoginCount{-1, -1}, err
	}
	logins.Count = count
	logins.DistinctCount = distinctCount
	logs.Logger.Debug("Login count: %s, distinct count: %s", count, distinctCount)
	return logins, nil
}
