package main

import (
	"fmt"
	"os"
)

func Remove() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: rm [file]")
		os.Exit(1)
	}

	file := os.Args[1]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
