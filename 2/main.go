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
	// timestamps is a map where key is a timestamp and value is a count of how many times that timestamp appears in the file
	timestamps := make(map[time.Time]int)
	lineCounts := make(map[time.Time]map[int]bool)

	lineNumber := 0
	for _, line := range lines {
		lineNumber++
		// parse line as a timestamp that is in the format "2019-06-05T08:15:00-04:00"
		timestamp, err := time.Parse(time.RFC3339, line)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}
		// add 1 to the count of the timestamp
		timestamps[timestamp.UTC()]++
		m, ok := lineCounts[timestamp.UTC()]
		if !ok {
			m = make(map[int]bool)
			lineCounts[timestamp.UTC()] = m
		}
		m[lineNumber] = true
	}
	for k, v := range timestamps {
		if v >= 4 {
			ts := k.Format(time.RFC3339)
			ts = strings.Replace(ts, "Z", "+00:00", 1)
			fmt.Println("Timestamp", ts, "appears", v, "times")
			for line, _ := range lineCounts[k] {
				fmt.Println("Line", line)
			}
		}
	}
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
