package main

import "testing"

func TestPartOne1(t *testing.T) {
	want := 25
	output := partOne(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("partOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	want := 362
	output := partOne(readInput("./input.txt"))

	if want != output {
		t.Errorf("partOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 286
	output := partTwo(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("partTwo(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 29895
	output := partTwo(readInput("./input.txt"))

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
