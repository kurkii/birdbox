package main

import (
	"fmt"
	"os"
)

func Mkdir() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: mkdir [directory]")
		os.Exit(1)
	}

	file := os.Args[1]

	err := os.Mkdir(file, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
