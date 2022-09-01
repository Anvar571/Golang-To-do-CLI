package main

import (
	"fmt"
	"os"
	"./internal/todo/"
)

func main() {
	lenOsArgs:= len(os.Args)
	if lenOsArgs < 2 {
		fmt.Println("not enouph argumments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		todo: todo.List()
	case "get":

	}
}
