package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const numWorkers int = 4

func sumInvalidProductIDs(data string) (int, error) {
	productIDRanges := strings.Split(data, ",")

	var wg sync.WaitGroup
	wg.Add(len(productIDRanges))

	var mu sync.Mutex
	errs := make([]error, 0)

	sum := 0
	workerPool := make(chan struct{}, numWorkers)

	for idx := range productIDRanges {
		go func() {
			workerPool <- struct{}{}

			defer func() {
				<-workerPool
				wg.Done()
			}()

			lower, upper, err := calculateProductIDRange(productIDRanges[idx])
			if err != nil {
				errs = append(errs, err)

				return
			}

			for id := lower; id <= upper; id++ {
				if productIDInvalid(strconv.Itoa(id)) {
					mu.Lock()
					{
						sum += id
					}
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	if len(errs) > 0 {
		return 0, fmt.Errorf(
			"at least one error occurred while calculating the sum of invalid product IDs: %w",
			errors.Join(errs...),
		)
	}

	return sum, nil
}

func calculateProductIDRange(data string) (int, int, error) {
	strArr := strings.Split(data, "-")
	if len(strArr) != 2 {
		return 0, 0, fmt.Errorf("unexpected number of splits found in product ID range: want 2, got %d", len(strArr))
	}

	lower, err := strconv.Atoi(strArr[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing the lower bound %q: %w", lower, err)
	}

	upper, err := strconv.Atoi(strArr[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing the upper bound %q: %w", upper, err)
	}

	if upper < lower {
		return 0, 0, errors.New("the upper bound (%q) is lower than the lower bound (%q)")
	}

	return lower, upper, nil
}

func productIDInvalid(productID string) bool {
	for subStrLim := 1; subStrLim <= len(productID)/2; subStrLim++ {
		if len(productID)%subStrLim != 0 {
			continue
		}

		repeats := regexp.MustCompile(productID[:subStrLim]).FindAllString(productID, -1)

		if len(repeats) == len(productID)/subStrLim {
			return true
		}
	}

	return false
}
