package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const (
	dot           = '.'
	potentialGear = '*'
)

type answer struct {
	sumOfValidEnginePartNumbers int
	sumOfGearRatios             int
}

func main() {
	schematic, err := engineSchematic("2023/day-3/files/input")
	if err != nil {
		log.Fatalf("ERROR: Unable to retrieve the engine schematic; %v\n", err)
	}

	answer, err := fixTheGondola(schematic)
	if err != nil {
		log.Fatalf("ERROR: Unable to calculate the sum of all valid engine part numbers; %v\n", err)
	}

	fmt.Printf("[Part 1] The sum of all valid engine part numbers: %d\n", answer.sumOfValidEnginePartNumbers)
	fmt.Printf("[Part 2] The sum of all Gear Ratios: %d\n", answer.sumOfGearRatios)
}

func fixTheGondola(schematic []string) (answer, error) {
	var (
		partNumSum          = 0
		partNum             = ""
		validPartNum        = false
		maxY                = len(schematic) - 1
		maxX                = len(schematic[0]) - 1
		mapOfPotentialGears = make(map[string][]int)
		potentialGearPos    = ""
	)

	for y := range schematic {
		for x, v := range schematic[y] {
			if unicode.IsDigit(v) {
				partNum = partNum + string(v)
				if !validPartNum {
					switch {
					// check above left
					case ((y - 1) >= 0) && ((x - 1) >= 0) && !unicode.IsDigit(rune(schematic[y-1][x-1])) && (schematic[y-1][x-1] != dot):
						validPartNum = true
						if schematic[y-1][x-1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y-1, x-1)
						}
					// check straight above
					case ((y - 1) >= 0) && !unicode.IsDigit(rune(schematic[y-1][x])) && (schematic[y-1][x] != dot):
						validPartNum = true
						if schematic[y-1][x] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y-1, x)
						}
					// check above right
					case ((y - 1) >= 0) && ((x + 1) <= maxX) && !unicode.IsDigit(rune(schematic[y-1][x+1])) && (schematic[y-1][x+1] != dot):
						validPartNum = true
						if schematic[y-1][x+1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y-1, x+1)
						}
					// check left
					case ((x - 1) >= 0) && !unicode.IsDigit(rune(schematic[y][x-1])) && (schematic[y][x-1] != dot):
						validPartNum = true
						if schematic[y][x-1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y, x-1)
						}
					// check right
					case ((x + 1) <= maxX) && !unicode.IsDigit(rune(schematic[y][x+1])) && (schematic[y][x+1] != dot):
						validPartNum = true
						if schematic[y][x+1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y, x+1)
						}
					// check below left
					case ((y + 1) <= maxY) && ((x - 1) >= 0) && !unicode.IsDigit(rune(schematic[y+1][x-1])) && (schematic[y+1][x-1] != dot):
						validPartNum = true
						if schematic[y+1][x-1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y+1, x-1)
						}
					// check straight below
					case ((y + 1) <= maxY) && !unicode.IsDigit(rune(schematic[y+1][x])) && (schematic[y+1][x] != dot):
						validPartNum = true
						if schematic[y+1][x] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y+1, x)
						}
					// check below right
					case ((y + 1) <= maxY) && ((x + 1) <= maxX) && !unicode.IsDigit(rune(schematic[y+1][x+1])) && (schematic[y+1][x+1] != dot):
						validPartNum = true
						if schematic[y+1][x+1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y+1, x+1)
						}
					}
				}
			} else {
				if validPartNum {
					value, err := strconv.Atoi(partNum)
					if err != nil {
						return answer{}, fmt.Errorf("unable to convert %q to int; %w", partNum, err)
					}
					partNumSum = partNumSum + value

					if len(potentialGearPos) > 0 {
						mapOfPotentialGears[potentialGearPos] = append(mapOfPotentialGears[potentialGearPos], value)
					}
				}

				partNum = ""
				validPartNum = false
				potentialGearPos = ""
			}
		}
	}

	// Calculate the sum of all gear ratios
	sumGearRatios := 0

	for _, v := range mapOfPotentialGears {
		if len(v) == 2 {
			sumGearRatios = sumGearRatios + (v[0] * v[1])
		}
	}

	ans := answer{
		sumOfValidEnginePartNumbers: partNumSum,
		sumOfGearRatios:             sumGearRatios,
	}

	return ans, nil
}

func engineSchematic(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q; %w", filename, err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
