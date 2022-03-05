package main

import (
	"fmt"
	"os"
)

func Move() {
	if len(os.Args) < 3 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: mv [file] [target]")
		os.Exit(1)
	}

	file := os.Args[1]
	target := os.Args[2]

	err := os.Rename(file, target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
