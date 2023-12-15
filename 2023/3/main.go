package main

import (
	helpers "AdventOfCode"
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()
	
	fmt.Println("GetSumAllPartNumbers(sampleInput.txt) =", GetSumAllPartNumbers(doc))
	// fmt.Println("GetSumIdsPossibleGames(sampleInput.txt) =", GetSumIdsPossibleGames(doc))
	// fmt.Println("GetSumGamePowers(actualInput.txt) =", GetSumGamePowers(doc))
}

// questions:
//   scan thru for symbols and look for surrounding numbers OR
//   scan for numbers and look for surroinding symbols
//   faster opt depends on if there are more symbols than numbers or the other way around
//   
//   based on sample input, assume: there are less symbols than numbers

// O(n * m)
//   first pass thru (text): parse into matrix
//   second pass thru (matrix): if char is non-dot symbol
//     check surroundings, and flag each surrounding number as part num so add to sum



// ProcessMatrix computes the sum of all the digits in the matrix that have
// a neighboring symbol (1 char away horizontally, vertically, or diagonally)
func ProcessMatrix(matrix [][]rune) int {
	set := map[string]struct{}{}
	aggrSum := 0
	for i, row := range matrix {
		for j := range row {
			char := matrix[i][j]
			if char != '.' && unicode.In(char, unicode.Punct, unicode.Symbol) {
				
			}
		}
	}
	
	for key, _ := range set {
		fmt.Printf("%s ", key)
	}
	fmt.Print("\n")
	return -1
}



func GetSumAllPartNumbers(doc *bufio.Scanner) int {
	partNumbersSum := 0
	matrix := [][]rune{}
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	ProcessMatrix(matrix)
	return partNumbersSum
}