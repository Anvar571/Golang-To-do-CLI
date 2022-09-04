package main

import (
	"fmt"
	"projects/todo"
	"os"
)

func main() {
	lenOsArgs := len(os.Args)
	if lenOsArgs < 2 {
		fmt.Println("can not argumets")
		os.Exit(1)
	}
	fmt.Println(todo.List())
	fmt.Println(os.Args[0])
}