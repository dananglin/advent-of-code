package main

import (
	"fmt"
	"log"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/common"
)

func main() {
	document, err := common.ReadFile("2023/day-1/files/input")
	if err != nil {
		log.Fatalf("Error: Unable to retrieve the calibration document; %v", err)
	}

	fmt.Printf("[Part 1] Total sum of calibration values: %d\n", partOneCalculateSumCalibrationValues(document))
	fmt.Printf("[Part 2] Total sum of calibration values: %d\n", partTwoCalculateSumCalibrationValues(document))
}
