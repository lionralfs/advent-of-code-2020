package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// A Coordinate in a 3D/4D space
type Coordinate struct {
	x, y, z, w int
}

func main() {
	part1, part2 := solve("./input.txt")

	fmt.Printf("[Part1]: %v\n", part1)
	fmt.Printf("[Part2]: %v\n", part2)
}

func solve(inputfile string) (int, int) {
	universe1 := readInput(inputfile)
	universe2 := universe1

	for i := 1; i <= 6; i++ {
		universe1 = nextCycle(universe1, getNeighbors3D)
		universe2 = nextCycle(universe2, getNeighbors4D)
	}

	return len(universe1), len(universe2)
}

func readInput(filename string) map[Coordinate]bool {
	result := map[Coordinate]bool{}

	bytes, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile(`\r?\n`)
	for y, line := range re.Split(string(bytes), -1) {
		for x, c := range strings.Split(line, "") {
			if c == "#" {
				result[Coordinate{x, y, 0, 0}] = true
			}
		}
	}

	return result
}

func nextCycle(universe map[Coordinate]bool, neighborFn func(Coordinate) []Coordinate) map[Coordinate]bool {
	newUniverse := map[Coordinate]bool{}
	activeNeighboursCount := map[Coordinate]int{}

	// go through all currently active coordinates
	for coordinate := range universe {
		// set a counter to how many active neighbors they have
		activeNeighbors := 0

		// go through each neighbor
		for _, neighbor := range neighborFn(coordinate) {
			// if that neighbor is active, increase the counter
			if universe[neighbor] {
				activeNeighbors++
			}

			// also, increment the amount of active neighbors that that neighbor has
			activeNeighboursCount[neighbor]++
		}

		// if an active coordinate has two or three active neighbors, it stays active
		if activeNeighbors == 2 || activeNeighbors == 3 {
			newUniverse[coordinate] = true
		}
	}

	// go through all other inactive coordinates and check if they get activated
	for coordinate, neighborCount := range activeNeighboursCount {
		if !universe[coordinate] && neighborCount == 3 {
			newUniverse[coordinate] = true
		}
	}

	return newUniverse
}

func getNeighbors4D(coord Coordinate) []Coordinate {
	result := []Coordinate{}

	for x := coord.x - 1; x <= coord.x+1; x++ {
		for y := coord.y - 1; y <= coord.y+1; y++ {
			for z := coord.z - 1; z <= coord.z+1; z++ {
				for w := coord.w - 1; w <= coord.w+1; w++ {
					if x == coord.x && y == coord.y && z == coord.z && w == coord.w {
						continue
					}
					result = append(result, Coordinate{x, y, z, w})
				}
			}
		}
	}

	return result
}

func getNeighbors3D(coord Coordinate) []Coordinate {
	result := []Coordinate{}

	for x := coord.x - 1; x <= coord.x+1; x++ {
		for y := coord.y - 1; y <= coord.y+1; y++ {
			for z := coord.z - 1; z <= coord.z+1; z++ {
				if x == coord.x && y == coord.y && z == coord.z {
					continue
				}
				result = append(result, Coordinate{x, y, z, coord.w})
			}
		}
	}

	return result
}
