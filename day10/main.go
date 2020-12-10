package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) []int {
	bytes, _ := ioutil.ReadFile(filename)

	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input := re.ReplaceAllString(string(bytes), "\n")
	list := strings.Split(input, "\n")

	result := make([]int, len(list))

	for i, line := range list {
		num, _ := strconv.Atoi(line)

		result[i] = num
	}

	return result
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	list := readInput("./input.txt")
	fmt.Printf("[Part1]: %v\n", countDiffs(list))
}

func countDiffs(list []int) int {
	sort.Slice(list, func(a, b int) bool {
		return list[a] < list[b]
	})

	diff1s := 0
	diff3s := 0
	prev := 0
	for _, num := range list {
		switch num - prev {
		case 1:
			diff1s++
			break
		case 3:
			diff3s++
		}
		prev = num
	}
	diff3s++
	return diff1s * diff3s
}

func partTwo() {
	list := readInput("./input.txt")

	fmt.Printf("[Part2]: %v\n", countChoices(list))
}

func countChoices(list []int) int {
	list = append(list, 0)
	sort.Slice(list, func(a, b int) bool {
		return list[a] < list[b]
	})

	choicesMap := make(map[int][]int)
	for i, num := range list {
		var choices []int

		for j := 1; j <= 3; j++ {
			if i+j >= len(list) {
				break
			}

			other := list[i+j]

			if other-num <= 3 {
				choices = append(choices, other)
			}
		}

		choicesMap[num] = choices
	}

	results := make(map[int]int)
	results[list[len(list)-1]] = 1

	// the trick is going backwards
	for i := len(list) - 2; i >= 0; i-- {
		num := list[i]
		// look at all the numbers that the current num can reach
		// and add up the possibilities
		for _, e := range choicesMap[num] {
			results[num] += results[e]
		}
	}

	return results[0]
}
