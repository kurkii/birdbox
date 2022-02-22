package main

import (
	"log"
	"os"
	"strings"
)

func Write() {
	file := os.Args[1]
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(strings.Join(os.Args[1:], " ")); err != nil {
		log.Println(err)
	}

}
