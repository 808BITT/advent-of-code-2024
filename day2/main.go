package main

import (
	"bufio"
	"fmt"
	"os"
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
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var rawReports []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rawReports = append(rawReports, scanner.Text())
	}

	safeReports := 0
	for _, report := range rawReports {
		if safe2(report) {
			safeReports++
		}
	}

	return safeReports
}

func safe2(line string) bool {
	var safe bool

	levels := strings.Split(line, " ")
	levelInts := make([]int, len(levels))
	for i, level := range levels {
		fmt.Sscanf(level, "%d", &levelInts[i])
	}

	// now we can remove any one level and check if it passes the 1-3 increasing or decreasing check
	dampener := -1
	combinations := len(levelInts) + 1

	for i := 0; i < combinations; i++ {
		testArray := make([]int, 0)
		if dampener > -1 { // we can try skipping a level
			for j := 0; j < len(levelInts); j++ {
				if j != dampener {
					testArray = append(testArray, levelInts[j])
				}
			}
		} else { // try the original levels no skipping
			for j := 0; j < len(levelInts); j++ {
				testArray = append(testArray, levelInts[j])
			}
		}

		var increasing bool
		if testArray[1] > testArray[0] {
			increasing = true
		} else {
			increasing = false
		}

		for i := 1; i < len(testArray); i++ {
			if increasing {
				if testArray[i]-testArray[i-1] > 3 || testArray[i]-testArray[i-1] < 1 {
					safe = false
					break
				}
				if testArray[i] < testArray[i-1] {
					safe = false
					break
				}
				safe = true
			} else {
				if testArray[i-1]-testArray[i] > 3 || testArray[i-1]-testArray[i] < 1 {
					safe = false
					break
				}
				if testArray[i-1] < testArray[i] {
					safe = false
					break
				}
				safe = true
			}
		}

		if safe { // if we find a safe report we can return
			return safe
		} else { // if not we can try skipping levels using the dampener
			dampener++
		}
	}

	return safe
}

func part1(inputFile string) int {
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var rawReports []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rawReports = append(rawReports, scanner.Text())
	}

	safeReports := 0
	for _, report := range rawReports {
		if safe(report) {
			safeReports++
		}
	}

	return safeReports
}

func safe(line string) bool {
	// safe if all adjacent numbers differ by 1-3 (inclusive)

	levels := strings.Split(line, " ")
	levelInts := make([]int, len(levels))
	for i, level := range levels {
		fmt.Sscanf(level, "%d", &levelInts[i])
	}

	var increasing bool
	if levelInts[1] > levelInts[0] {
		increasing = true
	} else {
		increasing = false
	}

	if increasing {
		for i := 1; i < len(levelInts); i++ {
			if levelInts[i]-levelInts[i-1] > 3 || levelInts[i]-levelInts[i-1] < 1 {
				return false
			}

			if levelInts[i] < levelInts[i-1] {
				return false
			}
		}
		return true
	} else {
		for i := 1; i < len(levelInts); i++ {
			if levelInts[i-1]-levelInts[i] > 3 || levelInts[i-1]-levelInts[i] < 1 {
				return false
			}
			if levelInts[i-1] < levelInts[i] {
				return false
			}
		}
		return true
	}
}
