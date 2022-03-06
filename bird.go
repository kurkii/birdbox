package main

import (
	"fmt"
	"os"
)

func Bird() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: bird [file]")
		os.Exit(1)
	}

	file := os.Args[1]

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(content))
}
