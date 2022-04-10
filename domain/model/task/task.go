package task

import (
	"database/sql"
	"time"
)

type Task struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}

func NewTask(id int, name string) Task {
	return Task{Id: id, Name: name}
}
