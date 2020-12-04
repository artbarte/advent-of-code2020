package main

import (
	"fmt"
	"io/ioutil"
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

func checkPassportFields(p passport, f []string) bool {
	for _, field := range f {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
}

func main() {
	// Read input file
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedData := strings.Split(string(data), "\r\n\r\n") // Will not work on files with LF EOL sequence (linux)
	valids := 0
	for _, p := range parsedData {
		parsedPassport := parsePassport(p)

		if checkPassportFields(parsedPassport, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}) {
			valids++
		}
	}
	fmt.Println("Answer:", valids)
}
