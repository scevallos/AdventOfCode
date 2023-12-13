package main

import (
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

const (
	sampleInputAnswer  = 8
	sampleInputAnswer2 = 2286
	firstLine          = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
)

func TestIsGamePossible(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestGetSumIdsPossibleGames(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()
	assert.Equal(t, sampleInputAnswer, GetSumIdsPossibleGames(doc))
}

func TestGetSumGamePowers(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()
	assert.Equal(t, sampleInputAnswer2, GetSumGamePowers(doc))
}

func TestParseLineToGame(t *testing.T) {
	cases := []struct {
		name         string
		inputLine    string
		expectedGame *Game
	}{
		{
			name:      "first line sample input",
			inputLine: firstLine,
			expectedGame: &Game{
				ID: 1,
				Sets: []*Set{
					{
						BluesDrawn: 3,
						RedsDrawn:  4,
					},
					{
						RedsDrawn:   1,
						GreensDrawn: 2,
						BluesDrawn:  6,
					},
					{
						GreensDrawn: 2,
					},
				},
			},
		},
		{
			name:      "second line sample input",
			inputLine: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expectedGame: &Game{
				ID: 2,
				Sets: []*Set{
					{
						BluesDrawn:  1,
						GreensDrawn: 2,
					},
					{
						GreensDrawn: 3,
						BluesDrawn:  4,
						RedsDrawn:   1,
					},
					{
						GreensDrawn: 1,
						BluesDrawn:  1,
					},
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedGame, ParseLineToGame(tc.inputLine))
		})
	}
}

func TestParseSets(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		expectedSets []*Set
	}{
		{
			name:  "game 1 set sample input",
			input: "8 green, 6 blue, 20 red",
			expectedSets: []*Set{
				{
					RedsDrawn:   20,
					GreensDrawn: 8,
					BluesDrawn:  6,
				},
			},
		},
		{
			name:  "game 1 line sample input",
			input: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expectedSets: []*Set{
				{
					BluesDrawn: 3,
					RedsDrawn:  4,
				},
				{
					RedsDrawn:   1,
					GreensDrawn: 2,
					BluesDrawn:  6,
				},
				{
					GreensDrawn: 2,
				},
			},
		},
		{
			name:  "game 3 line sample input",
			input: "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expectedSets: []*Set{
				{
					GreensDrawn: 8,
					BluesDrawn:  6,
					RedsDrawn:   20,
				},
				{
					BluesDrawn:  5,
					RedsDrawn:   4,
					GreensDrawn: 13,
				},
				{
					GreensDrawn: 5,
					RedsDrawn:   1,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSets, ParseSets(tc.input))
		})
	}
	// fmt.Println(ParseSets("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
}

func TestGetMinCubesNeeded(t *testing.T) {
	cases := []struct {
		name        string
		inputSets   []*Set
		expectedBag *Bag
	}{
		{
			name: "game 3 set 1 sample input",
			inputSets: []*Set{
				{
					RedsDrawn:   20,
					GreensDrawn: 8,
					BluesDrawn:  6,
				},
			},
			expectedBag: &Bag{
				Reds:   20,
				Greens: 8,
				Blues:  6,
			},
		},
		{
			name: "game 1 line sample input",
			inputSets: []*Set{
				{
					BluesDrawn: 3,
					RedsDrawn:  4,
				},
				{
					RedsDrawn:   1,
					GreensDrawn: 2,
					BluesDrawn:  6,
				},
				{
					GreensDrawn: 2,
				},
			},
			expectedBag: &Bag{
				Reds:   4,
				Greens: 2,
				Blues:  6,
			},
		},
		{
			name: "game 3 line sample input",
			inputSets: []*Set{
				{
					GreensDrawn: 8,
					BluesDrawn:  6,
					RedsDrawn:   20,
				},
				{
					BluesDrawn:  5,
					RedsDrawn:   4,
					GreensDrawn: 13,
				},
				{
					GreensDrawn: 5,
					RedsDrawn:   1,
				},
			},
			expectedBag: &Bag{
				Reds:   20,
				Greens: 13,
				Blues:  6,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedBag, GetMinCubesNeeded(tc.inputSets))
		})
	}
	// fmt.Println(ParseSets("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
}

func TestTranslateTokensToSets(t *testing.T) {
	assert.Equal(t,
		[]*Set{
			{
				BluesDrawn: 3,
				RedsDrawn:  4,
			},
			{
				RedsDrawn:   1,
				GreensDrawn: 2,
				BluesDrawn:  6,
			},
			{
				GreensDrawn: 2,
			},
		},
		translateTokensToSets([][]int{{3, 2, 4, 1}, {1, 1, 2, 0, 6, 2}, {2, 0}}),
	)
}
