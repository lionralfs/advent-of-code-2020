package main

import "testing"

func TestPartOne1(t *testing.T) {
	want := 694
	output := calc(readInput("./input.txt"), 2020)

	if want != output {
		t.Errorf("calc(input.txt, 2020) = %v, expected %v", output, want)
	}
}

func TestPartOne2(t *testing.T) {
	want := 1
	output := calc([]int{1, 3, 2}, 2020)

	if want != output {
		t.Errorf("calc({1, 3, 2}, 2020) = %v, expected %v", output, want)
	}
}

func TestPartOne3(t *testing.T) {
	want := 10
	output := calc([]int{2, 1, 3}, 2020)

	if want != output {
		t.Errorf("calc({2, 1, 3}, 2020) = %v, expected %v", output, want)
	}
}

func TestPartOne4(t *testing.T) {
	want := 27
	output := calc([]int{1, 2, 3}, 2020)

	if want != output {
		t.Errorf("calc({1, 2, 3}, 2020) = %v, expected %v", output, want)
	}
}

func TestPartOne5(t *testing.T) {
	want := 78
	output := calc([]int{2, 3, 1}, 2020)

	if want != output {
		t.Errorf("calc({2, 3, 1}, 2020) = %v, expected %v", output, want)
	}
}

func TestPartOne6(t *testing.T) {
	want := 438
	output := calc([]int{3, 2, 1}, 2020)

	if want != output {
		t.Errorf("calc({3, 2, 1}, 2020) = %v, expected %v", output, want)
	}
}

func TestPartOne7(t *testing.T) {
	want := 1836
	output := calc([]int{3, 1, 2}, 2020)

	if want != output {
		t.Errorf("calc({3, 1, 2}, 2020) = %v, expected %v", output, want)
	}
}

func TestPartTwo1(t *testing.T) {
	want := 21768614
	output := calc(readInput("./input.txt"), 30000000)

	if want != output {
		t.Errorf("calc(input.txt, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo2(t *testing.T) {
	want := 2578
	output := calc([]int{1, 3, 2}, 30000000)

	if want != output {
		t.Errorf("calc({1, 3, 2}, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo3(t *testing.T) {
	want := 3544142
	output := calc([]int{2, 1, 3}, 30000000)

	if want != output {
		t.Errorf("calc({2, 1, 3}, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo4(t *testing.T) {
	want := 261214
	output := calc([]int{1, 2, 3}, 30000000)

	if want != output {
		t.Errorf("calc({1, 2, 3}, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo5(t *testing.T) {
	want := 6895259
	output := calc([]int{2, 3, 1}, 30000000)

	if want != output {
		t.Errorf("calc({2, 3, 1}, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo6(t *testing.T) {
	want := 18
	output := calc([]int{3, 2, 1}, 30000000)

	if want != output {
		t.Errorf("calc({3, 2, 1}, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo7(t *testing.T) {
	want := 362
	output := calc([]int{3, 1, 2}, 30000000)

	if want != output {
		t.Errorf("calc({3, 1, 2}, 30000000) = %v, expected %v", output, want)
	}
}

func TestPartTwo8(t *testing.T) {
	want := 175594
	output := calc([]int{0, 3, 6}, 30000000)

	if want != output {
		t.Errorf("calc({0, 3, 6}, 30000000) = %v, expected %v", output, want)
	}
}
