package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	helpers "AdventOfCode"
)

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	doc2, closeFile2 := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile2()

	fmt.Println("GetLowestLocationNumberForSeedRanges(sampleInput.txt) =", GetLowestLocationNumberForSeedRanges(doc))
	fmt.Println("GetLowestLocationNumberForSeedRanges(actualInput.txt) =", GetLowestLocationNumberForSeedRanges(doc2))
}

// 50 98 2
// 52 50 48
// assumption: len(dests) == len(srcs) == len(rangeLengths)
func processMap(dests, srcs, rangeLengths []int, source int) int {
	for j := range helpers.MakeRange(0, len(dests)) {
		if srcs[j] <= source && source < srcs[j]+rangeLengths[j] {
			return dests[j] + (source - srcs[j])
		}
	}
	return source
}

// assumption: len(dests) == len(srcs) == len(rangeLengths)
func processMap2(dests, srcs, rangeLengths []int, seedRangeStart, seedRangeLen int) [][2]int {
	matches := [][2]int{}
	destRange := [2]int{}
	minDelta := math.MaxInt
	var delta int
	for rowNum := 0; rowNum < len(dests); rowNum++ {
		overlapStart, overlapEnd := getOverlappingRange(srcs[rowNum], srcs[rowNum]+rangeLengths[rowNum], seedRangeStart, seedRangeStart+seedRangeLen)
		if overlapStart == -1 && overlapEnd == -1 {
		  continue
		}
		delta = dests[rowNum] - srcs[rowNum]
		if delta < minDelta {
			minDelta = delta
		}
		matches = append(matches, [2]int{overlapStart, overlapEnd-1})
	}
	if len(matches) == 0 {
		matches = append(matches, [2]int{seedRangeStart, seedRangeStart+seedRangeLen-1})
	}
	for _, rangeMatch := range matches {

	}
	return matches
}

// starting clean
// seeds: 79 14 55 13
// 79+14 seed range

// seed-to-soil map:
// 50 98 2 --> no match
// 52 50 48 --> +2, match 79+14

// 79+14 => 81+14

// soil-to-fertilizer map:
// 0 15 37 --> no match
// 37 52 2 --> no match
// 39 0 15 --> no match

// 81+14 => 81+14

// fertilizer-to-water map:
// 49 53 8 --> no match
// 0 11 42 --> no match
// 42 0 7 --> no match
// 57 7 4 --> no match

// 81+14 => 81+14

// water-to-light map:
// 88 18 7 --> no match
// 18 25 70 --> -7, match 79+14

// 81+14 => 74+14

// light-to-temperature map:
// 45 77 23 --> -32, match 77+14
// 81 45 19 --> no match
// 68 64 13 --> +4, match 74+3

// 74+14 => (45+14, 78+3)

// temperature-to-humidity map:
// 0 69 1 --> no match
// 1 0 69 --> +1, match 45+14

// 45+14 => 46+14
// 78+3 => 78+3

// humidity-to-location map:
// 60 56 37 --> +4, match 78+3
// 56 93 4 --> no match

// 46+14 => 46+14
// 78+3 => 82+3

// (2-5),(4-7) --> 4-5
// (2-5),(1-3) --> 1-3
// (2-5),(1-7) --> 2-5
// (2-5),(3-4) --> 3-4
// assumptions:
//   - i < j
//   - x < y
//   - i != j && x != y
//   - i, j, x, y > 0
func getOverlappingRange(i, j, x, y int) (int, int) {
	if i <= y && y < j && x < i { // left half
		return i, y
	} else if x <= j && j < y && i < x { // right half
		return x, j
	} else if x <= i && y >= j {
		return i, j
	} else if x >= i && y <= j {
		return x, y
	}
	return -1, -1
}

func parseSeeds(text string) []int {
	return helpers.CollectNumsInLine(text, ':', 0)
}

func parseInput(doc *bufio.Scanner) ([]int, [][][]int) {
	// get seeds from first line
	doc.Scan()
	firstLine := doc.Text()
	seeds := parseSeeds(firstLine)

	collectingMapDests, collectingMapSrcs, collectingMapRanges := []int{}, []int{}, []int{}
	allTheMaps := [][][]int{}
	// workingOnMapIndex := 0

	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		nums := helpers.CollectNumsInLine(line, 0, 0)
		if len(nums) == 0 && len(collectingMapDests) != 0 {
			// collect map
			allTheMaps = append(allTheMaps, [][]int{
				collectingMapDests,
				collectingMapSrcs,
				collectingMapRanges,
			})
			collectingMapDests, collectingMapSrcs, collectingMapRanges = []int{}, []int{}, []int{}
			continue
		}
		if len(nums) != 3 {
			continue
		}
		collectingMapDests = append(collectingMapDests, nums[0])
		collectingMapSrcs = append(collectingMapSrcs, nums[1])
		collectingMapRanges = append(collectingMapRanges, nums[2])
	}

	// collect the last one
	if len(collectingMapDests) != 0 {
		allTheMaps = append(allTheMaps, [][]int{
			collectingMapDests,
			collectingMapSrcs,
			collectingMapRanges,
		})
	}
	return seeds, allTheMaps
}

func GetLowestLocationNumberForSeeds(doc *bufio.Scanner) int {
	lowestLocation := math.MaxInt

	seeds, allTheMaps := parseInput(doc)
	for _, seed := range seeds {
		src := seed
		results := []int{src}
		for i := 0; i < len(allTheMaps); i++ {
			dest := processMap(allTheMaps[i][0], allTheMaps[i][1], allTheMaps[i][2], src)
			results = append(results, dest)
			src = dest
		}
		// fmt.Println(results)
		location := results[len(results)-1]
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func GetLowestLocationNumberForSeedRanges(doc *bufio.Scanner) int {
	lowestLocation := math.MaxInt

	seedInputs, allTheMaps := parseInput(doc)
	seedRangeStart := -1
	seedRangeLength := -1
	seedRanges := [][2]int{}
	for i, value := range seedInputs {
		if i % 2 == 0 {
			seedRangeStart = value
		} else {
			seedRangeLength = value
		}
		if seedRangeLength != -1 && seedRangeStart != -1 {
			seedRanges = append(seedRanges, [2]int{seedRangeStart, seedRangeLength})
			seedRangeLength = -1
			seedRangeStart = -1
		}
	}

	for _, seedRangeStrtAndLen := range seedRanges {
		seedRangeStart, seedRangeLen := seedRangeStrtAndLen[0], seedRangeStrtAndLen[1]
		// results := []int{src}
		for i := 0; i < len(allTheMaps); i++ {
			nextRanges := processMap2(allTheMaps[i][0], allTheMaps[i][1], allTheMaps[i][2], seedRangeStart, seedRangeLen)
			_ = nextRanges
			// results = append(results, dest)
			// src = dest
		}
		// fmt.Println(results)
		// location := results[len(results)-1]
		location := 1
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}
