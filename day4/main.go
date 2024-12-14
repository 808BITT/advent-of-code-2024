package main

import (
	"fmt"
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

	wordSearch := make([][]rune, 0)
	row := 0
	for {
		var line string
		_, err := fmt.Fscanf(input, "%s\n", &line)
		if err != nil {
			break
		}

		wordSearch = append(wordSearch, make([]rune, len(line)))
		for i, c := range line {
			wordSearch[row][i] = c
		}
		row++
	}

	return crossSearchMap(wordSearch)
}

func crossSearchMap(wordSearch [][]rune) int {
	possibilities := make([][]int, len(wordSearch))
	for i := range possibilities {
		possibilities[i] = make([]int, len(wordSearch[i]))
	}

	count := 0
	for y, row := range wordSearch {
		for x, c := range row {
			if c == 'A' {
				if crossCheck(wordSearch, x, y) {
					count++
				}
			}
		}
	}
	return count
}

func crossCheck(wordSearch [][]rune, x, y int) bool {
	downLeft := false
	downRight := false
	if y-1 >= 0 && y+1 < len(wordSearch) && x-1 >= 0 && x+1 < len(wordSearch[y]) {
		fmt.Printf("Checking cross at %d,%d\n", x, y)
		fmt.Printf("Checking %c %c\n", wordSearch[y-1][x-1], wordSearch[y+1][x+1])
		if wordSearch[y-1][x-1] == 'M' && wordSearch[y+1][x+1] == 'S' {
			downLeft = true
		} else if wordSearch[y-1][x-1] == 'S' && wordSearch[y+1][x+1] == 'M' {
			downLeft = true
		}
		fmt.Printf("Checking %c %c\n", wordSearch[y+1][x-1], wordSearch[y-1][x+1])
		if wordSearch[y-1][x+1] == 'M' && wordSearch[y+1][x-1] == 'S' {
			downRight = true
		} else if wordSearch[y-1][x+1] == 'S' && wordSearch[y+1][x-1] == 'M' {
			downRight = true
		}
	}
	if downLeft && downRight {
		return true
	}
	return false
}

func part1(inputFile string) int {
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	wordSearch := make([][]rune, 0)
	row := 0
	for {
		var line string
		_, err := fmt.Fscanf(input, "%s\n", &line)
		if err != nil {
			break
		}

		wordSearch = append(wordSearch, make([]rune, len(line)))
		for i, c := range line {
			wordSearch[row][i] = c
		}
		row++
	}

	return searchMap(wordSearch)
}

// finds X's and then calls xmas to see if it makes 'XMAS'
func searchMap(wordSearch [][]rune) int {
	count := 0
	for y, row := range wordSearch {
		for x, c := range row {
			if c == 'X' {
				count += xmas(wordSearch, x, y)
			}
		}
	}
	return count
}

func xmas(wordSearch [][]rune, x, y int) int {
	count := 0

	// fmt.Printf("Checking XMAS at %d,%d\n", x, y)
	// check up
	if y-3 >= 0 {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y-1][x], wordSearch[y-2][x], wordSearch[y-3][x])
		if wordSearch[y-1][x] == 'M' && wordSearch[y-2][x] == 'A' && wordSearch[y-3][x] == 'S' {
			count++
		}
	}
	// check up-left
	if y-3 >= 0 && x-3 >= 0 {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y-1][x-1], wordSearch[y-2][x-2], wordSearch[y-3][x-3])
		if wordSearch[y-1][x-1] == 'M' && wordSearch[y-2][x-2] == 'A' && wordSearch[y-3][x-3] == 'S' {
			count++
		}
	}
	// check left
	if x-3 >= 0 {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y][x-1], wordSearch[y][x-2], wordSearch[y][x-3])
		if wordSearch[y][x-1] == 'M' && wordSearch[y][x-2] == 'A' && wordSearch[y][x-3] == 'S' {
			count++
		}
	}
	// check down-left
	if y+3 < len(wordSearch) && x-3 >= 0 {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y+1][x-1], wordSearch[y+2][x-2], wordSearch[y+3][x-3])
		if wordSearch[y+1][x-1] == 'M' && wordSearch[y+2][x-2] == 'A' && wordSearch[y+3][x-3] == 'S' {
			count++
		}
	}
	// check down
	if y+3 < len(wordSearch) {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y+1][x], wordSearch[y+2][x], wordSearch[y+3][x])
		if wordSearch[y+1][x] == 'M' && wordSearch[y+2][x] == 'A' && wordSearch[y+3][x] == 'S' {
			count++
		}
	}
	// check down-right
	if y+3 < len(wordSearch) && x+3 < len(wordSearch[y]) {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y+1][x+1], wordSearch[y+2][x+2], wordSearch[y+3][x+3])
		if wordSearch[y+1][x+1] == 'M' && wordSearch[y+2][x+2] == 'A' && wordSearch[y+3][x+3] == 'S' {
			count++
		}
	}
	// check right
	if x+3 < len(wordSearch[y]) {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y][x+1], wordSearch[y][x+2], wordSearch[y][x+3])
		if wordSearch[y][x+1] == 'M' && wordSearch[y][x+2] == 'A' && wordSearch[y][x+3] == 'S' {
			count++
		}
	}
	// check up-right
	if y-3 >= 0 && x+3 < len(wordSearch[y]) {
		// fmt.Printf("Checking %c %c %c\n", wordSearch[y-1][x+1], wordSearch[y-2][x+2], wordSearch[y-3][x+3])
		if wordSearch[y-1][x+1] == 'M' && wordSearch[y-2][x+2] == 'A' && wordSearch[y-3][x+3] == 'S' {
			count++
		}
	}
	return count
}
