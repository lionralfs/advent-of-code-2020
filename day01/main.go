package main

import (
	"fmt"
	"os"
)

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := 0

	fmt.Printf("[Part1]: %v\n", result)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := 0

	fmt.Printf("[Part2]: %v\n", result)
}

func main() {
	partOne()
	partTwo()
}
