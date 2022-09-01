package todo

import (
	"time"
)

type ToDoStatus int

const (
	StatusToDo = iota
	StatusInProgres
	StatusDone
	StatusStopped
)

type Todo struct {
	ID int
	Title string
	Discription string
	CreatedAt time.Time
	Status ToDoStatus
}

var todoList = []Todo{
	{
		ID:		1,
		Title:   "my bootcamp course",
		Discription: "learning go programming",
		CreatedAt: time.Now(),
		Status: StatusToDo,
	},
	{
		ID:		2,
		Title:   "my bootcamp my first app",
		Discription: "learning go programming",
		CreatedAt: time.Now(),
		Status: StatusInProgres,
	},
}

func List() []Todo {
	return todoList
}