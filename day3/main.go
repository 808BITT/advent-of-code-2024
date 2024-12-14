package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("Provide input file")
	}
	inputFile := os.Args[1]

	part1 := part1(inputFile)
	fmt.Printf("(%s) Part 1: %d\n", inputFile, part1)

	part2 := part2(inputFile)
	fmt.Printf("(%s) Part 2: %d\n", inputFile, part2)
}

func part2(inputFile string) int {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	keepParts := make([]string, 0)
	doParts := strings.Split(string(input), "do()")
	for _, part := range doParts {
		keepPart := strings.Split(part, "don't()")[0]
		keepParts = append(keepParts, keepPart)
	}

	total := 0
	for _, part := range keepParts {
		exp := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		occurrences := exp.FindAllString(string(part), -1)
		for _, occ := range occurrences {
			leftExp := regexp.MustCompile(`\((\d{1,3}),`)
			rightExp := regexp.MustCompile(`,(\d{1,3})\)`)
			leftNum := leftExp.FindStringSubmatch(occ)[1]
			rightNum := rightExp.FindStringSubmatch(occ)[1]
			total += toInt(leftNum) * toInt(rightNum)
		}
	}
	return total
}

func part1(inputFile string) int {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	// regex for finding all mul(###,###) where # can be 1-3 digits
	exp := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	occurrences := exp.FindAllString(string(input), -1)

	total := 0
	for _, occ := range occurrences {
		leftExp := regexp.MustCompile(`\((\d{1,3}),`)
		rightExp := regexp.MustCompile(`,(\d{1,3})\)`)
		leftNum := leftExp.FindStringSubmatch(occ)[1]
		rightNum := rightExp.FindStringSubmatch(occ)[1]
		total += toInt(leftNum) * toInt(rightNum)
	}
	return total
}

func toInt(s string) int {
	var num int
	for _, c := range s {
		num = num*10 + int(c-'0')
	}
	return num
}
