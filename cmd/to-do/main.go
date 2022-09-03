package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
	ID          int
	Title       string
	Discription string
	CreatedAt   time.Time
	Status      ToDoStatus
}

var (
	todoList     = make([]Todo, 0)
	lastID   int = 0
)

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

func Create(title, discription string) Todo {
	lastID++
	newTodo := Todo{
		ID:          lastID,
		Title:       title,
		Discription: discription,
		CreatedAt:   time.Now(),
		Status:      StatusToDo,
	}
	todoList = append(todoList, newTodo)

	return newTodo
}

func Save() error {
	f, err := os.OpenFile("./database.csv", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, todo := range todoList {
		todoStr := fmt.Sprintf("%d %s;%s;%s;%d\n", todo.ID, todo.Title, todo.Discription, todo.CreatedAt.Format(time.RFC3339), todo.Status)

		_, err := f.WriteString(todoStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile() error {
	f, err := os.OpenFile("./database.csv", os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func main() {
	lenOsArgs := len(os.Args)
	if lenOsArgs < 2 {
		fmt.Println("not enouph argumments")
		os.Exit(1)
	}
	defer Save()

	switch os.Args[1] {
	case "list":
		todos := List()
		for i := range todos {
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
	case "create":
		var title, discription string

		if lenOsArgs < 3 {
			fmt.Println("can not argument create")
			os.Exit(1)
		}
		title = os.Args[2]
		if lenOsArgs == 4 {
			discription = os.Args[3]
		}

		createdTodo := Create(title, discription)
		fmt.Printf("[%d] - %s %s %d\n", createdTodo.ID, createdTodo.Title, createdTodo.Discription, createdTodo.Status)
	default:
		fmt.Println("invalid command argument")
		os.Exit(1)
	}
}
