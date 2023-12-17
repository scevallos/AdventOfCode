package main

import (
	helpers "AdventOfCode"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var (
	scratchCardMatches = map[int]int{}
)

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	doc2, closeFile2 := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile2()

	// fmt.Println("GetSumAllScratchcard(sampleInput.txt) =", GetSumAllScratchcard(doc))
	// fmt.Println("GetSumAllScratchcard(actualInput.txt) =", GetSumAllScratchcard(doc2))
	fmt.Println("GetTotalScratchcard(sampleInput.txt) =", GetTotalScratchcard(doc))
	fmt.Println("GetTotalScratchcard(actualInput.txt) =", GetTotalScratchcard(doc2))
}

// "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53" --> 8
func GetScratchcardValue(text string) int {
	_, matches := getCardNumberAndMatches(text)
	return int(math.Pow(2, float64(matches-1)))
}

func getCardNumberAndMatches(text string) (int, int) {
	section := 1
	buildingNumber := false
	wipNumber := []rune{}
	winningNumbers := map[int]struct{}{}
	matches := 0
	gameNumber := 0
	for _, char := range text {
		switch section {
		case 1:
			if char == ':' {
				section = 2
				parsedNum, err := strconv.Atoi(string(wipNumber))
				if err != nil {
					panic(err)
				}
				gameNumber = parsedNum
				wipNumber = []rune{}
			} else if unicode.IsDigit(char) {
				wipNumber = append(wipNumber, char)
			}
			continue
		case 2:
			if unicode.IsSpace(char) {
				if buildingNumber {
					parsedNum, err := strconv.Atoi(string(wipNumber))
					if err != nil {
						panic(err)
					}
					// fmt.Printf("adding %d to winningNumbers\n", parsedNum)
					winningNumbers[parsedNum] = struct{}{}
					wipNumber = []rune{}
					buildingNumber = false
				} else {
					continue
				}
			} else if unicode.IsDigit(char) {
				wipNumber = append(wipNumber, char)
				buildingNumber = true
			} else if char == '|' {
				section = 3
				buildingNumber = false
				wipNumber = []rune{}
			}
		case 3:
			if unicode.IsSpace(char) {
				if buildingNumber {
					parsedNum, err := strconv.Atoi(string(wipNumber))
					if err != nil {
						panic(err)
					}
					wipNumber = []rune{}
					buildingNumber = false

					// fmt.Println("checking if", parsedNum, "is a match")
					if _, isMatch := winningNumbers[parsedNum]; isMatch {
						matches++
					}
				} else {
					continue
				}
			} else if unicode.IsDigit(char) {
				wipNumber = append(wipNumber, char)
				buildingNumber = true
			}
		}
	}
	// construct last number & check for match
	parsedNum, err := strconv.Atoi(string(wipNumber))
	if err != nil {
		panic(err)
	}
	if _, isMatch := winningNumbers[parsedNum]; isMatch {
		matches++
	}

	return gameNumber, matches
}

func processCard(cardNumber int) int {
	// process 1 --> 4 matches
	//    return 1 + process(2, 3, 4, 5)
	// ...
	// process i --> n matches
	//    return 1 + process(i+1, ..., i+n)

	matches, ok := scratchCardMatches[cardNumber]
	if !ok {
		panic(fmt.Sprintf("card number %d not in map!", cardNumber))
	}

	if matches == 0 {
		return 1
	}

	sum := 0
	for _, nextCardNumber := range helpers.MakeRange(cardNumber+1, cardNumber + matches + 1) {
		sum += processCard(nextCardNumber)
	}

	return 1 + sum
}

func GetSumAllScratchcard(doc *bufio.Scanner) int {
	points := 0
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		points += GetScratchcardValue(line)
	}
	return points
}

func GetTotalScratchcard(doc *bufio.Scanner) int {
	// process all the scratchcards
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		cardNumber, matches := getCardNumberAndMatches(line)
		scratchCardMatches[cardNumber] = matches
	}

	totalCards := 0
	for cardNumber := range scratchCardMatches {
		totalCards += processCard(cardNumber)
	}

	return totalCards
}
