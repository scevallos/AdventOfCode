package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Keys[T comparable, E any](someMap map[T]E) []T {
	keys := make([]T, len(someMap))
	i := 0
	for k := range someMap {
		keys[i] = k
		i++
	}
	return keys
}

func GetDocFromString(str string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(str))
}

// return scanner, and closer func which should be defer called
func GetDocFromFile(filename string) (*bufio.Scanner, func() error) {
	file, err := os.Open(filename)
	if err != nil {
		panic("couldn't open the file: " + err.Error())
	}

	return bufio.NewScanner(file), file.Close
}

func MakeRange(i, j int) []int {
	if i > j {
		panic(fmt.Sprintf("MakeRange(%d, %d) invalid: must have i <= j", i, j))
	}
	nums := make([]int, j-i)
	index := 0
	for a := i; a < j; a++ {
		nums[index] = a
		index++
	}
	return nums
}

// CollectNumsInLine(line, after, until)
//   - space separated integers
//   - don't start until `after` char seen
//   - stop once `until` is seen or EOL if unspecified
//
// examples:
//   - ("seeds: 79 14 55 13", ':', 0) --> [79 14 55 13]
//   - ("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", ':', '|') --> [41 48 83 86 17]
//   - ("1 2 3", 0, 0) --> [1 2 3]
func CollectNumsInLine(line string, after, until rune) []int {
	section := 1
	if after == 0 {
		// if no `after` char specified
		// start collecting nums from start
		section = 2
	}
	buildingNumber := false
	wipNumber := []rune{}
	numbers := []int{}
	for _, char := range line {
		switch section {
		case 1:
			if char == after {
				section = 2
			}
			continue
		case 2:
			if unicode.IsSpace(char) {
				if buildingNumber {
					parsedNum, err := strconv.Atoi(string(wipNumber))
					if err != nil {
						panic(err)
					}
					numbers = append(numbers, parsedNum)
					wipNumber = []rune{}
					buildingNumber = false
				} else {
					continue
				}
			} else if unicode.IsDigit(char) {
				wipNumber = append(wipNumber, char)
				buildingNumber = true
			} else if char == until {
				section = 3
				buildingNumber = false
				wipNumber = []rune{}
			}
		}
	}

	if buildingNumber && len(wipNumber) > 0 {
		parsedNum, err := strconv.Atoi(string(wipNumber))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, parsedNum)
	}

	return numbers
}

func IsByteDigit(input byte) bool {
	return input >= 48 && input <= 57
}