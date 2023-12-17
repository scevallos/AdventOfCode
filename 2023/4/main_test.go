package main

import (
	helpers "AdventOfCode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumAllScratchcard(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 13, GetSumAllScratchcard(doc))
}

func TestGetScratchcardValue(t *testing.T) {
	cases := []struct{
		name string
		line string
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