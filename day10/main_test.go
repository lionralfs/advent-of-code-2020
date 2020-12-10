package main

import "testing"

func TestCountDiffs1(t *testing.T) {
	want := 35
	output := countDiffs(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("countDiffs(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountDiffs2(t *testing.T) {
	want := 220
	output := countDiffs(readInput("./testinput2.txt"))

	if want != output {
		t.Errorf("countDiffs(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestCountDiffs3(t *testing.T) {
	want := 2380
	output := countDiffs(readInput("./input.txt"))

	if want != output {
		t.Errorf("countDiffs(input.txt) = %v, expected %v", output, want)
	}
}

func TestCountChoices1(t *testing.T) {
	want := 8
	output := countChoices(readInput("./testinput1.txt"))

	if want != output {
		t.Errorf("countChoices(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountChoices2(t *testing.T) {
	want := 19208
	output := countChoices(readInput("./testinput2.txt"))

	if want != output {
		t.Errorf("countChoices(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestCountChoices3(t *testing.T) {
	want := 48358655787008
	output := countChoices(readInput("./input.txt"))

	if want != output {
		t.Errorf("countChoices(input.txt) = %v, expected %v", output, want)
	}
}
