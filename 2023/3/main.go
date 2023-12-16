package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	helpers "AdventOfCode"
)

const (
	PartNumberSum = iota
	GearRatio
)

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	doc2, closeFile2 := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile2()

	// fmt.Println("GetSumAllPartNumbers(sampleInput.txt) =", GetSumAllPartNumbers(doc))
	// fmt.Println("GetSumAllPartNumbers(actualInput.txt) =", GetSumAllPartNumbers(doc2))
	fmt.Println("GetSumAllGearRatios(sampleInput.txt) =", GetSumAllGearRatios(doc))
	fmt.Println("GetSumAllGearRatios(actualInput.txt) =", GetSumAllGearRatios(doc2))
}

func AbsorbNumber(chars []rune, index int) (string, []int) {
	return absorbNumber(chars, index, index, "", []int{})
}

func absorbNumber(chars []rune, index, globalIndex int, parts string, indices []int) (string, []int) {
	if index < 0 || index >= len(chars) {
		return parts, indices
	}

	self := chars[index]
	if !unicode.IsDigit(self) {
		return parts, indices
	}

	var left, right string
	var leftIndices, rightIndices []int
	if index-1 >= 0 {
		left, leftIndices = absorbNumber(chars[:index], index-1, index-1, "", []int{})
	}

	if index+1 < len(chars) {
		right, rightIndices = absorbNumber(chars[index+1:], 0, globalIndex+1, "", []int{})
	}

	allIndices := append(leftIndices, globalIndex)
	allIndices = append(allIndices, rightIndices...)

	return left + string(self) + right, allIndices
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
func ProcessMatrix(matrix [][]rune, processingType int) int {
	aggrSum := 0
	// i,j to whether or not it's been constructed yet
	flaggedPoints := map[[2]int]struct{}{}
	for i, row := range matrix {
		for j := range row {
			char := matrix[i][j]
			switch processingType {
			case PartNumberSum:
				if char != '.' && unicode.In(char, unicode.Punct, unicode.Symbol) {
					for _, k := range []int{i - 1, i, i + 1} {
						var h int
						for _, h = range []int{j - 1, j, j + 1} {
							if k == i && h == j {
								continue
							}
							var surroundingRow []rune
							var surroundingPoint rune
							if k >= 0 && k < len(matrix) {
								surroundingRow = matrix[k]
							}
							if h >= 0 && h < len(surroundingRow) && surroundingRow != nil {
								surroundingPoint = surroundingRow[h]
							}
							if unicode.IsDigit(surroundingPoint) {
								_, alreadyFlagged := flaggedPoints[[2]int{k, h}]
								if alreadyFlagged {
									continue
								}
	
								// fmt.Printf("absorbing number from digit at %s\n", string(matrix[k][h]))
								absorbedNum, absorbedIndices := AbsorbNumber(matrix[k], h)
								partNumInt, err := strconv.Atoi(absorbedNum)
								if err != nil {
									panic(err)
								}
								// fmt.Printf("x) adding %d to aggrSum\n", partNumInt)
								aggrSum += partNumInt
								for _, absIndex := range absorbedIndices {
									flaggedPoints[[2]int{k, absIndex}] = struct{}{}
								}
							}
						}
					}
				}
			case GearRatio:
				if char == '*' {
					absorbedNumbers := []int{}
					for _, k := range []int{i - 1, i, i + 1} {
						var h int
						for _, h = range []int{j - 1, j, j + 1} {
							if k == i && h == j {
								continue
							}
							var surroundingRow []rune
							var surroundingPoint rune
							if k >= 0 && k < len(matrix) {
								surroundingRow = matrix[k]
							}
							if h >= 0 && h < len(surroundingRow) && surroundingRow != nil {
								surroundingPoint = surroundingRow[h]
							}
							if unicode.IsDigit(surroundingPoint) {
								_, alreadyFlagged := flaggedPoints[[2]int{k, h}]
								if alreadyFlagged {
									continue
								}
	
								// fmt.Printf("absorbing number from digit at %s\n", string(matrix[k][h]))
								absorbedNum, absorbedIndices := AbsorbNumber(matrix[k], h)
								partNumInt, err := strconv.Atoi(absorbedNum)
								if err != nil {
									panic(err)
								}
								// fmt.Printf("x) adding %d to aggrSum\n", partNumInt)
								absorbedNumbers = append(absorbedNumbers, partNumInt)
								for _, absIndex := range absorbedIndices {
									flaggedPoints[[2]int{k, absIndex}] = struct{}{}
								}
							}
						}
					}
					if len(absorbedNumbers) == 2 {
						aggrSum += absorbedNumbers[0] * absorbedNumbers[1]
					}
				}
			}
		}
	}
	// fmt.Println(flaggedPoints)
	return aggrSum
}

func GetSumAllPartNumbers(doc *bufio.Scanner) int {
	matrix := [][]rune{}
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	return ProcessMatrix(matrix, PartNumberSum)
}

func GetSumAllGearRatios(doc *bufio.Scanner) int {
	matrix := [][]rune{}
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	return ProcessMatrix(matrix, GearRatio)
}
