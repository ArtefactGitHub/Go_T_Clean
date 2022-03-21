package model

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewTask(id int, name string) Task {
	return Task{Id: id, Name: name}
}
