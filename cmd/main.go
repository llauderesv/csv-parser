package main

import (
	"fmt"
	parser "llauderesv/parser/lib"
	"log"
	"os"
)

func main() {
	input := os.Args[1:]
	if len(input) != 1 {
		log.Fatal("Should only accept one argument")
	}

	g, err := parser.ReadCsvFile(input[0])
	if err != nil {
		log.Fatal("Unable to process csv file", err)
	}

	for _, m := range g {
		fmt.Printf("%v", m)
	}
}
