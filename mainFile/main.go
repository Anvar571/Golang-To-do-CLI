package main

import (
	"To-Do-CLI/todos"
	"fmt"
	"os"
)

func main() {
	lenOsArgs := len(os.Args)
	if lenOsArgs < 2 {
		fmt.Println("can not argumets")
		os.Exit(1)
	}
	fmt.Println(todos.List())
	fmt.Println(os.Args[0])
}
