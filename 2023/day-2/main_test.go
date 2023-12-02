package main

import (
	"bytes"
	"testing"
)

var testStr = `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func TestSumOfGameIDs(t *testing.T) {
	buf := bytes.NewBufferString(testStr)

	got, err := sumOfGameIDs(buf)
	if err != nil {
		t.Fatalf("Received an error after running sumOfGameIDs(); %v\n", err)
	}

	want := 8

	if got != want {
		t.Errorf("unexpected result received from sumOfGameIDs(); want %d, got %d\n", want, got)
	}
}

func TestSumOfThePowerOfTheCubes(t *testing.T) {
	buf := bytes.NewBufferString(testStr)

	got, err := sumOfThePowerOfTheCubes(buf)
	if err != nil {
		t.Fatalf("Received an error after running sumOfThePowerOfTheCubes(); %v\n", err)
	}

	want := 2286

	if got != want {
		t.Errorf("unexpected result received from sumOfThePowerOfTheCubes(); want %d; got %d", want, got)
	}
}
