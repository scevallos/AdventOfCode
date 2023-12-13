package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	helpers "AdventOfCode"
)

var digitWordsRegex = []*regexp.Regexp{
	regexp.MustCompile("one"),
	regexp.MustCompile("two"),
	regexp.MustCompile("three"),
	regexp.MustCompile("four"),
	regexp.MustCompile("five"),
	regexp.MustCompile("six"),
	regexp.MustCompile("seven"),
	regexp.MustCompile("eight"),
	regexp.MustCompile("nine"),
}

var digitWordsToNums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	doc, closeFile := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile()
	fmt.Println("GetCalibrationValue(actualInput.txt) =", GetCalibrationValue(doc))
}

// O(n)
//   passthru string once
//   for char:
//      if char is digit, store in order
//      then peek and dequeue

// First attempt
// func processLine(line string) int {
// 	// fmt.Println("processing line:", line)
// 	nums := []rune{}
// 	for _, c := range line {
// 		if unicode.IsDigit(c) {
// 			nums = append(nums, c)
// 		}
// 	}
// 	firstAndLast := []rune{nums[0], nums[len(nums)-1]}
// 	out, err := strconv.Atoi(string(firstAndLast))
// 	if err != nil {
// 		panic("bad non-numeric input provided: " + err.Error())
// 	}
// 	return out
// }

// Given an alphanumeric string, returns the list of numbers contained in the string.
// Numbers can appear either as digits [0-9] or as their written-out form ("one", "two", ..., "nine")
func GetDigitsFromLine(line string) []int {
	nums := map[int]int{} // starting index to number repr

	// Get all the numbers represented as digits
	// store them & their index
	for i, c := range line {
		if unicode.IsDigit(c) {
			number, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			nums[i] = number
		}
	}

	// Get all the numbers represented in written-out form
	// store them & their index
	for _, digitWordRegex := range digitWordsRegex {
		matches := digitWordRegex.FindAllStringIndex(line, -1)
		if len(matches) != 0 {
			for _, match := range matches {
				matchedDigitWord := line[match[0]:match[1]]
				matchedDigitWordNum := digitWordsToNums[matchedDigitWord]
				nums[match[0]] = matchedDigitWordNum
			}
		}
	}

	// reduce the map into the list of ints in the order they appeared
	finalNumsList := []int{}
	i := 0
	for i < len(line) {
		val, ok := nums[i]
		if ok {
			finalNumsList = append(finalNumsList, val)
		}
		i++
	}

	return finalNumsList
}

func MakeTwoDigitNumFirstAndLast(digits []int) int {
	num, err := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]))
	if err != nil {
		panic(err)
	}
	return num
}

func GetCalibrationValue(doc *bufio.Scanner) int {
	lineValsSum := 0
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		// lineValsSum += processLine(line)
		digits := GetDigitsFromLine(line)
		lineValsSum += MakeTwoDigitNumFirstAndLast(digits)

	}
	return lineValsSum
}
