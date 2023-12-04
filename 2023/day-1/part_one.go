package main

import (
	"bufio"
	"fmt"
	"os"
)

func partOneCalculateSumCalibrationValues(document []string) int {
	sum := 0

	for _, line := range document {
		sum = sum + calibrationValue(line)
	}

	return sum
}

func calibrationValue(line string) int {
	var digits []int

	for _, r := range []rune(line) {
		if (r >= '0') && (r <= '9') {
			digits = append(digits, int(r-'0'))
		}
	}

	if len(digits) == 0 {
		return 0
	}

	return (10 * digits[0]) + digits[len(digits)-1]
}

func calibrationDocument(filename string) ([]string, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q; %w", filename, err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning %q; %w", filename, err)
	}

	return lines, nil
}
