package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Task contains todays puzzle input
type Task struct {
	target int
	busIDs []int
}

// Bus is a struct for part 2
type Bus struct {
	offset int
	ID     int
}

func readInput(filename string) Task {
	bytes, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile(`\r?\n`)

	lines := re.Split(string(bytes), -1)

	target, _ := strconv.Atoi(lines[0])

	busIDStrings := strings.Split(lines[1], ",")
	busIDs := make([]int, len(busIDStrings))
	for i, busID := range busIDStrings {
		if busID == "x" {
			busIDs[i] = -1
			continue
		}
		busID, _ := strconv.Atoi(busID)
		busIDs[i] = busID
	}

	return Task{target, busIDs}
}

func main() {
	fmt.Printf("[Part1]: %v\n", partOne(readInput("./input.txt")))
	fmt.Printf("[Part2]: %v\n", partTwo(readInput("./input.txt")))
}

func partOne(task Task) int {
	bestTime := task.target * 2
	bestBusID := -1

	for _, busID := range task.busIDs {
		if busID == -1 {
			continue
		}

		closest := ((task.target / busID) + 1) * busID

		if closest < bestTime {
			bestTime = closest
			bestBusID = busID
		}
	}

	return (bestTime - task.target) * bestBusID
}

func partTwo(task Task) int {
	list := make([]Bus, 0)

	for i, busID := range task.busIDs {
		if busID == -1 {
			continue
		}
		list = append(list, Bus{i, busID})
	}

	sort.Slice(list, func(a, b int) bool {
		return list[a].ID > list[b].ID
	})

	firstSync := firstSync(list[0], list[1])
	lcm := list[0].ID * list[1].ID

	for i := 1; ; i++ {
		target := firstSync + i*lcm

		if allHitTarget(list[1:], target) {
			return target
		}
	}
}

func allHitTarget(list []Bus, target int) bool {
	for _, bus := range list {
		if (target+bus.offset)%bus.ID != 0 {
			return false
		}
	}
	return true
}

func firstSync(a, b Bus) int {
	for i := 0; ; i++ {
		if (i+a.offset)%a.ID == 0 && (i+b.offset)%b.ID == 0 {
			return i
		}
	}
}
