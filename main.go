package main

import (
	"os"
	"path/filepath"
)

func main() {
	baseName := filepath.Base(os.Args[0])
	switch baseName {
	case "cp":
		Copy()
	case "mv":
		Move()
	case "echo":
		Echo()
	case "wr":
		Write()
	case "bird":
		Bird()
	case "wc":
		WordCount()
	case "mkdir":
		Mkdir()
	case "cr":
		Create()
	case "ls":
		List()
	case "rm":
		Remove()
	case "exit":
		Remove()
	}
}
