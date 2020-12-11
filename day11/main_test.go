package main

import "testing"

func TestPartOne1(t *testing.T) {
	want := 37
	output := partOne("./testinput1.txt")

	if want != output {
		t.Errorf("partOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	want := 2113
	output := partOne("./input.txt")

	if want != output {
		t.Errorf("partOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 26
	output := partTwo("./testinput1.txt")

	if want != output {
		t.Errorf("partTwo(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 1865
	output := partTwo("./input.txt")

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
