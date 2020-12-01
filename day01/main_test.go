package main

import "testing"

func TestDummyFunc(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 0,
			want:  0,
		},
	}

	for _, test := range tests {
		if output := 0; output != test.want {
			t.Errorf("mydummyfunc(%d) = %v; want %d", test.input, output, test.want)
		}
	}
}
