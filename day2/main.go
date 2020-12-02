package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var passRXP = regexp.MustCompile(`(\d+)-(\d+) (\D): (\D+)`)

func passwordValidationFunc(s string) (bool, bool) {
	parsedPass := passRXP.FindStringSubmatch(s)
	min, _ := strconv.Atoi(parsedPass[1])
	max, _ := strconv.Atoi(parsedPass[2])
	char := byte(parsedPass[3][0])
	password := parsedPass[4]

	passedNew := false

	if len(password) > max-1 {
		// Xor logic opeartor ;P
		x := password[min-1] == char
		y := password[max-1] == char
		if (x || y) && (x != y) {
			passedNew = true
		}
	}

	reqCharCount := 0

	for _, c := range password {
		if c == rune(char) {
			reqCharCount++
			if reqCharCount > max {
				return false, passedNew
			}
		}
	}
	if reqCharCount >= min {
		return true, passedNew
	}
	return false, passedNew
}

func main() {
	// Read input file
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var oldValids int
	var newValids int

	s := bufio.NewScanner(f)
	for s.Scan() {
		// Parse input to slice of string
		oldPassPolicy, newPassPolicy := passwordValidationFunc(s.Text())
		if oldPassPolicy {
			oldValids++
		}
		if newPassPolicy {
			newValids++
		}
	}
	f.Close()
	fmt.Println("Answer (part 1): ", oldValids)
	fmt.Println("Answer (part 2): ", newValids)
}
