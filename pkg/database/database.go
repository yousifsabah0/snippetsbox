package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	return db, err
}
