package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	poo := rune('💩')
	pooCount := 0
	width := len([]rune(lines[0]))
	height := len(lines)
	startX, startY := 0, 0
	for {
		if ([]rune(lines[startY])[startX]) == poo {
			pooCount++
		}
		// walk down 1
		startY++
		if startY >= height {
			break
		}
		// walk right 2
		startX++
		if startX >= width {
			startX = 0
		}
		startX++
		if startX >= width {
			startX = 0
		}
	}
	fmt.Println(pooCount)
}

// readLines reads a file and returns the lines as a slice of strings.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
