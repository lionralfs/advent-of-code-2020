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
	groups := readInput("./input.txt")

	fmt.Printf("[Part1]: %v\n", countPartOne(groups))
}

func partTwo() {
	groups := readInput("./input.txt")

	fmt.Printf("[Part2]: %v\n", countPartTwo(groups))
}

func countPartOne(groups [][]string) int {
	sum := 0
	for _, group := range groups {
		answers := make(map[string]bool)

		for _, person := range group {
			for _, c := range person {
				answers[string(c)] = true
			}
		}

		sum += len(answers)
	}

	return sum
}

func countPartTwo(groups [][]string) int {
	sum := 0
	for _, group := range groups {
		answers := make(map[string]int)

		for _, person := range group {
			for _, c := range person {
				answers[string(c)]++
			}
		}

		for _, answer := range answers {
			if answer == len(group) {
				sum++
			}
		}
	}

	return sum
}

func readInput(filename string) [][]string {
	bytes, _ := ioutil.ReadFile(filename)

	input := string(bytes)

	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")

	rawGroups := strings.Split(input, "\n\n")

	groups := make([][]string, len(rawGroups))

	for i, group := range rawGroups {
		groups[i] = strings.Split(group, "\n")
	}

	return groups
}
