package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, cs.expected, GetCalibrationValue(cs.doc))
		})
	}
}

func TestProcessLine(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{"only two nums", "12", 12},
		{"two nums and chars", "1abc2", 12},
		{"multiple nums and chars", "a1b2c3d4e5f", 15},
		{"single num and chars", "treb7uchet", 77},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			assert.Equal(t, cs.expected, processLine(cs.input))
		})
	}
}
