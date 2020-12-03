package main

import (
	"bufio"
	"fmt"
	"os"
)

type treeMap [][31]rune

func countTrees(m treeMap, x, y int) int {
	pos := [2]int{0, 0}
	trees := 0
	for pos[0] < len(m) {

		if m[pos[0]][pos[1]%31] == '#' {
			trees++
		}
		pos[0] += y
		pos[1] += x

	}
	return trees
}

func main() {
	// Read input file
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var areaMap treeMap = [][31]rune{}

	var i int
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		areaMap = append(areaMap, [31]rune{})
		for j, c := range l {
			areaMap[i][j] = c
		}
		i++
	}
	f.Close()

	fmt.Println("Answer: ", countTrees(areaMap, 3, 1))
}
