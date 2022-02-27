package main

import (
	"fmt"
	"os"
)

func Bird() {

	file := os.Args[1]

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(content))
}
