package cron

import (
	"database/sql"

	"github.com/cyverse-de/dbutil"

)

const (
	host	= "trotter.cyverse.org"
	port   	= 6432
	user   	= "de"
	dbname 	= "de"
	folder  = "/tmp"
	logfile = "logs-stdout-output"
)


func InitDB(dbURI string) *sql.DB{
	connector, err := dbutil.NewDefaultConnector("1m")
	if err != nil {
		panic(err)
	}

	db, err := connector.Connect("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func getStringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}