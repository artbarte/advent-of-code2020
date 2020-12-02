package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var passRXP = regexp.MustCompile(`(\d+)-(\d+) (\D): (\D+)`)

func passwordValidationFunc(s string) bool {
	parsedPass := passRXP.FindStringSubmatch(s)
	min, _ := strconv.Atoi(parsedPass[1])
	max, _ := strconv.Atoi(parsedPass[2])
	char := parsedPass[3]
	password := parsedPass[4]

	reqCharCount := 0
	for _, c := range password {
		if c == rune(char[0]) {
			reqCharCount++
			if reqCharCount > max {
				return false
			}
		}
	}
	if reqCharCount >= min {
		return true
	}
	return false
}

func main() {
	// Read input file
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var valids int

	s := bufio.NewScanner(f)
	for s.Scan() {
		// Parse input to slice of string
		if passwordValidationFunc(s.Text()) {
			valids++
		}
	}
	f.Close()
	fmt.Println("Answer: ", valids)
}
