package main
import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	//_ "github.com/lib/pq"
)

const (
	host	= "trotter.cyverse.org"
	port   	= 6432
	user   	= "de"
	dbname 	= "de"
	folder  = "/tmp"
	logfile = "logs-stdout-output"
)
func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}