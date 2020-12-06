package main

import "testing"

func TestCountPartOne1(t *testing.T) {
	groups := readInput("./testinput1.txt")
	want := 6
	output := countPartOne(groups)

	if want != output {
		t.Errorf("countPartOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountPartOne2(t *testing.T) {
	groups := readInput("./testinput2.txt")
	want := 11
	output := countPartOne(groups)

	if want != output {
		t.Errorf("countPartOne(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestCountPartOne3(t *testing.T) {
	groups := readInput("./input.txt")
	want := 6799
	output := countPartOne(groups)

	if want != output {
		t.Errorf("countPartOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestCountPartTwo1(t *testing.T) {
	groups := readInput("./testinput1.txt")
	want := 3
	output := countPartTwo(groups)

	if want != output {
		t.Errorf("countPartTwo(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountPartTwo2(t *testing.T) {
	groups := readInput("./testinput2.txt")
	want := 6
	output := countPartTwo(groups)

	if want != output {
		t.Errorf("countPartTwo(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestCountPartTwo3(t *testing.T) {
	groups := readInput("./input.txt")
	want := 3354
	output := countPartTwo(groups)

	if want != output {
		t.Errorf("countPartTwo(input.txt) = %v, expected %v", output, want)
	}
}
