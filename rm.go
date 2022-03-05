package main

import (
	"fmt"
	"os"
)

func Remove() {
	file := os.Args[1]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
