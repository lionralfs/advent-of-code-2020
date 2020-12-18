package main

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "1 + 2 * 3 + 4 * 5 + 6",
			want:  71,
		},
		{
			input: "1 + (2 * 3) + (4 * (5 + 6))",
			want:  51,
		},
		{
			input: "2 * 3 + (4 * 5)",
			want:  26,
		},
		{
			input: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want:  437,
		},
		{
			input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want:  12240,
		},
		{
			input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want:  13632,
		},
	}

	for i, test := range tests {
		output := evaluateEquationPartOne(test.input)

		if output != test.want {
			t.Errorf("Part1: Wrong output in test case %v; Got %v, expected, %v", i, output, test.want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "1 + 2 * 3 + 4 * 5 + 6",
			want:  231,
		},
		{
			input: "1 + (2 * 3) + (4 * (5 + 6))",
			want:  51,
		},
		{
			input: "2 * 3 + (4 * 5)",
			want:  46,
		},
		{
			input: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want:  1445,
		},
		{
			input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want:  669060,
		},
		{
			input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want:  23340,
		},
	}

	for i, test := range tests {
		output := evaluateEquationPartTwo(test.input)

		if output != test.want {
			t.Errorf("Part2: Wrong output in test case %v; Got %v, expected, %v", i, output, test.want)
		}
	}
}
