package main

import "testing"

func TestGetSeat1(t *testing.T) {
	wantRow, wantCol := 44, 5

	row, col := getSeat("FBFBBFFRLR")

	if row != wantRow || col != wantCol {
		t.Errorf("getSeat(FBFBBFFRLR) = (%v, %v), expected (%v, %v)", row, col, wantRow, wantCol)
	}
}

func TestGetSeat2(t *testing.T) {
	wantRow, wantCol := 70, 7

	row, col := getSeat("BFFFBBFRRR")

	if row != wantRow || col != wantCol {
		t.Errorf("getSeat(BFFFBBFRRR) = (%v, %v), expected (%v, %v)", row, col, wantRow, wantCol)
	}
}

func TestGetSeat3(t *testing.T) {
	wantRow, wantCol := 14, 7

	row, col := getSeat("FFFBBBFRRR")

	if row != wantRow || col != wantCol {
		t.Errorf("getSeat(FFFBBBFRRR) = (%v, %v), expected (%v, %v)", row, col, wantRow, wantCol)
	}
}

func TestGetSeat4(t *testing.T) {
	wantRow, wantCol := 102, 4

	row, col := getSeat("BBFFBBFRLL")

	if row != wantRow || col != wantCol {
		t.Errorf("getSeat(BBFFBBFRLL) = (%v, %v), expected (%v, %v)", row, col, wantRow, wantCol)
	}
}
