package main

import (
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

func partTwoCalculateSumCalibrationValues(document []string) int {
	sum := 0

	for _, line := range document {
		sum = sum + partTwoCalibrationValue(line)
	}

	return sum
}

func partTwoCalibrationValue(line string) int {
	var digits []int

	for len(line) > 0 {
		switch {
		case (line[0] >= '0') && (line[0] <= '9'):
			digits = append(digits, int(line[0]-'0'))
			line = line[1:]
		case strings.HasPrefix(line, one):
			digits = append(digits, 1)
			trimLength := len(one) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, two):
			digits = append(digits, 2)
			trimLength := len(two) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, three):
			digits = append(digits, 3)
			trimLength := len(three) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, four):
			digits = append(digits, 4)
			trimLength := len(four) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, five):
			digits = append(digits, 5)
			trimLength := len(five) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, six):
			digits = append(digits, 6)
			trimLength := len(six) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, seven):
			digits = append(digits, 7)
			trimLength := len(seven) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, eight):
			digits = append(digits, 8)
			trimLength := len(eight) - 1
			line = line[trimLength:]
		case strings.HasPrefix(line, nine):
			digits = append(digits, 9)
			trimLength := len(nine) - 1
			line = line[trimLength:]
		default:
			line = line[1:]
		}
	}

	if len(digits) == 0 {
		return 0
	}

	return (10*digits[0] + digits[len(digits)-1])
}
