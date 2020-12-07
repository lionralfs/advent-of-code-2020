package main

import "testing"

func TestCountPartOne1(t *testing.T) {
	rules := readInput("./testinput1.txt")

	want := 4
	output := countPartOne(rules)

	if want != output {
		t.Errorf("countPartOne(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountPartOne2(t *testing.T) {
	rules := readInput("./input.txt")

	want := 257
	output := countPartOne(rules)

	if want != output {
		t.Errorf("countPartOne(input.txt) = %v, expected %v", output, want)
	}
}

func TestCountBagsIn1(t *testing.T) {
	rules := readInput("./testinput1.txt")

	want := 32
	output := countBagsIn("shiny gold bags", rules)

	if want != output {
		t.Errorf("countBagsIn(testinput1.txt) = %v, expected %v", output, want)
	}
}

func TestCountBagsIn2(t *testing.T) {
	rules := readInput("./testinput2.txt")

	want := 126
	output := countBagsIn("shiny gold bags", rules)

	if want != output {
		t.Errorf("countBagsIn(testinput2.txt) = %v, expected %v", output, want)
	}
}

func TestCountBagsIn3(t *testing.T) {
	rules := readInput("./input.txt")

	want := 1038
	output := countBagsIn("shiny gold bags", rules)

	if want != output {
		t.Errorf("countBagsIn(input.txt) = %v, expected %v", output, want)
	}
}
