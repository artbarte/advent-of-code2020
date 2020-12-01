package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input file
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	// Prepare hashset for lookup
	dataHashSet := map[int]struct{}{}
	for s.Scan() {
		// Parse input to slice of string
		l, _ := strconv.Atoi(s.Text())
		dataHashSet[l] = struct{}{}
	}
	f.Close()
	// PART 1
	// Iterate over hashset and look for condition
	for l := range dataHashSet {
		dataHashSet[l] = struct{}{}
		if _, exists := dataHashSet[2020-l]; exists {
			// If found print naswer and exit program
			fmt.Println("Answer (part 1): ", l*(2020-l))
			break
		}
	}
	// PART 2
	for l := range dataHashSet {
		for m := range dataHashSet {
			if _, exists := dataHashSet[2020-l-m]; exists {
				fmt.Println("Answer (part 2): ", l*m*(2020-l-m))
				os.Exit(0)
			}
		}
	}
}
