package main

import (
	"fmt"
	"os"
	"strings"
)

func Write() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: wr 'file' 'input'")
		os.Exit(1)
	}

	file := os.Args[1]
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	if _, err := f.WriteString(strings.Join(os.Args[2:], " ")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
