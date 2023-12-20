package main

import (
	"fmt"
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

func TestProcessMatrix(t *testing.T) {
	assert.Equal(t, 199, ProcessMatrix([][]rune{
		{'4', '.', '.', '2', '1'}, // 4 + 21
		{'.', '#', '.', '*', '.'}, // -
		{'.', '.', '1', '7', '3'}, // 173
		{'1', '.', '.', '5', '6'}, // 1 (skip 5 6)
		{'$', '.', '5', '.', '.'}, // (skip 5)
		// total 199
	}, PartNumberSum))

	assert.Equal(t, 3633, ProcessMatrix([][]rune{
		{'4', '.', '.', '2', '1'}, // 21
		{'.', '#', '.', '*', '.'}, // -
		{'.', '.', '1', '7', '3'}, // 173
		{'1', '.', '.', '5', '6'}, //
		{'$', '.', '5', '.', '.'}, //
		// total 3633
	}, GearRatio))
}

func TestAbsorbNumber(t *testing.T) {
	cases := []struct {
		name       string
		inputChars []rune
		inputIndex int
		expected   string
		expIndices []int
	}{
		{
			name:       "1 char",
			inputChars: []rune{'1'},
			inputIndex: 0,
			expected:   "1",
			expIndices: []int{0},
		},
		{
			name:       "1 dot",
			inputChars: []rune{'.'},
			inputIndex: 0,
			expected:   "",
			expIndices: []int{},
		},
		{
			name:       "3 chars",
			inputChars: []rune{'1', '2', '3'},
			inputIndex: 2,
			expected:   "123",
			expIndices: []int{0, 1, 2},
		},
		{
			name:       "first row end of first",
			inputChars: []rune{'4', '6', '7', '.', '.', '1', '1', '4'},
			inputIndex: 2,
			expected:   "467",
			expIndices: []int{0, 1, 2},
		},
		{
			name:       "first row start of last",
			inputChars: []rune{'4', '6', '7', '.', '.', '1', '1', '4'},
			inputIndex: 5,
			expected:   "114",
			expIndices: []int{5, 6, 7},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			number, indices := AbsorbNumber(tc.inputChars, tc.inputIndex)
			assert.Equal(t, tc.expected, number)
			assert.Equal(t, tc.expIndices, indices)
		})
	}
}

func TestGetSumAllPartNumbers(t *testing.T) {
	// % = & $ * # @ + - /
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()
	fmt.Println(GetSumAllPartNumbers(doc))
}

func TestGetSumAllGearRatios(t *testing.T) {
	// % = & $ * # @ + - /
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()
	fmt.Println(GetSumAllGearRatios(doc))
}
