package main

import (
	"bytes"
	"testing"
)

func TestCalculateSumCalibrationValues(t *testing.T) {
	testStr := `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	buf := bytes.NewBuffer([]byte(testStr))

	want := 281
	got := calculateSumCalibrationValues(buf)

	if want != got {
		t.Errorf("Unexpected value returned from calculateSumCalibrationValues(); want %d, got %d", want, got)
	}
}
