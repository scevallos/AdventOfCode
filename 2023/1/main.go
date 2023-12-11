package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("foo1")
}

// O(n)
//   passthru string once
//   for char:
//      if char is digit, store in order
//      then peek and dequeue

func processLine(line string) int {
	// fmt.Println("processing line:", line)
	nums := []rune{}
	for _, c := range line {
		if unicode.IsDigit(c) {
			nums = append(nums, c)
		}
	}
	firstAndLast := []rune{nums[0], nums[len(nums)-1]}
	out, err := strconv.Atoi(string(firstAndLast))
	if err != nil {
		panic("bad non-numeric input provided: " + err.Error())
	}
	return out
}

func GetCalibrationValue(doc *bufio.Scanner) int {
	lineValsSum := 0
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		lineValsSum += processLine(line)
	}
	return lineValsSum
}

func getDocFromString(str string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(str))
}

// return scanner, and closer func which should be defer called
func getDocFromFile(filename string) (*bufio.Scanner, func() error) {
	file, err := os.Open(filename)
    if err != nil {
		panic("couldn't open the file: " + err.Error())
    }

    return bufio.NewScanner(file), file.Close
}