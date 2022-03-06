package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Create() {

	if len(os.Args) != 1 {
		fmt.Println("Invalid Operation!")
		fmt.Println("Syntax: cr [file]")
		os.Exit(1)
	}

	target := os.Args[1]

	err := ioutil.WriteFile(target, nil, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
