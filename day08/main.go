package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// A Program is a list of instructions along with some state
type Program struct {
	instructions       []Instruction
	accumulator        int
	instructionPointer int
}

// An Instruction describes an operation with an argument
type Instruction struct {
	operation string
	argument  int
}

func (p *Program) runNextInstruction() bool {
	if p.instructionPointer >= len(p.instructions) {
		return false
	}

	instruction := p.instructions[p.instructionPointer]

	switch instruction.operation {
	case "acc":
		p.accumulator += instruction.argument
		p.instructionPointer++
		break
	case "jmp":
		p.instructionPointer += instruction.argument
		break
	case "nop":
		p.instructionPointer++
		break
	default:
		panic(errors.New("Unhandled/Unknown operation: " + instruction.operation))
	}

	return true
}

func main() {
	partOne()
	fmt.Printf("[Part2]: %v\n", partTwo("./input.txt"))
}

func partOne() {
	program := readInput("./input.txt")

	loops, _ := loops(program)
	if loops {
		fmt.Printf("[Part1]: %v\n", program.accumulator)
	}
}

func partTwo(filename string) int {
	program := readInput(filename)

	// run once to find the main loop
	_, mainLoop := loops(program)

	// go through each instruction in the main loop,
	// try to find an instruction we can replace,
	// replace it and see if the program still loops
	for instructionIndex := range mainLoop {
		instruction := program.instructions[instructionIndex]

		if instruction.operation == "nop" || instruction.operation == "jmp" {
			other := "jmp"

			if instruction.operation == "jmp" {
				other = "nop"
			}

			instructionCopy := make([]Instruction, len(program.instructions))
			copy(instructionCopy, program.instructions)

			instructionCopy[instructionIndex] = Instruction{
				// replace the operation
				operation: other,
				// keep the argument
				argument: program.instructions[instructionIndex].argument,
			}

			newProgram := &Program{
				instructionPointer: 0,
				accumulator:        0,
				instructions:       instructionCopy,
			}

			// run the program
			if loops, _ := loops(newProgram); !loops {
				return newProgram.accumulator
			}
		}
	}

	return -1
}

// computer scientists hate this one simple trick
func loops(program *Program) (bool, map[int]bool) {
	visited := make(map[int]bool)
	visited[0] = true

	for program.runNextInstruction() {
		if visited[program.instructionPointer] {
			return true, visited
		}

		visited[program.instructionPointer] = true
	}

	return false, visited
}

func readInput(filename string) *Program {
	bytes, _ := ioutil.ReadFile(filename)

	input := string(bytes)

	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")

	rawInstructions := strings.Split(input, "\n")
	instructions := make([]Instruction, len(rawInstructions))

	for i, rawInstruction := range rawInstructions {
		re := regexp.MustCompile(`(acc|jmp|nop) ((\+|-)\d+)`)
		result := re.FindStringSubmatch(rawInstruction)
		if result == nil || len(result) != 4 {
			panic(errors.New("Unknown instruction: " + rawInstruction))
		}

		arg, err := strconv.Atoi(result[2])

		if err != nil {
			panic(errors.New("Not an integer in instruction: " + rawInstruction))
		}

		instructions[i] = Instruction{
			operation: result[1],
			argument:  arg,
		}
	}

	return &Program{
		instructionPointer: 0,
		accumulator:        0,
		instructions:       instructions,
	}
}
