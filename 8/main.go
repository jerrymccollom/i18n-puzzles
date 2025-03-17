package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/unicode/norm"
	"os"
	"regexp"
	"strings"
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
		oneVowel := false
		oneConsonant := false
		goodLength := numBytes >= 4 && numBytes <= 12
		repeated := false
		repeat := make(map[rune]bool)
		runeLine := []rune(strings.ToLower(line))
		for _, c := range runeLine {
			if unicode.IsDigit(c) {
				oneDigit = true
			}
			if IsVowel(c) {
				oneVowel = true
			}
			if IsConsonant(c) {
				oneConsonant = true
			}
			normalizedRune := []rune(string(norm.NFD.Bytes([]byte(string(c)))))
			if repeat[normalizedRune[0]] {
				repeated = true
			} else {
				repeat[normalizedRune[0]] = true
			}
		}
		if goodLength && oneDigit && oneVowel && oneConsonant && !repeated {
			goodPasswords++
			// fmt.Println(" OK: ", line)
		} else {
			//fmt.Println("Bad: ", line)
			//if !goodLength {
			//	fmt.Println("   - Bad length")
			//}
			//if !oneDigit {
			//	fmt.Println("   - Missing digit")
			//}
			//if !oneVowel {
			//	fmt.Println("   - Missing vowel")
			//}
			//if !oneConsonant {
			//	fmt.Println("   - Missing consonant")
			//}
			//if repeated {
			//	fmt.Println("   - Repeated character")
			//}
		}
	}
	fmt.Println("Total good passwords: ", goodPasswords)

}

func IsVowel(r rune) bool {
	r = unicode.ToLower(r)
	normalized := []rune(string(norm.NFD.Bytes([]byte(string(r)))))
	// fmt.Printf("%q\n", normalized)
	return regexp.MustCompile(`[aeiou]`).MatchString(string(normalized[0]))
}

func IsConsonant(r rune) bool {
	return unicode.IsLetter(r) && !IsVowel(r)
}

// readLines reads a file and returns the lines as a slice of strings.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
