package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func List() {

	if len(os.Args) == 1 {
		mydir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		files, err := ioutil.ReadDir(mydir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}
		os.Exit(0)
	}

	dir := os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	os.Exit(0)
}
