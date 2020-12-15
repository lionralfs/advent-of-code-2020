package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./input.txt")
	fmt.Printf("[Part1]: %v\n", calc(input, 2020))
	fmt.Printf("[Part2]: %v\n", calc(input, 30000000))
}

func readInput(filename string) []int {
	bytes, _ := ioutil.ReadFile(filename)
	stringList := strings.Split(string(bytes), ",")

	result := make([]int, len(stringList))

	for i, s := range stringList {
		num, _ := strconv.Atoi(s)
		result[i] = num
	}

	return result
}

func calc(input []int, turns int) int {
	lastSpokenMap := make(map[int]int)

	lastNumber := input[len(input)-1]

	for i, num := range input {
		lastSpokenMap[num] = i + 1
	}

	for turn := len(input) + 1; turn <= turns; turn++ {
		lastSpoken := lastSpokenMap[lastNumber]

		var next int
		if lastSpoken == 0 {
			next = 0
		} else {
			next = turn - 1 - lastSpoken
		}
		lastSpokenMap[lastNumber] = turn - 1
		lastNumber = next
	}

	return lastNumber
}
