package main

import "testing"

func TestPart1(t *testing.T) {
	entries := parseInput("./testinput1.txt")
	want := 2

	if output := countValidPasswords(entries); output != want {
		t.Errorf("countValidPasswords(testinput1.txt) = %v; want %d", output, want)
	}
}

func TestPart2OnRealInput(t *testing.T) {
	entries := parseInput("./input.txt")
	want := 673

	if output := countValidPasswordsPartTwo(entries); output != want {
		t.Errorf("countValidPasswordsPartTwo(input.txt) = %v; want %d", output, want)
	}
}
