package main

import (
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

func TestGetSumAllScratchcard(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 13, GetSumAllScratchcard(doc))
}

func TestGetScratchcardValue(t *testing.T) {
	cases := []struct {
		name     string
		line     string
		expected int
	}{
		{
			name:     "card 1 sample",
			line:     "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expected: 8,
		},
		{
			name:     "card 3 sample",
			line:     "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, GetScratchcardValue(tc.line))
		})
	}
}

func TestGetCardNumberAndMatches(t *testing.T) {
	cases := []struct {
		name               string
		line               string
		expectedCardNumber int
		expectedMatches    int
	}{
		{
			name:               "card 1 sample",
			line:               "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expectedCardNumber: 1,
			expectedMatches:    4,
		},
		{
			name:               "card 3 sample",
			line:               "Card 33:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expectedCardNumber: 33,
			expectedMatches:    2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cardNumber, matches := getCardNumberAndMatches(tc.line)
			assert.Equal(t, tc.expectedCardNumber, cardNumber)
			assert.Equal(t, tc.expectedMatches, matches)
		})
	}
}

func TestProcessCard(t *testing.T) {
	scratchCardMatches = map[int]int{
		1: 4,
		2: 2,
		3: 2,
		4: 1,
		5: 0,
		6: 0,
	}
	cases := []struct {
		name     string
		cardNum  int
		expected int
	}{
		{
			name:     "card 5 sample",
			cardNum:  5,
			expected: 1,
		},
		{
			name:     "card 4 sample",
			cardNum:  4,
			expected: 2,
		},
		{
			name:     "card 3 sample",
			cardNum:  3,
			expected: 4,
		},
		{
			name:     "card 2 sample",
			cardNum:  2,
			expected: 7,
		},
		{
			name:     "card 1 sample",
			cardNum:  1,
			expected: 15,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, processCard(tc.cardNum))
		})
	}
}

func TestGetTotalScratchcard(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 13, GetSumAllScratchcard(doc))
}
