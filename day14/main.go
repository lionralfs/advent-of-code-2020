package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// An Instruction consists of an operation and up to 2 arguments
type Instruction struct {
	operation string
	arg1      string
	arg2      string
}

func main() {
	fmt.Printf("[Part1]: %v\n", partOne(readInput("./input.txt")))
	fmt.Printf("[Part2]: %v\n", partTwo(readInput("./input.txt")))
}

func readInput(filename string) []Instruction {
	bytes, _ := ioutil.ReadFile(filename)
	lines := regexp.MustCompile(`\r?\n`).Split(string(bytes), -1)

	maskRe := regexp.MustCompile(`mask = (X|0|1){36}`)
	setMemRe := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

	result := make([]Instruction, len(lines))
	for i, line := range lines {
		if maskRe.MatchString(line) {
			mask := strings.Split(line, " = ")[1]

			result[i] = Instruction{operation: "mask", arg1: mask, arg2: ""}
			continue
		}

		if setMemMatches := setMemRe.FindStringSubmatch(line); len(setMemMatches) == 3 {
			result[i] = Instruction{operation: "setMem", arg1: setMemMatches[1], arg2: setMemMatches[2]}
			continue
		}

		panic(errors.New("unhandled program instruction: " + line))
	}

	return result
}

func partOne(program []Instruction) int {
	memory := make(map[int]int)
	currentMask := ""

	for _, instruction := range program {
		switch instruction.operation {
		case "mask":
			currentMask = instruction.arg1
			break
		case "setMem":
			arg1, _ := strconv.Atoi(instruction.arg1)
			arg2, _ := strconv.Atoi(instruction.arg2)

			forceOnes, _ := strconv.ParseInt(strings.ReplaceAll(currentMask, "X", "0"), 2, 0)
			forceZeroes, _ := strconv.ParseInt(strings.ReplaceAll(currentMask, "X", "1"), 2, 0)

			arg2 |= int(forceOnes)
			arg2 &= int(forceZeroes)

			memory[arg1] = arg2
			break
		default:
			panic(errors.New("unknown operation: " + instruction.operation))
		}
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}

	return sum
}

func partTwo(program []Instruction) int {
	memory := make(map[int]int)
	currentMask := ""

	for _, instruction := range program {
		switch instruction.operation {
		case "mask":
			currentMask = instruction.arg1
			break
		case "setMem":
			arg2, _ := strconv.Atoi(instruction.arg2)

			for _, perm := range getAllPermutations(applyMask(currentMask, instruction.arg1)) {
				address, _ := strconv.ParseInt(perm, 2, 0)

				memory[int(address)] = arg2
			}

			break
		default:
			panic(errors.New("unknown operation: " + instruction.operation))
		}
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}

	return sum
}

func getAllPermutations(input string) []string {
	for i, val := range input {
		if string(val) == "X" {
			result := make([]string, 0)

			for _, perm := range getAllPermutations(input[:i] + "0" + input[i+1:]) {
				result = append(result, perm)
			}

			for _, perm := range getAllPermutations(input[:i] + "1" + input[i+1:]) {
				result = append(result, perm)
			}

			return result
		}
	}
	return []string{input}
}

func applyMask(mask string, value string) string {
	val, _ := strconv.Atoi(value)

	result := mask

	asBinary := fmt.Sprintf("%b", val)

	for i := 0; i < len(asBinary); i++ {
		posInVal := len(asBinary) - i - 1
		posInMask := len(mask) - i - 1

		if mask[posInMask] == '0' {
			result = result[:posInMask] + string(asBinary[posInVal]) + result[posInMask+1:]
		}
	}

	return result
}
