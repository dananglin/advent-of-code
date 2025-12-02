package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q; %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning %q; %w", filename, err)
	}

	return lines, nil
}

func ParseIntList(line string) ([]int, error) {
	var values []int
	digitString := ""

	addValue := func(digit string) error {
		if len(digit) > 0 {
			value, err := strconv.Atoi(digit)
			if err != nil {
				return fmt.Errorf("unable to parse %q to int; %w", digit, err)
			}
			values = append(values, value)
		}

		return nil
	}

	for i, v := range line {
		if unicode.IsDigit(v) {
			digitString = digitString + string(v)
			if i == len(line)-1 {
				if err := addValue(digitString); err != nil {
					return nil, fmt.Errorf("unable to add %q to the list of ints; %w", digitString, err)
				}
			}
		} else {
			if err := addValue(digitString); err != nil {
				return nil, fmt.Errorf("unable to add %q to the list of ints; %w", digitString, err)
			}
			digitString = ""
		}
	}

	return values, nil
}

func Fatal(msg string, err error) {
	fmt.Printf("ERROR: %s: %v.\n", msg, err)

	os.Exit(1)
}
