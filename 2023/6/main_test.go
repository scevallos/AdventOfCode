package main

import (
	"fmt"
	"testing"

	helpers "AdventOfCode"

	"github.com/stretchr/testify/assert"
)

func TestGetProductOfNumberOfWaysRecordBeaten(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	assert.Equal(t, 288, GetProductOfNumberOfWaysRecordBeaten(doc))
}

func TestParseInput(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	races := parseInput(doc)
	assert.Equal(t, Race{7, 9}, races[0])
	assert.Equal(t, Race{15, 40}, races[1])
	assert.Equal(t, Race{30, 200}, races[2])
}

func TestParseInputSingleNumber(t *testing.T) {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	fmt.Println(parseInputSingleNumber(doc))
}