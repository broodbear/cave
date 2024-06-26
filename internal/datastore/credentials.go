package datastore

import (
	"database/sql"

	"github.com/broodbear/cave/internal/models"
)

const (
	insertCredentialQuery  = "insert into cave (project, target, username, password) values ($1, $2, $3, $4)"
	selectCredentialsQuery = "select project, target, username, password from cave"
)

type Credentials struct {
	db *sql.DB
}

func NewCredentials(db *sql.DB) Credentials {
	return Credentials{
		db: db,
	}
}

func (c Credentials) Save(project, target, username, password string) error {
	_, err := c.db.Exec(insertCredentialQuery, project, target, username, password)
	if err != nil {
		return err
	}

	return nil
}

func (c Credentials) All() ([]models.Record, error) {
	records := []models.Record{}

	rows, err := c.db.Query(selectCredentialsQuery)
	if err != nil {
		return records, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.Record

		err = rows.Scan(
			&record.Project,
			&record.Target,
			&record.Username,
			&record.Password,
		)
		if err != nil {
			return records, err
		}

		records = append(records, record)
	}

	return records, nil
}
