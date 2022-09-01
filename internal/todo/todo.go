package todo

import (
	"time"
)

const (
	StatusToDo = iota
	StatusInProgres
	StatusDone
	StatusStopped
)

type ToDo struct {
	ID int
	Title string
	Discription string
	CreatedAt time.Time
	Status ToDoStatus
}