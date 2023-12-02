package main

import (
	"reflect"
	"testing"
)

func TestParseGame(t *testing.T) {
	table := []struct {
		testGame string
		want     Game
	}{
		{
			testGame: "Game1: 10 green, 1 blue, 7red",
			want: Game{
				ID: 1,
				Sets: []Set{
					{
						Red:   7,
						Green: 10,
						Blue:  1,
					},
				},
			},
		},
		{
			testGame: "Game 100: 8 green, 6 blue, 20red; 5 blue, 4 red, 13green; 5 green, 1 red",
			want: Game{
				ID: 100,
				Sets: []Set{
					{
						Red:   20,
						Green: 8,
						Blue:  6,
					},
					{
						Red:   4,
						Green: 13,
						Blue:  5,
					},
					{
						Red:   1,
						Green: 5,
						Blue:  0,
					},
				},
			},
		},
	}

	for i := range table {
		got, err := parseGame(table[i].testGame)
		if err != nil {
			t.Fatalf("Error running parseGame(); %v", err)
		}

		if !reflect.DeepEqual(table[i].want, got) {
			t.Errorf("unexpected result after parsing the game; want %+v, got %+v", table[i].want, got)
		}

	}
}
