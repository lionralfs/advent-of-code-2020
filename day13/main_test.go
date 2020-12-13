package main

import "testing"

func TestPartOne1(t *testing.T) {
	want := 295
	output := partOne(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("partOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	want := 3035
	output := partOne(readInput("./input.txt"))

	if want != output {
		t.Errorf("partOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 1068781
	output := partTwo(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("partTwo(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 3417
	output := partTwo(readInput("./testinput2.txt"))

	if want != output {
		t.Errorf("partTwo(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo3(t *testing.T) {
	want := 754018
	output := partTwo(readInput("./testinput3.txt"))

	if want != output {
		t.Errorf("partTwo(testinput3.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo4(t *testing.T) {
	want := 779210
	output := partTwo(readInput("./testinput4.txt"))

	if want != output {
		t.Errorf("partTwo(testinput4.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo5(t *testing.T) {
	want := 1261476
	output := partTwo(readInput("./testinput5.txt"))

	if want != output {
		t.Errorf("partTwo(testinput5.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo6(t *testing.T) {
	want := 1202161486
	output := partTwo(readInput("./testinput6.txt"))

	if want != output {
		t.Errorf("partTwo(testinput6.txt) = %v, expected %v", output, want)
	}
}

func TestPartTwo7(t *testing.T) {
	want := 725169163285238
	output := partTwo(readInput("./input.txt"))

	if want != output {
		t.Errorf("partTwo(input.txt) = %v, expected %v", output, want)
	}
}
