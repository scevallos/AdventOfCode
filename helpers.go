package helpers

import (
	"bufio"
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
