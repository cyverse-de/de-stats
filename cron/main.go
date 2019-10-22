package cron

import (
	"database/sql"
	"fmt"
)

const (
	host	= "trotter.cyverse.org"
	port   	= 6432
	user   	= "de"
	dbname 	= "de"
	folder  = "/tmp"
	logfile = "logs-stdout-output"
)


func InitDB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}


func main(){
	db := InitDB()
	defer db.Close()
	amount := 10
	days := 100
	GetTopApps(db, amount, days)
}