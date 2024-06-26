package datastore

import (
	"os"
)

const createCave = `create table
  'cave' (
    'id' integer not null primary key autoincrement,
    'created_at' datetime not null default CURRENT_TIMESTAMP,
    'project' varchar(255) null,
    'target' varchar(255) null,
    'username' varchar(255) null,
    'password' varchar(255) null,
    unique ('id')
  );`

func Migrate(path, database string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsExist(err) {
		return nil
	}

	os.MkdirAll(path, 0700)

	db, err := NewDatastore(path + "/" + database)
	if err != nil {
		return err
	}

	db.Exec(createCave)

	return nil
}
