package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type policy struct {
	letter    rune
	countLow  int
	countHigh int
}

type entry struct {
	policy   policy
	password string
}

func parseInput(filename string) []entry {
	bytes, _ := ioutil.ReadFile(filename)

	entries := make([]entry, 0)

	for _, line := range strings.Split(string(bytes), "\n") {
		parts := strings.Split(line, " ")

		count := strings.Split(parts[0], "-")

		low, _ := strconv.Atoi(count[0])
		high, _ := strconv.Atoi(count[1])
		letter := []rune(parts[1])[0]

		e := entry{
			policy: policy{
				countLow:  low,
				countHigh: high,
				letter:    letter,
			},
			password: parts[2],
		}

		entries = append(entries, e)
	}

	return entries
}

func main() {
	partOne()
}

func partOne() {
	entries := parseInput("./input.txt")

	fmt.Printf("[Part1]: %v\n", countValidPasswords(entries))
	fmt.Printf("[Part2]: %v\n", countValidPasswordsPartTwo(entries))
}

func countValidPasswords(entries []entry) int {
	result := 0
	for _, entry := range entries {
		count := 0

		for _, c := range entry.password {
			if c == entry.policy.letter {
				count++
			}
		}

		if count >= entry.policy.countLow && count <= entry.policy.countHigh {
			result++
		}
	}

	return result
}

func countValidPasswordsPartTwo(entries []entry) int {
	result := 0

	for _, entry := range entries {
		count := 0

		if rune(entry.password[entry.policy.countLow-1]) == entry.policy.letter {
			count++
		}

		if rune(entry.password[entry.policy.countHigh-1]) == entry.policy.letter {
			count++
		}

		if count == 1 {
			result++
		}
	}

	return result
}
