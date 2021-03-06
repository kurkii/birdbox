package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Copy() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: cp [file] [target]")
		os.Exit(1)
	}

	file := os.Args[1]
	target := os.Args[2]

	bytesRead, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(target, bytesRead, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
