package main

import "testing"

func TestPartOne1(t *testing.T) {
	want := 127
	output := partOne(readInput("./testinput1.txt"), 5)

	if want != output {
		t.Errorf("partOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	want := 10884537
	output := partOne(readInput("./input.txt"), 25)

	if want != output {
		t.Errorf("partOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 62
	output := partTwo(readInput("./testinput1.txt"), 5)

	if want != output {
		t.Errorf("partTwo(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 1261309
	output := partTwo(readInput("./input.txt"), 25)

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
