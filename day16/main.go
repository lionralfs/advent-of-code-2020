package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// A Rule has a name and 2 intervals
type Rule struct {
	name  string
	low1  int
	high1 int
	low2  int
	high2 int
}

// A Ticket has a list of values
type Ticket struct {
	values []*TicketValue
}

// A TicketValue has a value and a list of possible rule names
// which apply for that value (relevant for part 2)
type TicketValue struct {
	value         int
	possibleRules []string
}

func main() {
	rules, myTicket, nearbyTickets := readInput("./input.txt")

	fmt.Printf("[Part1]: %v\n", totalErrorRate(nearbyTickets, rules))
	fmt.Printf("[Part2]: %v\n", partTwo(rules, myTicket, nearbyTickets))
}

// return the 3 sections in the input (all rules, my ticket, all nearby tickets)
func readInput(filename string) ([]*Rule, *Ticket, []*Ticket) {
	bytes, _ := ioutil.ReadFile(filename)

	doubleNewLineRe := regexp.MustCompile(`\r?\n\r?\n`)
	newLineRe := regexp.MustCompile(`\r?\n`)
	ruleRe := regexp.MustCompile(`^([a-z\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	myTicketRe := regexp.MustCompile(`^your ticket:\r?\n`)
	nearbyTicketsRe := regexp.MustCompile(`^nearby tickets:\r?\n`)

	parts := doubleNewLineRe.Split(string(bytes), -1)

	var rules []*Rule
	for _, rawRule := range newLineRe.Split(parts[0], -1) {
		parts := ruleRe.FindStringSubmatch(rawRule)
		rules = append(rules, &Rule{
			name:  parts[1],
			low1:  toInt(parts[2]),
			high1: toInt(parts[3]),
			low2:  toInt(parts[4]),
			high2: toInt(parts[5]),
		})
	}

	var myTicketValues []*TicketValue
	for _, rawValue := range strings.Split(myTicketRe.ReplaceAllString(parts[1], ""), ",") {
		myTicketValues = append(myTicketValues, &TicketValue{
			value:         toInt(rawValue),
			possibleRules: []string{},
		})
	}

	var nearbyTickets []*Ticket
	for _, rawTicket := range newLineRe.Split(nearbyTicketsRe.ReplaceAllString(parts[2], ""), -1) {
		var ticketValues []*TicketValue
		for _, rawValue := range strings.Split(rawTicket, ",") {
			ticketValues = append(ticketValues, &TicketValue{
				value:         toInt(rawValue),
				possibleRules: []string{},
			})
		}
		nearbyTickets = append(nearbyTickets, &Ticket{ticketValues})
	}

	return rules, &Ticket{myTicketValues}, nearbyTickets
}

func toInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return res
}

func totalErrorRate(tickets []*Ticket, rules []*Rule) int {
	result := 0

	for _, ticket := range tickets {
		result += ticketErrorRate(ticket, rules)
	}

	return result
}

func ticketErrorRate(ticket *Ticket, rules []*Rule) int {
	result := 0

	for _, ticketValue := range ticket.values {
		if !matchesAtLeastOneRule(ticketValue.value, rules) {
			result += ticketValue.value
		}
	}

	return result
}

func matchesAtLeastOneRule(value int, rules []*Rule) bool {
	for _, rule := range rules {
		if matchesRule(value, rule) {
			return true
		}
	}

	return false
}

func matchesRule(value int, rule *Rule) bool {
	if value >= rule.low1 && value <= rule.high1 {
		return true
	}

	if value >= rule.low2 && value <= rule.high2 {
		return true
	}

	return false
}

// modifies original list
func filter(tickets []*Ticket, filterFn func(*Ticket) bool) []*Ticket {
	count := 0
	for _, ticket := range tickets {
		if filterFn(ticket) {
			tickets[count] = ticket
			count++
		}
	}

	return tickets[:count]
}

// modifies original list
func filterStringList(input []string, filterFn func(string) bool) []string {
	count := 0
	for _, ticket := range input {
		if filterFn(ticket) {
			input[count] = ticket
			count++
		}
	}

	return input[:count]
}

func partTwo(rules []*Rule, myTicket *Ticket, tickets []*Ticket) int {
	// step 1: filter out invalid tickets
	validTickets := filter(append(tickets, myTicket), func(ticket *Ticket) bool {
		return ticketErrorRate(ticket, rules) == 0
	})

	// step 2: enrich ticket values with rules that could apply to them
	for _, ticket := range validTickets {
		for _, ticketValue := range ticket.values {
			for _, rule := range rules {
				if matchesRule(ticketValue.value, rule) {
					ticketValue.possibleRules = append(ticketValue.possibleRules, rule.name)
				}
			}
		}
	}

	// step 3: make list of possible rule values where each index i in the list stands for the possibilities
	// of field i being that rule
	possibleRules := make([][]string, len(rules))
	for i := range possibleRules {
		possibleRules[i] = []string{}

		for _, rule := range rules {
			if isValidRuleForPosition(validTickets, rule.name, i) {
				possibleRules[i] = append(possibleRules[i], rule.name)
			}
		}
	}

	// step 4: recursively shorten the possible rules
	alreadyShortened := make(map[int]bool)
	shortened := shortenRules(possibleRules, &alreadyShortened)

	// step 5: find all fields that start with "departure" and multiply the values from my ticket together
	departureRe := regexp.MustCompile(`^departure`)
	product := 1
	for i, values := range shortened {
		if departureRe.MatchString(values[0]) {
			product *= myTicket.values[i].value
		}
	}

	return product
}

func isValidRuleForPosition(tickets []*Ticket, ruleName string, position int) bool {
	for _, ticket := range tickets {
		if !includes(ticket.values[position].possibleRules, ruleName) {
			return false
		}
	}

	return true
}

func includes(haystack []string, needle string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}

func shortenRules(positions [][]string, alreadyShortened *map[int]bool) [][]string {
	totalShortenable := 0
	toShorten := -1

	for i := 0; i < len(positions); i++ {
		if len(positions[i]) == 1 {
			if !(*alreadyShortened)[i] {
				toShorten = i
			}
			totalShortenable++
		}
	}

	if toShorten == -1 {
		fmt.Println("cant shorten", positions)
	}

	if totalShortenable == len(positions) || toShorten == -1 {
		return positions
	}

	(*alreadyShortened)[toShorten] = true

	return shortenRules(
		removeRuleFromOthersExceptPosition(
			positions,
			positions[toShorten][0],
			toShorten,
		),
		alreadyShortened,
	)
}

func removeRuleFromOthersExceptPosition(positions [][]string, ruleName string, except int) [][]string {
	for i := 0; i < len(positions); i++ {
		if i == except {
			continue
		}

		positions[i] = filterStringList(positions[i], func(val string) bool {
			return val != ruleName
		})
	}

	return positions
}
