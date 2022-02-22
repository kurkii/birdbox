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
	}
}
