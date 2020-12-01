package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1721, 979, 366, 299, 675, 1456},
			want:  514579,
		},
	}

	for _, test := range tests {
		if output := twoSummands(test.input); output != test.want {
			t.Errorf("twoSummands(%d) = %v; want %d", test.input, output, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1721, 979, 366, 299, 675, 1456},
			want:  241861950,
		},
	}

	for _, test := range tests {
		if output := threeSummands(test.input); output != test.want {
			t.Errorf("twoSummands(%d) = %v; want %d", test.input, output, test.want)
		}
	}
}
