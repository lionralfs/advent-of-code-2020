package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	rules := readInput("./input.txt")
	total := countPartOne(rules)

	fmt.Printf("[Part1]: %v\n", total)
}

func countPartOne(rules map[string][]struct {
	count int
	name  string
}) int {

	total := 0
	for i := range rules {
		if i == "shiny gold bags" {
			continue
		}

		if canHoldShinyGoldBag(i, rules) {
			total++
		}
	}

	return total
}

func partTwo() {
	rules := readInput("./input.txt")

	total := countBagsIn("shiny gold bags", rules)

	fmt.Printf("[Part2]: %v\n", total)
}

func countBagsIn(outer string, rules map[string][]struct {
	count int
	name  string
}) int {
	total := 0

	for _, rule := range rules[outer] {
		total += rule.count
		total += rule.count * countBagsIn(rule.name, rules)
	}

	return total
}

func canHoldShinyGoldBag(outer string, rules map[string][]struct {
	count int
	name  string
}) bool {

	if outer == "shiny gold bags" {
		return true
	}

	for _, rule := range rules[outer] {
		if canHoldShinyGoldBag(rule.name, rules) {
			return true
		}
	}

	return false
}

func readInput(filename string) map[string][]struct {
	count int
	name  string
} {
	bytes, _ := ioutil.ReadFile(filename)

	input := string(bytes)

	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")

	rawRules := strings.Split(input, "\n")

	rules := make(map[string][]struct {
		count int
		name  string
	})

	for _, rawRule := range rawRules {
		parts := strings.Split(rawRule[:len(rawRule)-1], " contain ")
		outer := parts[0]
		inner := strings.Split(parts[1], ", ")

		list := make([]struct {
			count int
			name  string
		}, 0)

		for _, innerRaw := range inner {
			if innerRaw == "no other bags" {
				continue
			}

			re := regexp.MustCompile(`(\d) (.*)`)

			result := re.FindStringSubmatch(innerRaw)
			bagCount, _ := strconv.Atoi(result[1])
			bagName := result[2]
			if bagCount == 1 {
				bagName = strings.ReplaceAll(result[2], "bag", "bags")
			}

			value := struct {
				count int
				name  string
			}{count: bagCount, name: bagName}

			list = append(list, value)
		}

		rules[outer] = list
	}

	return rules
}
