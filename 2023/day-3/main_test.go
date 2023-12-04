package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay3GearRatios(t *testing.T) {
	t.Logf("This is the test suite for Advent of Code - Day 3 - Gear Ratio")

	testEngineSchematic := `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`
	schematic := strings.Split(testEngineSchematic, "\n")

	// remove the first and last lines of the test schematic as they are blank.
	schematic = schematic[1 : len(schematic)-1]

	t.Run("Test the fix to the Gondola issue", testFixTheGondola(schematic))
}

func testFixTheGondola(schematic []string) func(t *testing.T) {
	return func(t *testing.T) {
		got, err := fixTheGondola(schematic)
		if err != nil {
			t.Fatalf("Received an error after running sumOfValidEnginePartNumbers(); %v\n", err)
		}

		want := solution{
			sumOfValidEnginePartNumbers: 4361,
			sumOfGearRatios: 467835,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("unexpected result received from sumOfValidEnginePartNumbers(); want %d, got %d\n", want, got)
		}
	}
}
