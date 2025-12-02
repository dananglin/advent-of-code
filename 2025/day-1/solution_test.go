package main

import (
	"testing"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/utilities"
)

func TestDay1SecretEntrance(t *testing.T) {
	t.Log("This is the test suite for the [2025] Advent of Code - Day 1 - Secret Entrance")
	t.Run("Test Calculate Secret Entrance Code", testCalculateSecretEntranceCode)
}

func testCalculateSecretEntranceCode(t *testing.T) {
	filename := "files/test"

	data, err := utilities.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading the test data from %s: %v", filename, err)
	}

	got, err := calculateSecretEntranceCode(data)
	if err != nil {
		t.Fatalf("Error calculating the code to the secret entrance: %v", err)
	}

	want := 6

	if want != got {
		t.Errorf(
			"Unexpected code received from the calculation:\nwant: %d\n got: %d",
			want,
			got,
		)
	} else {
		t.Logf(
			"Expected code received from the calculation:\ngot: %d",
			got,
		)
	}
}
