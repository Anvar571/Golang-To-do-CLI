package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"errors"
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

func Get(id int) (Todo, error) {
	for _, todo := range todoList {
		if todo.ID == id {
			return todo, nil
		}
	}
	var error = errors.New("todo not found ")
	return Todo{}, error
}

func main() {
	lenOsArgs:= len(os.Args)
	if lenOsArgs < 2 {
		fmt.Println("not enouph argumments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		todos := List()
		for i:= range todos {
			fmt.Printf("[%d] - %s %s %d\n", todos[i].ID, todos[i].Title, todos[i].Discription, todos[i].Status)
		}
	case "get":
		if lenOsArgs < 3 {
			fmt.Println("is not argument")
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, _ := strconv.ParseInt(idStr, 10, 32)

		FoundTodo, err := Get(int(id))
		if err != nil {
			fmt.Printf("can not todo get: %v", err)
			os.Exit(1)
		}
		fmt.Printf("[%d] - %s %s %d\n", FoundTodo.ID, FoundTodo.Title, FoundTodo.Discription, FoundTodo.Status)
	default:
		fmt.Println("invalid command argument")
		os.Exit(1)
	}
}
