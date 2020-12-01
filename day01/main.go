package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := make([]int, 0)

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		list = append(list, mass)
	}

	fmt.Printf("[Part1]: %v\n", twoSummands(list))
}

func twoSummands(nums []int) int {
	for i, val1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			val2 := nums[j]

			if val1+val2 == 2020 {
				return val1 * val2
			}
		}
	}

	return -1
}

func threeSummands(nums []int) int {
	for i, val1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			val2 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				val3 := nums[k]

				if val1+val2+val3 == 2020 {
					return val1 * val2 * val3
				}
			}
		}
	}

	return -1
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := make([]int, 0)

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		list = append(list, mass)
	}

	fmt.Printf("[Part2]: %v\n", threeSummands(list))
}

func main() {
	partOne()
	partTwo()
}
