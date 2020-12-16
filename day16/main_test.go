package main

import "testing"

func TestPartOne1(t *testing.T) {
	rules, _, tickets := readInput("./testinput1.txt")
	want := 71
	output := totalErrorRate(tickets, rules)

	if want != output {
		t.Errorf("totalErrorRate(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	rules, _, tickets := readInput("./input.txt")
	want := 21071
	output := totalErrorRate(tickets, rules)

	if want != output {
		t.Errorf("totalErrorRate(input.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 1716
	output := partTwo(readInput("./testinput2.txt"))

	if want != output {
		t.Errorf("partTwo(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 3429967441937
	output := partTwo(readInput("./input.txt"))

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
