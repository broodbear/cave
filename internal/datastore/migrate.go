package datastore

import (
	"os"
)

const createCave = `create table
  'credentials' (
    'id' integer not null primary key autoincrement,
    'created_at' datetime not null default CURRENT_TIMESTAMP,
    'project' varchar(255) null,
    'target' varchar(255) null,
    'username' varchar(255) null,
    'password' varchar(255) null,
    unique ('id')
  );
create table
  'nmap' (
    'id' integer not null primary key autoincrement,
    'created_at' datetime not null default CURRENT_TIMESTAMP,
    'project' varchar(255) null,
    'target' varchar(255) null,
    'results' TEXT null,
    unique ('id')
  )`

func Migrate(path, database string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsExist(err) {
		return nil
	}

	err = os.MkdirAll(path, 0o700)
	if err != nil {
		return err
	}

	db, err := NewDatastore(path + "/" + database)
	if err != nil {
		return err
	}

	_, err = db.Exec(createCave)
	if err != nil {
		return err
	}

	return nil
}
