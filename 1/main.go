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
	totalCost := 0
	// Print the lines
	for _, line := range lines {
		// count bytes in line
		byteCount := len(line)
		characterCount := len([]rune(line))
		canTweet := characterCount <= 140
		canSMS := byteCount <= 160
		cost := 0
		if canTweet && canSMS {
			cost = 13
		} else if canTweet {
			cost = 7
		} else if canSMS {
			cost = 11
		}
		totalCost += cost
	}
	fmt.Println("Total cost:", totalCost)
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
