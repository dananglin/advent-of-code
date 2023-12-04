package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay4Scratchcards(t *testing.T) {
	t.Log("This is the test suite for Advent of Code - Day 4 - Scratchcards")

	testScratchcard := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

	scratchcards := strings.Split(testScratchcard, "\n")

	// remove the first and last lines of the test card pack as they are blank lines
	scratchcards = scratchcards[1 : len(scratchcards)-1]

	t.Run("Test the solution of the scratch card winnings", testSolveScratchCardWinnings(scratchcards))
}

func testSolveScratchCardWinnings(scratchcards []string) func(t *testing.T) {
	return func(t *testing.T) {
		got, err := solveScratchCardWinnings(scratchcards)
		if err != nil {
			t.Fatalf("Received an error after running solveScratchCardWinnings(); %v\n", err)
		}

		want := solution{
			sumUniqueScratchcardScores: 13,
			totalNumberOfScratchcards:  30,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("unexpected result received from solveScratchCardWinnings(); want %d, got %d\n", want, got)
		}
	}
}
