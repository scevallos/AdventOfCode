package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	sampleInput = `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`
	expectedOutput = 142
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

func GetCalibrationValue(doc string) int {
	scanner := bufio.NewScanner(strings.NewReader(sampleInput))
	lineValsSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineValsSum += processLine(line)
	}
	return lineValsSum
}