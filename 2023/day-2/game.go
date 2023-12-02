package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func parseGame(line string) (Game, error) {
	splits := strings.SplitN(strings.ReplaceAll(line, " ", ""), ":", 2)

	if len(splits) != 2 {
		return Game{}, fmt.Errorf("unexpected number of splits from %s; want 2; got %d", splits, len(splits))
	}

	gameStr, setsStr := splits[0], splits[1]

	gameID, err := parseGameID(gameStr)
	if err != nil {
		return Game{}, fmt.Errorf("unable to parse the Game ID; %w", err)
	}

	sets, err := parseSets(setsStr)
	if err != nil {
		return Game{}, fmt.Errorf("unable to parse the sets; %w", err)
	}

	game := Game{
		ID:   gameID,
		Sets: sets,
	}

	return game, nil
}

func parseGameID(game string) (int, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(game, "Game"))
	if err != nil {
		return 0, fmt.Errorf("error running strconv.Atoi(); %w", err)
	}

	return id, nil
}

func parseSets(str string) ([]Set, error) {
	setList := strings.Split(str, ";")
	var sets []Set

	const (
		red   = "red"
		green = "green"
		blue  = "blue"
	)

	for _, s := range setList {
		set := Set{0, 0, 0}
		split := strings.Split(s, ",")
		for _, i := range split {
			switch {
			case strings.HasSuffix(i, red):
				value, err := strconv.Atoi(strings.TrimSuffix(i, red))
				if err != nil {
					return nil, fmt.Errorf("unable to run strconv.Atoi() on %s; %w", i, err)
				}
				set.Red = value
			case strings.HasSuffix(i, green):
				value, err := strconv.Atoi(strings.TrimSuffix(i, green))
				if err != nil {
					return nil, fmt.Errorf("unable to run strconv.Atoi() on %s; %w", i, err)
				}
				set.Green = value
			case strings.HasSuffix(i, "blue"):
				value, err := strconv.Atoi(strings.TrimSuffix(i, blue))
				if err != nil {
					return nil, fmt.Errorf("unable to run strconv.Atoi() on %s; %w", i, err)
				}
				set.Blue = value
			}
		}
		sets = append(sets, set)
	}

	return sets, nil
}
