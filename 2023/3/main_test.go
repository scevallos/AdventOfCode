package main

import (
	helpers "AdventOfCode"

	"fmt"
	"testing"
)

func TestProcessMatrix(t *testing.T) {
	fmt.Println(ProcessMatrix([][]rune{
		{'4', '.', '.', '2', '1'}, // 4 + 21
		{'.', '#', '.', '*', '.'}, // -
		{'.', '.', '1', '7', '3'}, // 173
		{'1', '.', '.', '5', '6'}, // 1
		{'$', '.', '5', '.', '.'}, // (skip 5)
		// total 199
	}))
}

func TestGetSumAllPartNumbers(t *testing.T) {
	// % = & $ * # @ + - /
	doc, closeFile := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile()
	GetSumAllPartNumbers(doc)
}