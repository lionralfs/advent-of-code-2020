package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// An Action has an "action type" (for example "N" = North) and a value (for example 25)
type Action struct {
	action string
	value  int
}

func readInput(filename string) []Action {
	bytes, _ := ioutil.ReadFile(filename)

	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input := re.ReplaceAllString(string(bytes), "\n")
	list := strings.Split(input, "\n")

	result := make([]Action, len(list))

	for i, line := range list {
		re := regexp.MustCompile(`(N|S|E|W|L|R|F)(\d+)`)
		res := re.FindStringSubmatch(line)

		val, _ := strconv.Atoi(res[2])

		result[i] = Action{res[1], val}
	}

	return result
}

func main() {
	input := readInput("./input.txt")
	fmt.Printf("[Part1]: %v\n", partOne(input))
	fmt.Printf("[Part2]: %v\n", partTwo(input))
}

func partOne(actions []Action) int {
	facing := 0
	x := 0
	y := 0

	for _, action := range actions {
		switch action.action {
		case "N":
			x -= action.value
			break
		case "S":
			x += action.value
			break
		case "E":
			y += action.value
			break
		case "W":
			y -= action.value
			break
		case "L":
			facing -= action.value

			if facing < 0 {
				facing = 360 + facing
			}

			facing %= 360
			break
		case "R":
			facing += action.value
			facing %= 360
			break
		case "F":
			switch facing {
			case 0:
				y += action.value
				break
			case 90:
				x += action.value
				break
			case 180:
				y -= action.value
				break
			case 270:
				x -= action.value
				break
			default:
				panic(errors.New("found invalid angle: " + strconv.Itoa(facing)))
			}
			break
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func partTwo(actions []Action) int {
	x := 0
	y := 0
	waypointX := 10
	waypointY := -1

	for _, action := range actions {
		switch action.action {
		case "N":
			waypointY -= action.value
			break
		case "S":
			waypointY += action.value
			break
		case "E":
			waypointX += action.value
			break
		case "W":
			waypointX -= action.value
			break
		case "L":
			waypointX, waypointY = moveWaypointRight(waypointX, waypointY, (360-action.value)%360)
			break
		case "R":
			waypointX, waypointY = moveWaypointRight(waypointX, waypointY, action.value%360)
			break
		case "F":
			x += action.value * waypointX
			y += action.value * waypointY
			break
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func moveWaypointRight(x, y, angle int) (int, int) {
	switch angle {
	case 0:
		return x, y
	case 90:
		return -y, x
	case 270:
		return y, -x
	case 180:
		return -x, -y
	default:
		panic(errors.New("cant rotate by this angle: " + strconv.Itoa(angle)))
	}
}
