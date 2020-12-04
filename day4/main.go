package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

func parsePassport(p string) passport {
	passp := make(passport)
	for _, a := range strings.Fields(p) {
		kv := strings.Split(a, ":")
		passp[kv[0]] = kv[1]
	}
	return passp
}

func basicPassportFieldCheck(p passport, f []string) bool {
	for _, field := range f {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
}

var hclRegEx = regexp.MustCompile(`#(?:[a-f]|[0-9]){6}`)

func advPassportFieldCheck(p passport) bool {
	// Validate byr
	byr, _ := strconv.Atoi(p["byr"])
	if !(byr >= 1920 && byr <= 2002) {
		return false
	}
	// Validate iyr
	iyr, _ := strconv.Atoi(p["iyr"])
	if !(iyr >= 2010 && iyr <= 2020) {
		return false
	}
	// Validate eyr
	eyr, _ := strconv.Atoi(p["eyr"])
	if !(eyr >= 2020 && iyr <= 2030) {
		return false
	}
	// Validate hgt
	if strings.HasSuffix(p["hgt"], "cm") {
		hgtS := strings.TrimSuffix(p["hgt"], "cm")
		hgt, _ := strconv.Atoi(hgtS)
		if !(hgt >= 150 && hgt <= 193) {
			return false
		}
	} else if strings.HasSuffix(p["hgt"], "in") {
		hgtS := strings.TrimSuffix(p["hgt"], "in")
		hgt, _ := strconv.Atoi(hgtS)
		if !(hgt >= 59 && hgt <= 76) {
			return false
		}
	} else {
		return false
	}
	// Validate hcl
	if !hclRegEx.MatchString(p["hcl"]) {
		return false
	}
	// Validate ecl
	approvedColors := map[string]struct{}{"amb": struct{}{}, "blu": struct{}{}, "brn": struct{}{}, "gry": struct{}{}, "grn": struct{}{}, "hzl": struct{}{}, "oth": struct{}{}} // What an ugly hashset in Go...
	if _, ok := approvedColors[p["ecl"]]; !ok {
		return false
	}
	// Validate pid
	if len(p["pid"]) == 9 {
		_, err := strconv.Atoi(p["pid"])
		if err != nil {
			return false
		}
	}
	// Validate cid
	// ignoring...
	return true
}

func main() {
	// Read input file
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedData := strings.Split(string(data), "\r\n\r\n") // Will not work on files with LF EOL sequence (linux)
	basicValids := 0
	advValids := 0
	for _, p := range parsedData {
		parsedPassport := parsePassport(p)

		if basicPassportFieldCheck(parsedPassport, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}) {
			basicValids++
		}
		if advPassportFieldCheck(parsedPassport) {
			advValids++
		}
	}
	fmt.Println("Answer (part 1):", basicValids)
	fmt.Println("Answer (part 2):", advValids)
}
