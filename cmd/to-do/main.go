package main

import (
	"flag"
	"fmt"
	// "os"
)

func main() {
	var (
		cmdList   string
		// cmdGet    int
		// cmdCreate string
		// cmdEdit   string
		// cmdStatus string
	)
	flag.StringVar(&cmdList, "list", "", "return all todos")
	flag.Parse()

	fmt.Println(cmdList)
}
