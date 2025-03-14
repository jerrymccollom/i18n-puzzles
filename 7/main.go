package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines, err := readLines("test-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	halifax, _ := time.LoadLocation("America/Halifax")
	santiago, _ := time.LoadLocation("America/Santiago")

	sum := 0
	for lineNumber, line := range lines {
		s := strings.Split(line, "\t")
		ts, err := time.Parse(time.RFC3339, s[0])
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		// Timestamp in the timezone
		hts := ts.In(halifax)
		sts := ts.In(santiago)

		// Find the UTC offset in hours for both hts and sts
		_, originalOffset := ts.Zone()
		originalOffset /= 3600
		_, htsOffset := hts.Zone()
		htsOffset /= 3600
		_, stsOffset := sts.Zone()
		stsOffset /= 3600

		correctMinutes, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println("Error parsing correctMinutes:", err)
			return
		}
		incorrectMinutes, err := strconv.Atoi(s[2])
		if err != nil {
			fmt.Println("Error parsing incorrectMinutes:", err)
			return
		}

		calcTs := hts
		if originalOffset == htsOffset && originalOffset == stsOffset {
			fmt.Println(lineNumber+1, "Ambiguous")
		} else if originalOffset == htsOffset {
			fmt.Println(lineNumber+1, "Halifax")
		} else if originalOffset == stsOffset {
			fmt.Println(lineNumber+1, "Santiago")
			calcTs = sts
		} else {
			fmt.Println(lineNumber+1, "UNKNOWN")
		}
		calcTs = calcTs.Add(time.Duration(-incorrectMinutes) * time.Minute).Add(time.Duration(correctMinutes) * time.Minute)
		fmt.Println(calcTs.Format(time.RFC3339))
		sum += calcTs.Hour() * (lineNumber + 1)

	}
	fmt.Println(sum)

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
