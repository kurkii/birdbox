package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func count(value string) int {

	re := regexp.MustCompile(`[\S]+`)

	results := re.FindAllString(value, -1)
	return len(results)
}

func WordCount() {

	arg := os.Args[1]

	switch arg {

	case "-w":
		fmt.Println(count((strings.Join(os.Args[1:], " "))))
	default:
		fmt.Println("Invalid operation!")
		fmt.Println("Syntax: wc [-w] [message]")
		os.Exit(1)
	}

}
