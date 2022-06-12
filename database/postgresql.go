package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitPostgresql(driver string, source string) *sql.DB {

	var db, err = sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return db
}
