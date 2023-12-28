package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	helpers "AdventOfCode"
)

type Race struct {
	Time           int
	RecordDistance int
}

func (r Race) GetBeatableWays() int {
	ways := 0
	for sec := 0; sec < r.Time; sec++ {
		travelDistance := GetDistanceTravelled(sec, r.Time-sec)
		if travelDistance > r.RecordDistance {
			ways++
		}
	}
	return ways
}

func GetDistanceTravelled(heldButtonFor, timeRemaining int) int {
	if heldButtonFor == 0 || timeRemaining == 0 {
		return 0
	}
	return heldButtonFor * timeRemaining
}

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	doc2, closeFile2 := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile2()

	fmt.Println("GetProductOfNumberOfWaysRecordBeatenSingleRace(sampleInput.txt) =", GetProductOfNumberOfWaysRecordBeatenSingleRace(doc))
	fmt.Println("GetProductOfNumberOfWaysRecordBeatenSingleRace(actualInput.txt) =", GetProductOfNumberOfWaysRecordBeatenSingleRace(doc2))
}

func parseInput(doc *bufio.Scanner) []Race {
	// get times from first line
	doc.Scan()
	line := strings.TrimSpace(doc.Text())
	times := helpers.CollectNumsInLine(line, ':', 0)

	// get distance records from second line
	doc.Scan()
	line = strings.TrimSpace(doc.Text())
	distanceRecords := helpers.CollectNumsInLine(line, ':', 0)

	races := make([]Race, len(times))

	for i := 0; i < len(times); i++ {
		races[i] = Race{
			Time:           times[i],
			RecordDistance: distanceRecords[i],
		}
	}

	return races
}

func parseInputSingleNumber(doc *bufio.Scanner) Race {
	// get time parts from first line
	doc.Scan()
	line := strings.TrimSpace(doc.Text())
	parts := strings.Split(line[strings.Index(line, ":")+1:], " ")
	bigTime, err := strconv.Atoi(strings.Join(parts, ""))
	if err != nil {
		panic(err)
	}

	// get distnace parts from second line
	doc.Scan()
	line = strings.TrimSpace(doc.Text())
	parts = strings.Split(line[strings.Index(line, ":")+1:], " ")
	bigRecordDistance, err := strconv.Atoi(strings.Join(parts, ""))
	if err != nil {
		panic(err)
	}
	return Race{
		Time: bigTime,
		RecordDistance: bigRecordDistance,
	}
}

func GetProductOfNumberOfWaysRecordBeaten(doc *bufio.Scanner) int {
	races := parseInput(doc)
	product := 1
	for _, race := range races {
		product *= race.GetBeatableWays()
	}
	return product
}

func GetProductOfNumberOfWaysRecordBeatenSingleRace(doc *bufio.Scanner) int {
	race := parseInputSingleNumber(doc)
	return race.GetBeatableWays()
}
