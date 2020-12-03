package main

import "testing"

func TestEncounteredTrees1(t *testing.T) {
	slopes := [][3]int{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

	for _, slope := range slopes {
		output := countEncounteredTrees("./testinput1.txt", slope[0], slope[1])
		want := slope[2]

		if output := output; output != want {
			t.Errorf("countEncounteredTrees(testinput1.txt, %v, %v) = %v; want %d", slope[0], slope[1], output, want)
		}
	}
}

func TestEncounteredTrees2(t *testing.T) {
	slopes := [][3]int{
		{1, 1, 84},
		{3, 1, 198},
		{5, 1, 72},
		{7, 1, 81},
		{1, 2, 53},
	}

	for _, slope := range slopes {
		output := countEncounteredTrees("./input.txt", slope[0], slope[1])
		want := slope[2]

		if output := output; output != want {
			t.Errorf("countEncounteredTrees(input.txt, %v, %v) = %v; want %d", slope[0], slope[1], output, want)
		}
	}
}
