package main

import (
	"fmt"
	"io/ioutil"
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
func (f *Field) Get(x, y int) string {
	return f.data[y][x]
}

// Set sets the value v at position (x, y)
func (f *Field) Set(x, y int, v string) {
	f.data[y][x] = v
}

// NewField creates a new field from raw string input
func NewField(input string) *Field {
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
	partOne()
	partTwo()
}

func partOne() {
	fmt.Printf("[Part1]: %v\n", countEncounteredTrees("./input.txt", 3, 1))
}

func partTwo() {
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, slope := range slopes {
		trees := countEncounteredTrees("./input.txt", slope[0], slope[1])
		product *= trees
	}

	fmt.Printf("[Part2]: %v\n", product)
}

func countEncounteredTrees(inputfile string, slopeX, slopeY int) int {
	bytes, _ := ioutil.ReadFile(inputfile)

	field := NewField(string(bytes))

	posX, posY := 0, 0
	sizeX, sizeY := len(field.data[0]), len(field.data)

	encouteredTrees := 0

	for posY < sizeY {
		// check if on tree
		if field.Get(posX, posY) == "#" {
			encouteredTrees++
			field.Set(posX, posY, "X")
		} else {
			field.Set(posX, posY, "O")
		}

		// field.Print()

		// update position
		posX = (posX + slopeX) % sizeX
		posY += slopeY
	}

	return encouteredTrees
}
