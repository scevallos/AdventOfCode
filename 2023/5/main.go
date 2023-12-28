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
	// fmt.Println("processMap2 called with", dests, srcs, rangeLengths, seedRangeStart, seedRangeLen)
	destRanges := [][2]int{}
	var delta int
	maybePartialMatches := [][2]int{}
	for rowNum := 0; rowNum < len(dests); rowNum++ {
		overlapStart, overlapEnd := getOverlappingRange(srcs[rowNum], srcs[rowNum]+rangeLengths[rowNum], seedRangeStart, seedRangeStart+seedRangeLen)
		if overlapStart == -1 && overlapEnd == -1 {
		  continue
		}
		overlapSize := overlapEnd - overlapStart
		delta = dests[rowNum] - srcs[rowNum]
		// if delta < minDelta {
		// 	minDelta = delta
		// }
		destRanges = append(destRanges, [2]int{overlapStart+delta, overlapSize})
		
		// if only partial match, grab the non-matched portion too
		if overlapSize != seedRangeLen {
			maybePartialMatches = append(maybePartialMatches, [2]int{seedRangeStart, seedRangeLen - overlapSize})
		}
	}
	if len(destRanges) == 0 {
		destRanges = append(destRanges, [2]int{seedRangeStart, seedRangeLen})
	}
	// if dest ranges doesn't cover the whole range we looked at
	// look at partial range

	destRangeLength := 0
	for _, destRange := range destRanges {
		destRangeLength += destRange[1]
	}

	// fmt.Println("maybePartials", maybePartialMatches)
	// fmt.Println("destRanges", destRanges)
	// fmt.Println("srcRange", seedRangeStart, seedRangeLen)

	finalRanges := destRanges

	if destRangeLength != seedRangeLen {
		// destRanges = append(destRanges, maybePartialMatches...)
		for _, partialRange := range maybePartialMatches {
			for _, destRange := range destRanges {
				overlapStart, overlapEnd := getOverlappingRange(partialRange[0], partialRange[0]+partialRange[1], destRange[0], destRange[0]+destRange[1])
				if overlapStart == -1 && overlapEnd == -1 {
					finalRanges = append(finalRanges, partialRange)
				}
			}
		}
	}



	// fmt.Println("Returning destRanges", destRanges)
	// for _, rangeMatch := range matches {

	// }
	return finalRanges
}

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
		lookingAtRange := [][2]int{{seedRangeStart, seedRangeLen}}
		nextRange := lookingAtRange
		allTheNumbersInOrder := [][2]int{}
		for i := 0; i < len(allTheMaps); i++ {
			// fmt.Println("nextRange (before", nextRange)
			var y int
			var numRange [2]int
			for y, numRange = range nextRange {
				// TODO something weird happening here
				// like bc iterating over lookingAtRange while changing its value 
				// lookingAtRange := lookingAtRange
				// fmt.Println("looking at (before)", numRange)
				// fmt.Printf("calling processMap2(%v, %v, %v, %v, %v)\n", allTheMaps[i][0], allTheMaps[i][1], allTheMaps[i][2], numRange[0], numRange[1])
				// aggregate into a nextRange var
				lookingAtRange = processMap2(allTheMaps[i][0], allTheMaps[i][1], allTheMaps[i][2], numRange[0], numRange[1])
				allTheNumbersInOrder = append(allTheNumbersInOrder, lookingAtRange...)
				// fmt.Println("looking at (after)", lookingAtRange)
			}

			if y == 0 {
				nextRange = lookingAtRange
			} else {
				nextRange = allTheNumbersInOrder[len(allTheNumbersInOrder)-y-1:]
				// fmt.Println("y", y)
				// fmt.Println("lookingAtRange", lookingAtRange)
				// fmt.Println("allTheNumbersInOrder", allTheNumbersInOrder)
			}
			// 
			// fmt.Println("nextRange (outside)", nextRange)
		}

		// fmt.Println("allTheNumbersInOrder", allTheNumbersInOrder)
		fmt.Println("possible locations", nextRange)

		for _, locationRanges := range nextRange {
			if locationRanges[0] < lowestLocation {
				// fmt.Println("setting new lowest location to", locationRanges[0])
				lowestLocation = locationRanges[0]
			}
		}
	}

	return lowestLocation
}
