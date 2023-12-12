package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sampleInput = `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`
	expectedOutput   = 142
	submittedAnswer1 = 54573
	submittedAnswer2 = 54591
)

func TestGetCalibrationValue(t *testing.T) {
	cases := []struct {
		name     string
		doc      string
		expected int
	}{
		{"sample input", sampleInput, expectedOutput},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			assert.Equal(t, cs.expected, GetCalibrationValue(getDocFromString(cs.doc)))
		})
	}
}

func TestGetDocFromFile(t *testing.T) {
	doc, closeFile := GetDocFromFile("sampleInput.txt")
	defer closeFile()
	assert.Equal(t, expectedOutput, GetCalibrationValue(doc))

	doc, closeFile = GetDocFromFile("actualInput.txt")
	defer closeFile()
	assert.Equal(t, submittedAnswer2, GetCalibrationValue(doc))
}

// func TestProcessLine(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		input    string
// 		expected int
// 	}{
// 		{"only two nums", "12", 12},
// 		{"two nums and chars", "1abc2", 12},
// 		{"multiple nums and chars", "a1b2c3d4e5f", 15},
// 		{"single num and chars", "treb7uchet", 77},
// 	}
// 	for _, cs := range cases {
// 		t.Run(cs.name, func(t *testing.T) {
// 			assert.Equal(t, cs.expected, processLine(cs.input))
// 		})
// 	}
// }

func TestGetDigitsFromLine(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []int
	}{
		{"only two nums", "12", []int{1, 2}},
		{"two nums and chars", "1abc2", []int{1, 2}},
		{"multiple nums and chars", "a1b2c3d4e5f", []int{1, 2, 3, 4, 5}},
		{"single num and chars", "treb7uchet", []int{7}},
		{"digit words and num", "two1nine", []int{2, 1, 9}},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			assert.Equal(t, cs.expected, GetDigitsFromLine(cs.input))
		})
	}
}
