package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/VarunSharma3520/varC/internal/lexer"
	"github.com/VarunSharma3520/varC/internal/utils"
)

func main() {
	// Define a command-line flag for input file
	filePath := flag.String("file", "", "Path to the input varC file")
	flag.Parse()

	// Check if the flag was provided
	if *filePath == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s --file <filename>\n", os.Args[0])
		os.Exit(1)
	}

	// Read file
	lines, err := utils.ReadFile(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %q: %v\n", *filePath, err)
		os.Exit(1)
	}

	// Join lines into a single string
	source := strings.Join(lines, "\n")

	// Tokenize entire file
	jsonTokens, err := lexer.FlattenTokens(source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lexer error: %v\n", err)
		os.Exit(1)
	}

	// Print JSON output
	fmt.Println(jsonTokens)
}
