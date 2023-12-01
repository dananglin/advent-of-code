package main

import (
	"bytes"
	"testing"
)

func TestCalculateSumCalibrationValues(t *testing.T) {
	testStr := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	buf := bytes.NewBuffer([]byte(testStr))

	want := 142
	got := calculateSumCalibrationValues(buf)

	if want != got {
		t.Errorf("Unexpected value returned from calculateSumCalibrationValues(); want %d, got %d", want, got)
	}
}
