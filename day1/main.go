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
	// Iterate over hashset and look for condition
	for l := range dataHashSet {
		if _, exists := dataHashSet[2020-l]; exists {
			// If found print naswer and exit program
			fmt.Println("Answer: ", l*(2020-l))
			os.Exit(0)
		}
		dataHashSet[l] = struct{}{}
	}
}
