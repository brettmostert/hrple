package data

import "time"

type Activity struct {
	Id           string
	Name         string
	Description  string
	Type         string // later to be a type - behaviour, task,
	IsArchived   bool
	IsDeleted    bool
	CreatedTime  time.Time
	ModifiedTime time.Time
}
