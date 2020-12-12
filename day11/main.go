package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// Field represents a 2D field
type Field struct {
	data [][]string
}

// Print prints the field in human readable format
func (f *Field) Print() {
	field := ""
	for _, row := range f.data {
		for _, col := range row {
			field += col
		}
		field += "\n"
	}
	fmt.Println(field)
}

// Get returns the value at position (x, y)
func (f *Field) Get(x, y int) (string, error) {
	if y < 0 || y >= len(f.data) {
		return "", errors.New("out of bounds")
	}

	if x < 0 || x >= len(f.data[y]) {
		return "", errors.New("out of bounds")
	}

	return f.data[y][x], nil
}

// Set sets the value v at position (x, y)
func (f *Field) Set(x, y int, v string) {
	f.data[y][x] = v
}

var directionsToCheck = [8][2]int{
	// up-left diagonal
	{-1, -1},
	// up
	{0, -1},
	// up-right diagonal
	{1, -1},
	// right
	{1, 0},
	// down-right diagonal
	{1, 1},
	// down
	{0, 1},
	// down-left diagonal
	{-1, 1},
	// left
	{-1, 0},
}

// CountNeighbors returns the amount of directly adjacent occupied seats
func (f *Field) CountNeighbors(x, y int) int {
	result := 0

	for _, direction := range directionsToCheck {
		if val, err := f.Get(x+direction[0], y+direction[1]); err == nil && val == "#" {
			result++
		}
	}

	return result
}

// CountVisibleNeighbors returns the amount of occupied seats that the seat (x, y) can see
func (f *Field) CountVisibleNeighbors(x, y int) int {
	result := 0

	for _, direction := range directionsToCheck {
		for i := 1; ; i++ {

			newX := x + i*direction[0]
			newY := y + i*direction[1]

			val, err := f.Get(newX, newY)
			if err != nil {
				break
			}

			if val == "L" {
				break
			}

			if val == "#" {
				result++
				break
			}
		}
	}

	return result
}

// ApplyNextGeneration transforms the field according to the rules for 1 generation
func (f *Field) ApplyNextGeneration(max int, neighborFn func(x, y int) int) int {
	newField := make(map[struct{ x, y int }]string)

	changes := 0

	for y, row := range f.data {
		for x := range row {
			val, _ := f.Get(x, y)
			if val == "." {
				continue
			}

			n := neighborFn(x, y)

			if n == 0 && val == "L" {
				newField[struct {
					x int
					y int
				}{x, y}] = "#"

				changes++
				continue
			}

			if n >= max && val == "#" {
				newField[struct {
					x int
					y int
				}{x, y}] = "L"

				changes++
			}
		}
	}

	for key, value := range newField {
		f.Set(key.x, key.y, value)
	}

	return changes
}

// CountOccupiedSeats counts the total amount of occupied seats
func (f *Field) CountOccupiedSeats() int {
	result := 0
	for _, row := range f.data {
		for _, val := range row {
			if val == "#" {
				result++
			}
		}
	}
	return result
}

// NewField creates a new field from raw string input
func NewField(input string) *Field {
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")
	rows := strings.Split(input, "\n")

	rowList := make([][]string, len(rows))

	for rowI, row := range rows {
		cols := strings.Split(row, "")
		colList := make([]string, len(cols))

		for colI, col := range cols {
			colList[colI] = col
		}

		rowList[rowI] = colList
	}

	return &Field{
		data: rowList,
	}
}

func main() {
	fmt.Printf("[Part1]: %v\n", partOne("./input.txt"))
	fmt.Printf("[Part2]: %v\n", partTwo("./input.txt"))
}

func partOne(inputfile string) int {
	rawInput, _ := ioutil.ReadFile(inputfile)
	field := NewField(string(rawInput))

	for {
		changes := field.ApplyNextGeneration(4, field.CountNeighbors)

		if changes == 0 {
			return field.CountOccupiedSeats()
		}
	}
}

func partTwo(inputfile string) int {
	rawInput, _ := ioutil.ReadFile(inputfile)
	field := NewField(string(rawInput))

	for {
		changes := field.ApplyNextGeneration(5, field.CountVisibleNeighbors)
		if changes == 0 {
			return field.CountOccupiedSeats()
		}
	}
}
