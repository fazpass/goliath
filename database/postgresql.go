package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitPostgresql(driver string, source string) (*sql.DB, error) {

	var db, err = sql.Open(driver, source)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, nil
	}

	return db, nil
}
