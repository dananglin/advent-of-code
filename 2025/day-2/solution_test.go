package main

import (
	"testing"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/utilities"
)

func TestDay2GiftShop(t *testing.T) {
	t.Log("This is the test suite for the [2025] Advent of Code - Day 2 - Gift Shop")
	t.Run("Test the solution for Part 1", testSumInvalidIDs)
}

func testSumInvalidIDs(t *testing.T) {
	filename := "files/test"

	data, err := utilities.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading the test data from %s: %v", filename, err)
	}

	if len(data) != 1 {
		t.Fatalf(
			"Unexpected number of lines from %s: want 1, got %d",
			filename,
			len(data),
		)
	}

	got, err := sumInvalidProductIDs(data[0])
	if err != nil {
		t.Fatalf(
			"Error calculating the sum of invalid product IDs: %v",
			err,
		)
	}

	want := 4174379265

	if want != got {
		t.Errorf(
			"Unexpected sum of invalid product IDs found:\nwant: %d\n got: %d",
			want,
			got,
		)
	} else {
		t.Logf(
			"Expected sum of invalid product IDs found:\ngot: %d",
			got,
		)
	}
}
