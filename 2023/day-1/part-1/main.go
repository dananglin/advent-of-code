package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2023/day-1/files/input")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	fmt.Printf("Sum total: %d\n", calculateSumCalibrationValues(file))

}

func calculateSumCalibrationValues(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum = sum + calibrationValue(line)
	}
	
	return sum
}

func calibrationValue(line string) int {
	var digits []int

	for _, r := range []rune(line) {
		if (r >= '0') && (r <= '9') {
			digits = append(digits, int(r - '0'))
		}
	}

	if len(digits) == 0 {
		return 0
	}

	return (10*digits[0]) + digits[len(digits)-1]
}
