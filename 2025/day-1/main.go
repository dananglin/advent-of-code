package main

import (
	"fmt"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/utilities"
)

func main() {
	filename := "2025/day-1/files/input"

	data, err := utilities.ReadFile(filename)
	if err != nil {
		utilities.Fatal(
			"Unable to read the data from "+filename,
			err,
		)
	}

	secretEntranceCode, err := calculateSecretEntranceCode(data)
	if err != nil {
		utilities.Fatal(
			"Unable to calculate the code to the secret entrance",
			err,
		)
	}

	fmt.Printf("The code to the secret entrance: %d\n", secretEntranceCode)
}
