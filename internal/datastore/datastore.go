package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatastore(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return db, err
	}

	return db, nil
}
