package todo

import (
	"time"
)

type Todo struct {
	ID int
	Title string
	Description string
	CreatedAt time.Time
}

var (
	todoList = make([]Todo, 0)
	idCount = 0
)

func List() Todo {
	return Todo{}
}

func Create(title, description string) Todo {
	idCount++
	newTodo := Todo{
		ID: idCount,
		Title: title,
		Description: description,
		CreatedAt: time.Now(),
	}
	todoList = append(todoList, newTodo)
	
	return newTodo
}