package main

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
