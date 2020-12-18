package main

import "testing"

func TestWithTestInput(t *testing.T) {
	part1, part2 := solve("./testinput1.txt")
	want1, want2 := 112, 848

	if want1 != part1 {
		t.Errorf("solve part 1 (testinput1.txt) = %v, expected %v", part1, want1)
	}

	if want2 != part2 {
		t.Errorf("solve part 2 (testinput1.txt) = %v, expected %v", part2, want2)
	}
}

func TestWithRealInput(t *testing.T) {
	part1, part2 := solve("./input.txt")
	want1, want2 := 255, 2340

	if want1 != part1 {
		t.Errorf("solve part 1 (input.txt) = %v, expected %v", part1, want1)
	}

	if want2 != part2 {
		t.Errorf("solve part 2 (input.txt) = %v, expected %v", part2, want2)
	}
}
