package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	startVal int = 50
	minVal   int = 0
	maxVal   int = 99
	maxDist  int = 100
)

type direction int

const (
	left  direction = -1
	right direction = 1
)

type processedSecretEntrance struct {
	dialValue int
	code      int
}

func calculateSecretEntranceCode(data []string) (int, error) {
	var err error

	processed := processedSecretEntrance{
		dialValue: startVal,
		code:      0,
	}

	for idx := range data {
		processed, err = processRotation(processed, data[idx])
		if err != nil {
			return 0, fmt.Errorf(
				"error calculating the new value from %q: %w",
				data[idx],
				err,
			)
		}
	}

	return processed.code, nil
}

func processRotation(processed processedSecretEntrance, data string) (processedSecretEntrance, error) {
	move := right

	if strings.HasPrefix(strings.ToLower(data), "l") {
		move = left
	}

	dist, err := strconv.Atoi(data[1:])
	if err != nil {
		return processedSecretEntrance{dialValue: 0, code: 0},
			fmt.Errorf("error parsing the distance: %w", err)
	}

	for dist > maxDist {
		dist -= maxDist
		processed.code += 1
	}

	oldDialVal := processed.dialValue
	dialVal := processed.dialValue + (dist * int(move))

	switch {
	case dialVal < minVal:
		processed.dialValue = (maxVal + 1) - (minVal - dialVal)
		if oldDialVal != 0 {
			processed.code += 1
		}
	case dialVal > maxVal:
		processed.dialValue = (minVal - 1) + (dialVal - maxVal)
		processed.code += 1
	case dialVal == 0:
		processed.dialValue = dialVal
		processed.code += 1
	default:
		processed.dialValue = dialVal
	}

	return processed, nil
}
