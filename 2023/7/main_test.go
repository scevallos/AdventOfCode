package main

import (
	"fmt"
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalWinnings(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 6440, GetTotalWinnings(doc))
}

func TestParseInput(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	players := parseInput(doc)
	fmt.Println("OrderPlayers")
	OrderPlayers(players)
	fmt.Println(players)
}

func TestOrderPlayers(t *testing.T) {
	cases := []struct {
		name          string
		players       []Player
		expectedOrder []Player
	}{
		{
			name:          "sampleInput",
			players:       []Player{
				{
					Hand: Hand{
						Cards: "QQQJA",
					},
					Bid:  483,
				},
				{
					Hand: Hand{
						Cards: "T55J5",
					},
					Bid:  684,
				},
				{
					Hand: Hand{
						Cards: "KK677",
					},
					Bid:  28,
				},
				{
					Hand: Hand{
						Cards: "KTJJT",
					},
					Bid:  220,
				},
				{
					Hand: Hand{
						Cards: "32T3K",
					},
					Bid:  765,
				},
			},
			expectedOrder: []Player{
				{
					Hand: Hand{
						Cards: "KTJJT",
					},
					Bid:  220,
				},
				{
					Hand: Hand{
						Cards: "QQQJA",
					},
					Bid:  483,
				},
				{
					Hand: Hand{
						Cards: "T55J5",
					},
					Bid:  684,
				},
				{
					Hand: Hand{
						Cards: "KK677",
					},
					Bid:  28,
				},
				{
					Hand: Hand{
						Cards: "32T3K",
					},
					Bid:  765,
				},
			},
		},
		{
			name:          "jokers in play",
			players:       []Player{
				{
					Hand: Hand{
						Cards: "QQQJA",
					},
					Bid:  483,
				},
				{
					Hand: Hand{
						Cards: "T55J5",
					},
					Bid:  684,
				},
			},
			expectedOrder: []Player{
				{
					Hand: Hand{
						Cards: "QQQJA",
					},
					Bid:  483,
				},
				{
					Hand: Hand{
						Cards: "T55J5",
					},
					Bid:  684,
				},
			},
		},
		{
			name:          "jokers in play high card",
			players:       []Player{
				{
					Hand: Hand{
						Cards: "2345J",
					},
				},
				{
					Hand: Hand{
						Cards: "23457",
					},
				},
			},
			expectedOrder: []Player{
				{
					Hand: Hand{
						Cards: "2345J",
					},
				},
				{
					Hand: Hand{
						Cards: "23457",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			OrderPlayers(tc.players)
			assert.Equal(t, tc.expectedOrder, tc.players)
		})
	}
}
