package main

import (
	"fmt"
	"math"
	"os"
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

	// simMap tracks occurence count of right list
	simMap := make(map[int]int)

	// vals are values to check from left list
	vals := make([]int, 0)

	for {
		var x, y int
		_, err := fmt.Fscanf(input, "%d %d\n", &x, &y)
		if err != nil {
			break
		}
		vals = append(vals, x)
		simMap[y] += 1
	}

	totalSimilarity := 0
	for _, val := range vals {
		totalSimilarity += val * simMap[val]
	}

	return totalSimilarity
}

func part1(inputFile string) int {
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	left := make([]int, 0)
	right := make([]int, 0)

	for {
		var x, y int
		_, err := fmt.Fscanf(input, "%d %d\n", &x, &y)
		if err != nil {
			break
		}
		left = append(left, x)
		right = append(right, y)
	}
	sortedLeft := sort(left)
	sortedRight := sort(right)

	totalDifference := 0
	for i := 0; i < len(sortedLeft); i++ {
		totalDifference += int(math.Abs(float64(sortedLeft[i]) - float64(sortedRight[i])))
	}

	return totalDifference
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
