package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
	"unicode"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// latin := make(map[string]string)
	utf := make(map[string]int)
	lineNumber := 0
	for {
		line := lines[lineNumber]
		lineNumber++
		if len(line) == 0 {
			break
		}
		coded := ContainsGarbledCharacters(line)
		if !coded {
			utf[line] = lineNumber
			fmt.Println("NONE  ", lineNumber, line)
		} else {
			s, _ := UTF8ToLatin1(line)
			if !ContainsGarbledCharacters(s) {
				utf[s] = lineNumber
				fmt.Println("SINGLE", lineNumber, s)
			} else {
				s, _ = UTF8ToLatin1(s)
				utf[s] = lineNumber
				fmt.Println("DOUBLE", lineNumber, s)
			}
		}
	}

	// Now get the patterns
	total := 0
	for {
		line := lines[lineNumber]
		length, letter, position := GetPattern(line)
		match := false
		for k, v := range utf {
			if len([]rune(k)) == length {
				if []rune(k)[position] == letter {
					fmt.Println(line, k, v)
					total += v
					match = true
					break
				}
			}
		}
		if !match {
			fmt.Println(line, "NO MATCH")
		}
		lineNumber++
		if lineNumber >= len(lines) {
			break
		}
	}

	fmt.Println("Total:", total)
}

func GetPattern(line string) (int, rune, int) {
	pos := 0
	length := 0
	matchChar := rune(0)
	matchPos := 0
	for _, char := range line {
		if char == ' ' {
			continue
		}
		if char == '.' {
			pos++
			length++
			continue
		}
		matchChar = char
		matchPos = pos
		length++
	}
	return length, matchChar, matchPos
}

// ContainsGarbledCharacters checks if a UTF-8 string contains any Latin-1 characters.
func ContainsGarbledCharacters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func UTF8ToLatin1(utf8String string) (latin1String string, err error) {
	latin1Bytes, _, err := transform.String(charmap.ISO8859_1.NewEncoder(), utf8String)
	if err != nil {
		return "", err
	}
	return latin1Bytes, nil
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
