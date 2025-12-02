package main

import (
	"fmt"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/utilities"
)

func main() {
	filename := "2025/day-2/files/input"

	data, err := utilities.ReadFile(filename)
	if err != nil {
		utilities.Fatal(
			fmt.Sprintf("Error reading the test data from %s", filename),
			err,
		)
	}

	if len(data) != 1 {
		utilities.Fatal(
			fmt.Sprintf(
				"Unexpected number of lines from %s: want 1, got %d",
				filename,
				len(data),
			),
			nil,
		)
	}

	sum, err := sumInvalidProductIDs(data[0])
	if err != nil {
		utilities.Fatal(
			"Error calculating the sum of invalid product IDs",
			err,
		)
	}

	fmt.Printf("Sum total of invalid product IDs: %d\n", sum)
}
