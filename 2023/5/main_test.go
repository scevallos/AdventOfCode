package main

import (
	"bufio"
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

func TestGetLowestLocationNumberForSeeds(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 35, GetLowestLocationNumberForSeeds(doc))
}

func TestProcessMap(t *testing.T) {
	cases := []struct {
		name         string
		dsts         []int
		srcs         []int
		rangeLengths []int
		inputSrc     int
		expected     int
	}{
		{
			name:         "sample seed-to-soil first seed",
			dsts:         []int{50, 52},
			srcs:         []int{98, 50},
			rangeLengths: []int{2, 48},
			inputSrc:     79,
			expected:     81,
		},
		{
			name:         "sample seed-to-soil second seed",
			dsts:         []int{50, 52},
			srcs:         []int{98, 50},
			rangeLengths: []int{2, 48},
			inputSrc:     14,
			expected:     14,
		},
		{
			name:         "sample seed-to-soil third seed",
			dsts:         []int{50, 52},
			srcs:         []int{98, 50},
			rangeLengths: []int{2, 48},
			inputSrc:     55,
			expected:     57,
		},
		{
			name:         "sample seed-to-soil fourth seed",
			dsts:         []int{50, 52},
			srcs:         []int{98, 50},
			rangeLengths: []int{2, 48},
			inputSrc:     13,
			expected:     13,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, processMap(tc.dsts, tc.srcs, tc.rangeLengths, tc.inputSrc))
		})
	}
}

func TestParseInput(t *testing.T) {
	cases := []struct {
		name         string
		filename     string
		inputStr     string
		expectedMaps [][][]int
	}{
		{
			name: "first map sample input",
			inputStr: `seeds: 79 14 55 13

			seed-to-soil map:
			50 98 2
			52 50 48
			`,
			expectedMaps: [][][]int{
				{
					{50, 52},
					{98, 50},
					{2, 48},
				},
			},
		},
		{
			name:     "sample input",
			filename: "sampleInput.txt",
			expectedMaps: [][][]int{
				{
					{50, 52},
					{98, 50},
					{2, 48},
				},
				{
					{0, 37, 39},
					{15, 52, 0},
					{37, 2, 15},
				},
				{
					{49, 0, 42, 57},
					{53, 11, 0, 7},
					{8, 42, 7, 4},
				},
				{
					{88, 18},
					{18, 25},
					{7, 70},
				},
				{
					{45, 81, 68},
					{77, 45, 64},
					{23, 19, 13},
				},
				{
					{0, 1},
					{69, 0},
					{1, 69},
				},
				{
					{60, 56},
					{56, 93},
					{37, 4},
				},
			},
		},
	}

	for _, tc := range cases {
		var doc *bufio.Scanner
		if tc.filename != "" {
			var closeFile func() error
			doc, closeFile = helpers.GetDocFromFile(tc.filename)
			defer closeFile()
		} else {
			doc = helpers.GetDocFromString(tc.inputStr)
		}
		_, allTheMaps := parseInput(doc)
		assert.Equal(t, tc.expectedMaps, allTheMaps)
	}
}
