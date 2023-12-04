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

type solution struct {
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

func fixTheGondola(schematic []string) (solution, error) {
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
					case ((y - 1) >= 0) && ((x - 1) >= 0) && isSymbol(rune(schematic[y-1][x-1])):
						validPartNum = true
						if schematic[y-1][x-1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y-1, x-1)
						}
					// check straight above
					case ((y - 1) >= 0) && isSymbol(rune(schematic[y-1][x])):
						validPartNum = true
						if schematic[y-1][x] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y-1, x)
						}
					// check above right
					case ((y - 1) >= 0) && ((x + 1) <= maxX) && isSymbol(rune(schematic[y-1][x+1])):
						validPartNum = true
						if schematic[y-1][x+1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y-1, x+1)
						}
					// check left
					case ((x - 1) >= 0) && isSymbol(rune(schematic[y][x-1])):
						validPartNum = true
						if schematic[y][x-1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y, x-1)
						}
					// check right
					case ((x + 1) <= maxX) && isSymbol(rune(schematic[y][x+1])):
						validPartNum = true
						if schematic[y][x+1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y, x+1)
						}
					// check below left
					case ((y + 1) <= maxY) && ((x - 1) >= 0) && isSymbol(rune(schematic[y+1][x-1])):
						validPartNum = true
						if schematic[y+1][x-1] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y+1, x-1)
						}
					// check straight below
					case ((y + 1) <= maxY) && isSymbol(rune(schematic[y+1][x])):
						validPartNum = true
						if schematic[y+1][x] == potentialGear {
							potentialGearPos = fmt.Sprintf("%d-%d", y+1, x)
						}
					// check below right
					case ((y + 1) <= maxY) && ((x + 1) <= maxX) && isSymbol(rune(schematic[y+1][x+1])):
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
						return solution{}, fmt.Errorf("unable to convert %q to int; %w", partNum, err)
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

	answer := solution{
		sumOfValidEnginePartNumbers: partNumSum,
		sumOfGearRatios:             sumGearRatios,
	}

	return answer, nil
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

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && (r != dot)
}
