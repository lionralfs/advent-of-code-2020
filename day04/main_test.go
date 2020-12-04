package main

import "testing"

func TestCountValidPasswords1(t *testing.T) {
	passports := readInput("./testinput1.txt")
	want := 2
	output := countValidPassports(passports)

	if want != output {
		t.Errorf("countValidPassports(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountValidPasswords2(t *testing.T) {
	passports := readInput("./input.txt")
	want := 233
	output := countValidPassports(passports)

	if want != output {
		t.Errorf("countValidPassports(input.txt) = %v, expected %v", output, want)
	}
}

func TestCountValidPasswordsPartTwo1(t *testing.T) {
	passports := readInput("./testinput2.txt")
	want := 0
	output := countValidPassportsPartTwo(passports)

	if want != output {
		t.Errorf("countValidPassportsPartTwo(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestCountValidPasswordsPartTwo2(t *testing.T) {
	passports := readInput("./testinput3.txt")
	want := 4
	output := countValidPassportsPartTwo(passports)

	if want != output {
		t.Errorf("countValidPassportsPartTwo(testinput3.txt) = %v, expected %v", output, want)
	}
}
