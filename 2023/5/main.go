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

	fmt.Println("GetLowestLocationNumberForSeeds(sampleInput.txt) =", GetLowestLocationNumberForSeeds(doc))
	fmt.Println("GetLowestLocationNumberForSeeds(actualInput.txt) =", GetLowestLocationNumberForSeeds(doc2))
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
