package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(filename string) []int {
	bytes, _ := ioutil.ReadFile(filename)

	input := strings.Split(string(bytes), "\n")

	result := make([]int, len(input))

	for i, line := range input {
		num, _ := strconv.Atoi(line)

		result[i] = num
	}

	return result
}

func main() {
	nums := readInput("./input.txt")

	fmt.Printf("[Part1]: %v\n", partOne(nums, 25))
	fmt.Printf("[Part2]: %v\n", partTwo(nums, 25))
}

func partOne(nums []int, preambleSize int) int {
	low := 0
	high := preambleSize

	for high <= len(nums)-1 {
		preamble := nums[low:high]
		next := nums[high]

		if !containsSum(preamble, next) {
			return next
		}
		low++
		high++
	}

	return -1
}

func partTwo(nums []int, preambleSize int) int {
	return findContiguousSum(nums, partOne(nums, preambleSize))
}

func findContiguousSum(nums []int, sum int) int {
	for low := 0; low < len(nums); low++ {
		currentSum := nums[low]
		for high := low + 1; high < len(nums); high++ {
			currentSum += nums[high]

			if currentSum == sum {
				smallest, highest := getSmallestAndHighest(nums[low : high+1])
				return smallest + highest
			}

			if currentSum > sum {
				break
			}
		}
	}

	return -1
}

func getSmallestAndHighest(nums []int) (int, int) {
	smallest, highest := nums[0], nums[0]

	for _, num := range nums {
		if num < smallest {
			smallest = num
		}

		if num > highest {
			highest = num
		}
	}

	return smallest, highest
}

func containsSum(nums []int, sum int) bool {
	for i, n1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]

			if n1 == n2 {
				continue
			}

			if n1+n2 == sum {
				return true
			}
		}
	}

	return false
}
