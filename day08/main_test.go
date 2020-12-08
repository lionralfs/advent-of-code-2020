package main

import "testing"

func TestLoops1(t *testing.T) {
	program := readInput("./testinput1.txt")

	loops(program)

	want := 5
	output := program.accumulator

	if want != output {
		t.Errorf("loops(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestLoops2(t *testing.T) {
	program := readInput("./input.txt")

	loops(program)

	want := 2003
	output := program.accumulator

	if want != output {
		t.Errorf("loops(input.txt) = %v, expected %v", output, want)
	}
}

func TestFixLoop1(t *testing.T) {
	want := 8
	output := partTwo("./testinput1.txt")

	if want != output {
		t.Errorf("partTwo(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestFixLoop2(t *testing.T) {
	want := 1984
	output := partTwo("./input.txt")

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
