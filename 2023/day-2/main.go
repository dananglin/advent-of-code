package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "2023/day-2/files/input"

	file1, err := os.Open(filename)
	if err != nil {
		log.Fatalf("ERROR: unable to open %s, %v\n", filename, err)
	}
	defer file1.Close()

	sumOfIDs, err := sumOfGameIDs(file1)
	if err != nil {
		log.Fatalf("ERROR: unable to calculate the sum of the IDs of the valid games; %v\n", err)
	}

	fmt.Printf("Total sum of the IDs of the valid games: %d\n", sumOfIDs)

	file2, err := os.Open(filename)
	if err != nil {
		log.Fatalf("ERROR: unable to open %s, %v\n", filename, err)
	}
	defer file2.Close()

	sumOfPower, err := sumOfThePowerOfTheCubes(file2)
	if err != nil {
		log.Fatalf("ERROR: unable to calculate the sum of the power of cubes; %v\n", err)
	}

	fmt.Printf("Total sum of the power of cubes: %d\n", sumOfPower)
}

// sumOfGameIDs solves Part 1
func sumOfGameIDs(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		game, err := parseGame(line)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %q; %w", game, err)
		}

		if validGame(game.Sets) {
			sum = sum + game.ID
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("received an error after scanning file; %w", err)
	}

	return sum, nil
}

// sumOfThePowerOfTheCubes solves Part 2
func sumOfThePowerOfTheCubes(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		game, err := parseGame(line)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %q; %w", game, err)
		}

		fewestCubeSet := fewestSet(game.Sets)

		sum = sum + (fewestCubeSet.Red * fewestCubeSet.Green * fewestCubeSet.Blue)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("received an error after scanning file; %w", err)
	}

	return sum, nil
}

func validGame(sets []Set) bool {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	for _, set := range sets {
		switch {
		case set.Red > maxRed:
			return false
		case set.Blue > maxBlue:
			return false
		case set.Green > maxGreen:
			return false
		}
	}

	return true
}

func fewestSet(sets []Set) Set {
	result := Set{0, 0, 0}

	for _, set := range sets {
		if set.Red > result.Red {
			result.Red = set.Red
		}

		if set.Green > result.Green {
			result.Green = set.Green
		}

		if set.Blue > result.Blue {
			result.Blue = set.Blue
		}
	}

	return result
}
