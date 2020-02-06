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


func InitDB(dbURI string) (*sql.DB, error){
	connector, err := dbutil.NewDefaultConnector("1m")
	if err != nil {
		return nil, err
	}

	db, err := connector.Connect("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getStringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}