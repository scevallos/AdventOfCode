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

func TestGetLowestLocationNumberForSeedRanges(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 46, GetLowestLocationNumberForSeedRanges(doc))
}

func TestGetOverlappingRange(t *testing.T) {
	cases := []struct {
		name          string
		ijxy          []int
		expectedRange []int
	}{
		{
			name:          "right half",
			ijxy:          []int{2, 5, 4, 7},
			expectedRange: []int{4, 5},
		},
		{
			name:          "left half",
			ijxy:          []int{2, 5, 1, 3},
			expectedRange: []int{2, 3},
		},
		{
			name:          "bigger than",
			ijxy:          []int{2, 5, 1, 7},
			expectedRange: []int{2, 5},
		},
		{
			name:          "smaller than",
			ijxy:          []int{2, 5, 3, 4},
			expectedRange: []int{3, 4},
		},
		{
			name:          "no overlap",
			ijxy:          []int{1, 5, 20, 24},
			expectedRange: []int{-1, -1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a, b := getOverlappingRange(tc.ijxy[0], tc.ijxy[1], tc.ijxy[2], tc.ijxy[3])
			assert.Equal(t, tc.expectedRange[0], a, "start of range")
			assert.Equal(t, tc.expectedRange[1], b, "end of range")
		})
	}
}

func TestProcessMap2(t *testing.T) {
	cases := []struct {
		name           string
		dsts           []int
		srcs           []int
		rangeLengths   []int
		seedRangeStart int
		seedRangeLen   int
		expected       [][2]int
	}{
		{
			name: "sample seed-to-soil first seed range",
			dsts: []int{50, 52},
			srcs: []int{98, 50},
			rangeLengths: []int{2, 48},
			seedRangeStart: 79,
			seedRangeLen:   14,
			expected:       [][2]int{{81, 14}},
		},
		{
			name: "sample seed-to-soil second seed range",
			dsts: []int{50, 52},
			srcs: []int{98, 50},
			rangeLengths: []int{2, 48},
			seedRangeStart: 55,
			seedRangeLen:   13,
			expected:       [][2]int{{57, 13}},
		},
		{
			name: "sample soil-to-fertalizer first seed match",
			dsts: []int{0, 37, 39},
			srcs: []int{15, 52, 0},
			rangeLengths: []int{37, 2, 15},
			seedRangeStart: 81,
			seedRangeLen:   14,
			expected:       [][2]int{{81, 14}},
		},
		{
			name: "sample fertalizer-to-water first seed match",
			dsts: []int{49, 0, 42, 57},
			srcs: []int{53, 11, 0, 7},
			rangeLengths: []int{8, 42, 7, 4},
			seedRangeStart: 81,
			seedRangeLen:   14,
			expected:       [][2]int{{81, 14}},
		},
		{
			name: "sample water-to-light first seed match",
			dsts: []int{88, 18},
			srcs: []int{18, 25},
			rangeLengths: []int{7, 70},
			seedRangeStart: 81,
			seedRangeLen:   14,
			expected:       [][2]int{{74, 14}},
		},
		{
			name: "sample light-to-temp first seed match",
			dsts: []int{45, 81, 68},
			srcs: []int{77, 45, 64},
			rangeLengths: []int{23, 19, 13},
			seedRangeStart: 74,
			seedRangeLen:   14,
			expected:       [][2]int{{45, 11}, {78, 3}},
		},
		{
			name: "custom partial range where it's not all matched",
			dsts: []int{45, 81, 68},
			srcs: []int{77, 45, 64},
			rangeLengths: []int{23, 19, 13},
			seedRangeStart: 74,
			seedRangeLen:   14,
			expected:       [][2]int{{45, 11}, {78, 3}},
		},
		{
			name: "sample temp-to-humid first seed match a",
			dsts: []int{0, 1},
			srcs: []int{69, 0},
			rangeLengths: []int{1, 69},
			seedRangeStart: 45,
			seedRangeLen:   11,
			expected:       [][2]int{{46, 11}},
		},
		{
			name: "sample temp-to-humid first seed match b",
			dsts: []int{0, 1},
			srcs: []int{69, 0},
			rangeLengths: []int{1, 69},
			seedRangeStart: 78,
			seedRangeLen:   3,
			expected:       [][2]int{{78, 3}},
		},
		{
			name: "sample humid-to-loc first seed match a",
			dsts: []int{60, 56},
			srcs: []int{56, 93},
			rangeLengths: []int{37, 4},
			seedRangeStart: 46,
			seedRangeLen:   11,
			expected:       [][2]int{{60, 1}, {46, 10}},
		},
		{
			name: "sample humid-to-loc first seed match b",
			dsts: []int{60, 56},
			srcs: []int{56, 93},
			rangeLengths: []int{37, 4},
			seedRangeStart: 78,
			seedRangeLen:   3,
			expected:       [][2]int{{82, 3}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, processMap2(tc.dsts, tc.srcs, tc.rangeLengths, tc.seedRangeStart, tc.seedRangeLen))
		})
	}
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
