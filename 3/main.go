package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	goodPasswords := 0
	for _, line := range lines {
		numBytes := len([]rune(line))
		oneDigit := false
		oneUpper := false
		oneLower := false
		oneEightBit := false
		goodLength := numBytes >= 4 && numBytes <= 12
		for _, c := range []rune(line) {
			if unicode.IsDigit(c) {
				oneDigit = true
			}
			if unicode.IsUpper(c) {
				oneUpper = true
			}
			if unicode.IsLower(c) {
				oneLower = true
			}
			if c > 127 {
				oneEightBit = true
			}
		}
		if goodLength && oneDigit && oneUpper && oneLower && oneEightBit {
			goodPasswords++
		}
	}
	fmt.Println("Total good passwords: ", goodPasswords)

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
