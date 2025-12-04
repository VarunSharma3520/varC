package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile reads a file line by line and returns a slice of lines.
// Returns an error if the file cannot be opened or read.
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %w", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %q: %w", path, err)
	}

	return lines, nil
}
