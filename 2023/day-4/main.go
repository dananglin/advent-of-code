package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"codeflow.dananglin.me.uk/apollo/advent-of-code/internal/utilities"
)

type cardResult struct {
	numMatches int
	cardScore  int
}

type solution struct {
	sumUniqueScratchcardScores int
	totalNumberOfScratchcards  int
}

func main() {
	cards, err := utilities.ReadFile("2023/day-4/files/input")
	if err != nil {
		log.Fatalf("ERROR: Unable to retrieve the pack of scratchcards; %v\n", err)
	}

	answer, err := solveScratchCardWinnings(cards)
	if err != nil {
		log.Fatalf("ERROR: Unable to calculate the sum of all scratchcard points; %v\n", err)
	}

	fmt.Printf("[Part 1] The sum of all individual scratchcard scores: %d\n", answer.sumUniqueScratchcardScores)
	fmt.Printf("[Part 2] The total number of scratchcards: %d\n", answer.totalNumberOfScratchcards)
}

func solveScratchCardWinnings(scratchcards []string) (solution, error) {
	sumCardScores := 0
	results := make([]cardResult, len(scratchcards))

	var err error

	// Calculate the total sum of the scores of each unique scratchcard.
	for i, card := range scratchcards {
		results[i], err = cardScore(card)
		if err != nil {
			return solution{}, fmt.Errorf("unable to calculate the card score for %q, %w", card, err)
		}

		sumCardScores = sumCardScores + results[i].cardScore
	}

	// Calculate the total number of unique and duplicated scratchcards.
	cardCopies := make(map[int]int, len(scratchcards))
	sumCardCopies := 0

	for i := 0; i < len(scratchcards); i++ {
		cardCopies[i] = 1
	}

	for i := 0; i < len(scratchcards); i++ {
		for j := i + 1; j <= i+results[i].numMatches; j++ {
			cardCopies[j] = cardCopies[j] + cardCopies[i]
		}
		sumCardCopies = sumCardCopies + cardCopies[i]
	}

	answer := solution{
		sumUniqueScratchcardScores: sumCardScores,
		totalNumberOfScratchcards:  sumCardCopies,
	}

	return answer, nil
}

func cardScore(card string) (cardResult, error) {
	colonIndex := strings.Index(card, ":")
	if colonIndex == -1 {
		return cardResult{}, fmt.Errorf("unable to find the colon (:) in %q", card)
	}

	pipeIndex := strings.Index(card, "|")
	if pipeIndex == -1 {
		return cardResult{}, fmt.Errorf("unable to find the pipe (|) in %q", card)
	}

	winningNumbers := card[colonIndex+1 : pipeIndex]
	playedNumbers := card[pipeIndex+1:]

	// create a map of all winning numbers
	winnningDigit := ""
	cardMap := make(map[string]int)

	for _, d := range winningNumbers {
		if unicode.IsDigit(d) {
			winnningDigit = winnningDigit + string(d)
		} else {
			if len(winnningDigit) > 0 {
				cardMap[winnningDigit] = 0
			}

			winnningDigit = ""
		}
	}

	// record the matching cards to the card map
	playedDigit := ""

	updateCardMap := func(digit string) {
		if len(digit) > 0 {
			if _, ok := cardMap[digit]; ok {
				cardMap[digit] = cardMap[digit] + 1
			}
		}
	}

	for i, d := range playedNumbers {
		if unicode.IsDigit(d) {
			playedDigit = playedDigit + string(d)
			if i == len(playedNumbers)-1 {
				updateCardMap(playedDigit)
			}
		} else {
			updateCardMap(playedDigit)
			playedDigit = ""
		}
	}

	// count the matches
	matches := 0

	for _, v := range cardMap {
		if v == 1 {
			matches++
		}
	}

	// calculate the score
	score := 0

	if matches > 0 {
		score = 1 << (matches - 1)
	}

	result := cardResult{
		numMatches: matches,
		cardScore:  score,
	}

	return result, nil
}
