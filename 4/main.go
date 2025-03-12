package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	depart := time.Time{}
	arrive := time.Time{}
	timeTotal := time.Duration(0)

	for _, line := range lines {
		colon := strings.Index(line, ":")
		if colon > 0 {
			action := line[:strings.Index(line, ":")]
			tz := strings.Trim(line[colon+1:], " \t")
			idx := IndexOfFirstWhitespace(tz)
			date := strings.Trim(tz[idx+1:], " \t")
			tz = tz[:idx]
			loc, err := time.LoadLocation(tz)
			if err != nil {
				fmt.Println("Error parsing location:", err)
				return
			}
			ts, err := time.ParseInLocation("Jan 02, 2006, 15:04", date, loc)
			if err != nil {
				fmt.Println("Error parsing time:", err)
				return
			}

			if action == "Departure" {
				depart = ts
			} else if action == "Arrival" {
				arrive = ts
				timeTotal += arrive.Sub(depart)
			}
		}
	}
	fmt.Println(timeTotal.Minutes())
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

func IndexOfFirstWhitespace(s string) int {
	return strings.IndexFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n' || r == '\r'
	})
}
