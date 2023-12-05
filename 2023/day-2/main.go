package main

import (
	"fmt"
	"log"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/common"
)

func main() {
	filename := "2023/day-2/files/input"

	games, err := common.ReadFile(filename)
	if err != nil {
		log.Fatalf("ERROR: unable to read from %s, %v\n", filename, err)
	}

	sumOfIDs, err := sumOfGameIDs(games)
	if err != nil {
		log.Fatalf("ERROR: unable to calculate the sum of the IDs of the valid games; %v\n", err)
	}

	fmt.Printf("Total sum of the IDs of the valid games: %d\n", sumOfIDs)

	sumOfPower, err := sumOfThePowerOfTheCubes(games)
	if err != nil {
		log.Fatalf("ERROR: unable to calculate the sum of the power of cubes; %v\n", err)
	}

	fmt.Printf("Total sum of the power of cubes: %d\n", sumOfPower)
}

// sumOfGameIDs solves Part 1
func sumOfGameIDs(content []string) (int, error) {
	sum := 0

	for i := range content {
		if len(content[i]) == 0 {
			continue
		}

		game, err := parseGame(content[i])
		if err != nil {
			return 0, fmt.Errorf("unable to parse %q; %w", game, err)
		}

		if validGame(game.Sets) {
			sum = sum + game.ID
		}

	}

	return sum, nil
}

// sumOfThePowerOfTheCubes solves Part 2
func sumOfThePowerOfTheCubes(content []string) (int, error) {
	sum := 0

	for i := range content {
		if len(content[i]) == 0 {
			continue
		}

		game, err := parseGame(content[i])
		if err != nil {
			return 0, fmt.Errorf("unable to parse %q; %w", game, err)
		}

		fewestCubeSet := fewestSet(game.Sets)

		sum = sum + (fewestCubeSet.Red * fewestCubeSet.Green * fewestCubeSet.Blue)
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
