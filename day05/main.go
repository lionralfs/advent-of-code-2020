package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	passes := readInput("./input.txt")

	highest := 0
	for _, pass := range passes {
		row, col := getSeat(pass)

		seatID := (row * 8) + col

		if seatID > highest {
			highest = seatID
		}
	}

	fmt.Printf("[Part1]: %v\n", highest)
}

func partTwo() {
	passes := readInput("./input.txt")

	seats := make(map[int]bool)

	for _, pass := range passes {
		row, col := getSeat(pass)

		seatID := (row * 8) + col
		seats[seatID] = true
	}

	for i := 0; i < 128*8; i++ {
		if !seats[i] {
			if seats[i-1] && seats[i+1] {
				fmt.Printf("[Part2]: %v\n", i)
			}
		}
	}

}

func readInput(filename string) []string {
	bytes, _ := ioutil.ReadFile(filename)

	input := string(bytes)
	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")
	return strings.Split(input, "\n")
}

func binarySearch(input string, low, high int) int {
	inputIndex := 0
	for low < high {
		middle := (low + high) / 2

		if string(input[inputIndex]) == "F" {
			high = middle
		} else {
			low = middle + 1
		}

		inputIndex++
	}

	return low
}

func getSeat(input string) (int, int) {
	rowIdentifier := input[:7]
	colIdentifier := input[7:]
	colIdentifier = strings.ReplaceAll(colIdentifier, "L", "F")
	colIdentifier = strings.ReplaceAll(colIdentifier, "R", "B")

	row := binarySearch(rowIdentifier, 0, 127)
	col := binarySearch(colIdentifier, 0, 7)

	return row, col
}
