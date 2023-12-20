package main

import (
	"reflect"
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

func TestGetLowestLocationNumberForSeeds(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 35, GetLowestLocationNumberForSeeds(doc))
}

func TestParseMap(t *testing.T) {
	cases := []struct {
		name     string
		mapRows  [][3]int
		expected any
	}{
		{
			name:     "sample seed-to-soil",
			mapRows:  [][3]int{{50, 98, 2}, {52, 50, 48}},
			expected: 8,
		},
		{
			name:     "sample soil-to-fertilizer",
			mapRows:  [][3]int{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}},
			expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, parseMap(tc.mapRows))
		})
	}
}

func TestParseMapLine(t *testing.T) {
	cases := []struct {
		name          string
		dstRangeStart int
		srcRangeStart int
		rangeLength   int
		expectedFunc  func(out map[int]int) bool
	}{
		{
			name:          "sample seed-to-soil first line",
			dstRangeStart: 50,
			srcRangeStart: 98,
			rangeLength:   2,
			expectedFunc: func(out map[int]int) bool {
				return reflect.DeepEqual(map[int]int{
					98: 50,
					99: 51,
				}, out)
			},
		},
		{
			name:          "actual soil-to-fertilizer random",
			dstRangeStart: 3261136238,
			srcRangeStart: 2193516168,
			rangeLength:   29269446,
			expectedFunc: func(out map[int]int) bool {
				return false
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t,
				tc.expectedFunc(
					parseMapLine(
						tc.dstRangeStart,
						tc.srcRangeStart,
						tc.rangeLength,
					)))
		})
	}
}
