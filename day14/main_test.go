package main

import "testing"

func TestPartOne1(t *testing.T) {
	want := 165
	output := partOne(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("partOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	want := 5902420735773
	output := partOne(readInput("./input.txt"))

	if want != output {
		t.Errorf("partOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 208
	output := partTwo(readInput("./testinput2.txt"))

	if want != output {
		t.Errorf("partTwo(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 3801988250775
	output := partTwo(readInput("./input.txt"))

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
