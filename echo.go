package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {

	if len(os.Args) < 2 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: echo [text]")
		os.Exit(1)
	}

	fmt.Println(strings.Join(os.Args[1:], " "))
}
