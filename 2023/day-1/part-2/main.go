package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

func main() {
	file, err := os.Open("2023/day-1/files/input")
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	defer file.Close()

	fmt.Printf("Sum total: %d\n", calculateSumCalibrationValues(file))
}

func calculateSumCalibrationValues(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var (
		digits []int
		sum    int
	)

	for scanner.Scan() {
		line := scanner.Text()
		sum = sum + calibrationValue(line, digits)
	}

	return sum
}

func calibrationValue(line string, digits []int) int {
	if len(line) > 0 {
		switch {
		case (line[0] >= '0') && (line[0] <= '9'):
			digits = append(digits, int(line[0]-'0'))
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, one):
			digits = append(digits, 1)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, two):
			digits = append(digits, 2)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, three):
			digits = append(digits, 3)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, four):
			digits = append(digits, 4)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, five):
			digits = append(digits, 5)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, six):
			digits = append(digits, 6)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, seven):
			digits = append(digits, 7)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, eight):
			digits = append(digits, 8)
			return calibrationValue(line[1:], digits)
		case strings.HasPrefix(line, nine):
			digits = append(digits, 9)
			return calibrationValue(line[1:], digits)
		default:
			return calibrationValue(line[1:], digits)
		}
	}

	if len(digits) == 0 {
		return 0
	}

	return (10*digits[0] + digits[len(digits)-1])
}
