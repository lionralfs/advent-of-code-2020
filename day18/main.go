package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	newLineRe := regexp.MustCompile(`\r?\n`)

	sum1 := 0
	sum2 := 0
	for _, line := range newLineRe.Split(string(bytes), -1) {
		sum1 += evaluateEquationPartOne(line)
		sum2 += evaluateEquationPartTwo(line)
	}

	fmt.Printf("[Part1]: %v\n", sum1)
	fmt.Printf("[Part2]: %v\n", sum2)
}

func evaluateEquationPartOne(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	return evaluate(strings.Split(input, ""))
}

func evaluateEquationPartTwo(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	return parseMulExpressions(input)
}

// not very clean
func evaluate(parts []string) int {
	result := -1
	inParentheses := 0
	latestOperator := ""

	chunk := []string{}

	for _, part := range parts {
		switch part {
		case "(":
			inParentheses++
			if inParentheses > 1 {
				chunk = append(chunk, "(")
			}
			break
		case ")":
			inParentheses--
			if inParentheses > 0 {
				chunk = append(chunk, ")")
			} else {
				innerVal := evaluate(chunk)
				chunk = []string{}

				if result == -1 {
					result = innerVal
					break
				}

				switch latestOperator {
				case "+":
					result += innerVal
					break
				case "*":
					result *= innerVal
					break
				}
			}
			break
		case "+":
			fallthrough
		case "*":
			if inParentheses > 0 {
				chunk = append(chunk, part)
				break
			}
			latestOperator = part
			break
		default:
			if inParentheses > 0 {
				chunk = append(chunk, part)
				break
			}

			val, _ := strconv.Atoi(part)
			if result == -1 {
				result = val
			} else {
				switch latestOperator {
				case "+":
					result += val
					break
				case "*":
					result *= val
					break
				}
			}
		}
	}
	return result
}

func splitBy(input string, operator string) []string {
	inParentheses := 0
	chunk := ""
	result := []string{}

	for _, char := range input {
		if char == '(' {
			inParentheses++
		}
		if char == ')' {
			inParentheses--
		}
		if inParentheses == 0 && string(char) == operator {
			result = append(result, chunk)
			chunk = ""
		} else {
			chunk += string(char)
		}
	}

	result = append(result, chunk)

	return result
}

func parseMulExpressions(input string) int {
	stringParts := splitBy(input, "*")

	product := 1
	for _, str := range stringParts {
		product *= parsePlusExpressions(str)
	}

	return product
}

func parsePlusExpressions(input string) int {
	sum := 0
	nums := splitBy(input, "+")

	for _, num := range nums {
		if num[0] == '(' {
			sum += parseMulExpressions(num[1 : len(num)-1])
			continue
		}

		asInt, _ := strconv.Atoi(num)
		sum += asInt
	}

	return sum
}
