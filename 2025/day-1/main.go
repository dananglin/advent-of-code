package main

import (
	"fmt"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/common"
)

func main() {
	filename := "2025/day-1/files/input"

	data, err := common.ReadFile(filename)
	if err != nil {
		common.Fatal(
			"Unable to read the data from "+filename,
			err,
		)
	}

	partOneSecretEntranceCode, err := calculateSecretEntranceCode(data)
	if err != nil {
		common.Fatal(
			"Unable to calculate the code to the secret entrance",
			err,
		)
	}

	fmt.Printf("[Part 1] Code to the secret entrance: %d\n", partOneSecretEntranceCode)
}
