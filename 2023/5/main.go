package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	helpers "AdventOfCode"
)

var mapNameToSequenceIndex = map[string]int{
	"seed-to-soil map:":            0,
	"soil-to-fertilizer map:":      1,
	"fertilizer-to-water map:":     2,
	"water-to-light map:":          3,
	"light-to-temperature map:":    4,
	"temperature-to-humidity map:": 5,
	"humidity-to-location map:":    6,
}

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	// doc2, closeFile2 := helpers.GetDocFromFile("actualInput.txt")
	// defer closeFile2()

	fmt.Println("GetLowestLocationNumberForSeeds(sampleInput.txt) =", GetLowestLocationNumberForSeeds(doc))
}

// (50 98 2) --> {98: 50, 99: 51}
func parseMapLine(dstRangeStart, srcRangeStart, rangeLength int) map[int]int {
	srcToDestMap := make(map[int]int, rangeLength)
	for i := 0; i < rangeLength; i++ {
		srcToDestMap[srcRangeStart+i] = dstRangeStart + i
	}
	return srcToDestMap
}

func parseMap(mapRows [][3]int) map[int]int {
	rowMaps := []map[int]int{}
	mergedMap := map[int]int{} // src:dest
	for _, row := range mapRows {
		rowMaps = append(rowMaps, parseMapLine(row[0], row[1], row[2]))
	}

	for _, rowMap := range rowMaps {
		for key, val := range rowMap {
			mergedMap[key] = val
		}
	}

	return mergedMap
}

func parseSeeds(text string) []int {
	return helpers.CollectNumsInLine(text, ':', 0)
}

func GetLowestLocationNumberForSeeds(doc *bufio.Scanner) int {
	lowestLocation := math.MaxInt

	// get seeds from first line
	doc.Scan()
	firstLine := doc.Text()
	seeds := parseSeeds(firstLine)
	_ = seeds

	isBuildingMap := false
	collectMapRows := [][3]int{}
	allTheMaps := [7][][3]int{}
	workingOnMapIndex := 0

	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		if _, isMapStarter := mapNameToSequenceIndex[line]; isMapStarter {
			isBuildingMap = true
			continue
		} else if line == "" {
			if !isBuildingMap {
				continue
			}
			// collect map
			// allTheMaps = append(allTheMaps, collectMapRows)
			allTheMaps[workingOnMapIndex] = collectMapRows
			workingOnMapIndex++
		} else {
			// assume line is line of numbers
			collectMapRows = append(collectMapRows, ([3]int)(helpers.CollectNumsInLine(line, 0, 0)))
		}
	}
	_ = allTheMaps
	return lowestLocation
}
