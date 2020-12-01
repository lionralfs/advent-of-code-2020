package intcode

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// returns a list of length 4 where position:
// 0 represents the optcode
// 1 represents the mode of the first argument
// 2 represents the mode of the second argument
// 3 represents the mode of the third argument
func getOperation(instruction int) []int {
	result := make([]int, 4)

	result[0] = instruction % 100
	result[1] = (instruction / 100) % 10
	result[2] = (instruction / 1000) % 10
	result[3] = (instruction / 10000) % 10

	return result
}

// A Program can execute a set of instructions (intcode)
// on some inputs which produces some outputs
type Program struct {
	pointer int
	code    map[int]int
	output  int
	inputs  []int
	done    bool
	offset  int
}

// returns the address of the argument
func (p *Program) getArg(mode, offsetFromPointer int) int {
	position := p.pointer
	switch mode {
	case 0:
		return p.code[position+offsetFromPointer]
	case 1:
		return position + offsetFromPointer
	case 2:
		return p.code[position+offsetFromPointer] + p.offset
	}

	panic(errors.New("Unknown mode: " + strconv.Itoa(mode)))
}

// NewProgram creates a new instance based on an intcode
func NewProgram(intcode []int) Program {
	code := map[int]int{}
	for i, e := range intcode {
		code[i] = e
	}

	return Program{
		code:   code,
		offset: 0,
		inputs: []int{},
	}
}

// AddInput adds an input to the end of the input list for the program
func (p *Program) AddInput(input int) {
	p.inputs = append(p.inputs, input)
}

// Run runs the program, returning the next result
func (p *Program) Run() (int, bool) {
	for {
		position := p.pointer
		operation := getOperation(p.code[position])

		switch operation[0] {
		case 1: // addition
			arg1 := p.getArg(operation[1], 1)
			arg2 := p.getArg(operation[2], 2)
			arg3 := p.getArg(operation[3], 3)
			p.code[arg3] = p.code[arg1] + p.code[arg2]

			p.pointer += 4
		case 2: // multiplication
			arg1 := p.getArg(operation[1], 1)
			arg2 := p.getArg(operation[2], 2)
			arg3 := p.getArg(operation[3], 3)
			p.code[arg3] = p.code[arg1] * p.code[arg2]

			p.pointer += 4
		case 3: // input
			arg1 := p.getArg(operation[1], 1)
			// take the first position off the input list and remove it
			input := p.inputs[0]
			p.inputs = p.inputs[1:]
			p.code[arg1] = input

			p.pointer += 2
		case 4: // output
			arg1 := p.getArg(operation[1], 1)
			value := p.code[arg1]
			p.output = value

			p.pointer += 2

			return value, false
		case 5: // jump-if-true
			arg1 := p.getArg(operation[1], 1)
			arg2 := p.getArg(operation[2], 2)

			if p.code[arg1] != 0 {
				p.pointer = p.code[arg2]
			} else {
				p.pointer += 3
			}
		case 6: // jump-if-false
			arg1 := p.getArg(operation[1], 1)
			arg2 := p.getArg(operation[2], 2)

			if p.code[arg1] == 0 {
				p.pointer = p.code[arg2]
			} else {
				p.pointer += 3
			}
		case 7: // less than
			arg1 := p.getArg(operation[1], 1)
			arg2 := p.getArg(operation[2], 2)
			arg3 := p.getArg(operation[3], 3)

			if p.code[arg1] < p.code[arg2] {
				p.code[arg3] = 1
			} else {
				p.code[arg3] = 0
			}

			if arg3 != position {
				p.pointer += 4
			}
		case 8: // equals
			arg1 := p.getArg(operation[1], 1)
			arg2 := p.getArg(operation[2], 2)
			arg3 := p.getArg(operation[3], 3)

			if p.code[arg1] == p.code[arg2] {
				p.code[arg3] = 1
			} else {
				p.code[arg3] = 0
			}

			if arg3 != position {
				p.pointer += 4
			}
		case 9: // adjust offset
			arg1 := p.getArg(operation[1], 1)

			p.offset += p.code[arg1]
			p.pointer += 2
		case 99: // halt
			p.done = true
			return 0, true
		default:
			panic(errors.New("Unknown operation: " + strconv.Itoa(operation[0])))
		}
	}
}

// ReadInput reads the input of the file that's passed as its argument and turns it into an intcode program
func ReadInput(inputFilePath string) []int {
	bytes, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	// split the string of integers at each comma
	// which results in a list of strings
	list := strings.Split(string(bytes), ",")

	// parse the strings to integers
	opcodes := make([]int, len(list))

	for i, e := range list {
		code, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		opcodes[i] = code
	}

	return opcodes
}
