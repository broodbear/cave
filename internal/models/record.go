package models

import "time"

type Record struct {
	ID        int
	Project   string
	Target    string
	Username  string
	Password  string
	CreatedAt time.Time
}
