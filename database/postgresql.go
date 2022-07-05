package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitPostgresql(driver string, source string) (*sql.DB, error) {

	var db, err = sql.Open(driver, source)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
