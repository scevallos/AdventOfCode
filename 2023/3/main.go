package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	helpers "AdventOfCode"
)

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	fmt.Println("GetSumAllPartNumbers(sampleInput.txt) =", GetSumAllPartNumbers(doc))
	// fmt.Println("GetSumIdsPossibleGames(sampleInput.txt) =", GetSumIdsPossibleGames(doc))
	// fmt.Println("GetSumGamePowers(actualInput.txt) =", GetSumGamePowers(doc))
}

// assume: i < j
func makeRange(i, j int) []int {
	nums := []int{}
	for; i <= j; i++ {
		nums = append(nums, i)
	}
	return nums
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
	aggrSum := 0
	// i,j to whether or not it's been constructed yet
	flaggedPoints := map[[2]int]bool{}
	for i, row := range matrix {
		for j := range row {
			char := matrix[i][j]
			if char != '.' && unicode.In(char, unicode.Punct, unicode.Symbol) {
				for _, k := range []int{i - 1, i, i + 1} {
					var partNumber string
					var constructingNumber bool
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
						// fmt.Println("srPt", string(surroundingPoint))
						if unicode.IsDigit(surroundingPoint) {
							// check if flagged already
							_, alreadyFlagged := flaggedPoints[[2]int{k, h}]
							if alreadyFlagged {
								continue
							}

							if constructingNumber {
								// append it to WIP number
								partNumber = partNumber + string(surroundingPoint)
								// fmt.Printf("a) appending %s to partNumber\n", string(surroundingPoint))
							} else {
								// start new WIP number
								constructingNumber = true
								partNumber = partNumber + string(surroundingPoint)
								// fmt.Printf("b) appending %s to partNumber\n", string(surroundingPoint))
							}
							flaggedPoints[[2]int{k, h}] = false
							
						} else if constructingNumber {
							// terminate WIP number
							constructingNumber = false
							// mark as constructed
							flaggedPoints[[2]int{k, h-1}] = true
							// fmt.Printf("a) marking (%d, %d) as constructed\n", k, h-1)
							partNumInt, err := strconv.Atoi(partNumber)
							if err != nil {
								panic(err)
							}
							aggrSum += partNumInt
							// fmt.Printf("a) adding %d to aggrSum\n", partNumInt)
							
							partNumber = ""
						}
					}
					// finished row scan & was constructingNumber, so finish the construction and add
					if constructingNumber {
						var hDiff int
						// fmt.Println("looking over ", string(matrix[k][h:]))
						for hDiff, char = range matrix[k][h:] {
							// check if flagged already
							_, alreadyFlagged := flaggedPoints[[2]int{k, h + hDiff}]
							if alreadyFlagged {
								continue
							}

							if unicode.IsDigit(char) {
								partNumber = partNumber + string(char)
								// fmt.Printf("c) appending %s to partNumber\n", string(char))
							} else {
								// terminate number construction and break
								constructingNumber = false
								partNumInt, err := strconv.Atoi(partNumber)
								if err != nil {
									panic(err)
								}
								aggrSum += partNumInt
								fmt.Printf("b) adding %d to aggrSum\n", partNumInt)
								// mark as constructed
								flaggedPoints[[2]int{k, h + hDiff}] = true
								fmt.Printf("b) marking (%d, %d) as constructed\n", k, h + hDiff)
								break
							}
						}

						// if not constructed, do it
						if !flaggedPoints[[2]int{k, h}] {
							constructingNumber = false
							for _, index := range makeRange(h, hDiff + h) {
								// mark (k, h) --> (k, h + hDiff) as true
								flaggedPoints[[2]int{k, index}] = true
								// fmt.Printf("c) marking (%d, %d) as constructed\n", k, index)
							}
							partNumInt, err := strconv.Atoi(partNumber)
							if err != nil {
								panic(err)
							}
							aggrSum += partNumInt
							// fmt.Printf("c) adding %d to aggrSum\n", partNumInt)
						}
					}
				}
			}
		}
	}

	// fmt.Println(flaggedPoints)

	// for point, _ := range flaggedPoints {
	// 	row, col := point[0], point[1]
	// 	fmt.Println(string(matrix[row][col]))
	// }

	// for key := range set {
	// 	fmt.Printf("%s ", key)
	// }
	// fmt.Print("\n")
	return aggrSum
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
