package main

import (
	"strings"
	"testing"
)

func TestDay1Trebuchet(t *testing.T) {
	t.Logf("This is the test suite for Advent of Code - Day 1 - Trebuchet?!")
	t.Run("Test the solution for Part 1", testPartOneCalculateSumCalibrationValues)
	t.Run("Test the solution for Part 2", testPartTwoCalculateSumCalibrationValues)
}

func testPartOneCalculateSumCalibrationValues(t *testing.T) {
	testStr := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	document := strings.Split(testStr, "\n")
	got := partOneCalculateSumCalibrationValues(document)

	want := 142

	if want != got {
		t.Errorf("Unexpected value returned from partOneCalculateSumCalibrationValues(); want %d, got %d", want, got)
	}
}

func testPartTwoCalculateSumCalibrationValues(t *testing.T) {
	testStr := `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	document := strings.Split(testStr, "\n")
	got := partTwoCalculateSumCalibrationValues(document)

	want := 281

	if want != got {
		t.Errorf("Unexpected value returned from calculateSumCalibrationValues(); want %d, got %d", want, got)
	}
}
