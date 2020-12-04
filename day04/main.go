package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	// Birth Year
	byr string
	// Issue Year
	iyr string
	// Expiration Year
	eyr string
	// Height
	hgt string
	// Hair Color
	hcl string
	// Eye Color
	ecl string
	// Passport ID
	pid string
	// Country ID
	cid string
}

// for part 2
func (p passport) isValid() bool {
	// byr
	{
		byr, err := strconv.Atoi(p.byr)
		if err != nil {
			return false
		}

		if byr < 1920 || byr > 2002 {
			return false
		}
	}

	// iyr
	{
		iyr, err := strconv.Atoi(p.iyr)
		if err != nil {
			return false
		}

		if iyr < 2010 || iyr > 2020 {
			return false
		}
	}

	// eyr
	{
		eyr, err := strconv.Atoi(p.eyr)
		if err != nil {
			return false
		}

		if eyr < 2020 || eyr > 2030 {
			return false
		}
	}

	// hgt
	{
		if len(p.hgt) < 4 {
			return false
		}

		unit := p.hgt[len(p.hgt)-2:]
		if unit == "in" {
			parts := strings.Split(p.hgt, unit)
			height, err := strconv.Atoi(parts[0])
			if err != nil {
				return false
			}
			if height < 59 || height > 76 {
				return false
			}
		} else if unit == "cm" {
			parts := strings.Split(p.hgt, unit)
			height, err := strconv.Atoi(parts[0])
			if err != nil {
				return false
			}
			if height < 150 || height > 193 {
				return false
			}
		} else {
			return false
		}
	}

	// hcl
	{
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		if !re.MatchString(p.hcl) {
			return false
		}
	}

	// ecl
	{
		re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		if !re.MatchString(p.ecl) {
			return false
		}
	}

	// pid
	{
		re := regexp.MustCompile(`^[0-9]{9}$`)
		if !re.MatchString(p.pid) {
			return false
		}
	}

	return true
}

func readInput(inputfile string) []passport {
	bytes, _ := ioutil.ReadFile(inputfile)
	input := string(bytes)

	// for windows
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")

	rawPassports := strings.Split(input, "\n\n")

	passports := make([]passport, 0)
	for _, rawPassport := range rawPassports {
		passport := passport{}
		for _, field := range strings.Split(re.ReplaceAllString(rawPassport, " "), " ") {
			splitField := strings.Split(field, ":")
			key := splitField[0]
			value := splitField[1]

			switch key {
			case "byr":
				passport.byr = value
				break
			case "iyr":
				passport.iyr = value
				break
			case "eyr":
				passport.eyr = value
				break
			case "hgt":
				passport.hgt = value
				break
			case "hcl":
				passport.hcl = value
				break
			case "ecl":
				passport.ecl = value
				break
			case "pid":
				passport.pid = value
				break
			case "cid":
				passport.cid = value
				break
			default:
				panic(errors.New("unknown passport field " + key))
			}
		}
		passports = append(passports, passport)
	}
	return passports
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	passports := readInput("./input.txt")

	fmt.Printf("[Part1]: %v\n", countValidPassports(passports))
}

func partTwo() {
	passports := readInput("./input.txt")

	fmt.Printf("[Part2]: %v\n", countValidPassportsPartTwo(passports))
}

func countValidPassports(passports []passport) int {
	validPassports := 0
	for _, passport := range passports {

		if passport.byr == "" {
			continue
		}
		if passport.iyr == "" {
			continue
		}
		if passport.eyr == "" {
			continue
		}
		if passport.hgt == "" {
			continue
		}
		if passport.hcl == "" {
			continue
		}
		if passport.ecl == "" {
			continue
		}
		if passport.pid == "" {
			continue
		}
		// if passport.cid == "" {
		// 	continue
		// }

		validPassports++
	}

	return validPassports
}

func countValidPassportsPartTwo(passports []passport) int {
	validPassports := 0
	for _, passport := range passports {
		if passport.isValid() {
			validPassports++
		}
	}

	return validPassports
}
