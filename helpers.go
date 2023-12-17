package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func MakeRange(i, j int) []int{
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
