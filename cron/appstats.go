package cron
import (
	//"bufio"
	//"bytes"
	"database/sql"
	"fmt"
	//"io"
	//"os"
	//"path/filepath"
	//"strings"
	//"text/tabwriter"
	//"time"

	_ "github.com/lib/pq"
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

func getAllApps(db *sql.DB){
	var name *string
	query := "SELECT name FROM apps where deleted = false"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		output := fmt.Sprintf("App name %[1]v", getStringValue(name))
		fmt.Println(output)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func getStringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}


func main(){
	db := InitDB()
	defer db.Close()
	getAllApps(db)
}

