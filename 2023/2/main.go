package main

import (
	helpers "AdventOfCode"
	"strings"

	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

type CubeColor int

const (
	Green CubeColor = iota
	Red
	Blue
)

var (
	StartingBag = Bag{
		Reds:   12,
		Greens: 13,
		Blues:  14,
	}
	GameFormatRegex = regexp.MustCompile(`Game (\d+):(.*)`)
	TextToCubeColor = map[string]CubeColor{
		"green": Green,
		"red":   Red,
		"blue":  Blue,
	}
)

type Bag struct {
	Reds   int
	Greens int
	Blues  int
}

type Game struct {
	ID   int
	Sets []*Set
}

type Set struct {
	RedsDrawn   int
	GreensDrawn int
	BluesDrawn  int
}

func main() {
	doc, closeFile := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile()
	fmt.Println("GetSumIdsPossibleGames(actualInput.txt) =", GetSumIdsPossibleGames(doc))
}

func ParseSets(text string) []*Set {
	// fmt.Println("parsing set", text)
	var startedParsingToken, tokenIsNumber, tokenIsText bool
	tokens := [][]int{}
	tokensInSet := []int{}
	tokenStore := []rune{}
	for _, char := range text {
		if unicode.IsDigit(char) {
			// parse num & continue to store until end
			if !startedParsingToken {
				startedParsingToken = true
				tokenIsNumber = true
				tokenStore = append(tokenStore, char)
			} else if tokenIsNumber {
				tokenStore = append(tokenStore, char)
			}
		} else if unicode.IsLetter(char) {
			// parse color & continue to store until end
			if !startedParsingToken {
				startedParsingToken = true
				tokenIsText = true
				tokenStore = append(tokenStore, char)
			} else if tokenIsText {
				tokenStore = append(tokenStore, char)
			} else if tokenIsNumber {
				panic("ran into number char in the middle of parsing text token")
			}
		} else if unicode.IsPunct(char) && startedParsingToken {
			// we hit a comma or semi-colon so we're at the end of a (text) token
			// complete previous token
			if tokenIsText {
				parsedToken := int(TextToCubeColor[string(tokenStore)])
				tokensInSet = append(tokensInSet, parsedToken)
				tokenStore = []rune{}
				startedParsingToken = false
				tokenIsNumber = false
				tokenIsText = false
			} else if tokenIsNumber {
				panic("unexpected number token followed by punct")
			}
			// if it's a semi-colon, complete the set & reset it
			if char == ';' {
				tokens = append(tokens, tokensInSet)
				tokensInSet = []int{}
			}
		} else if unicode.IsSpace(char) {
			// in between tokens - complete previous token if we've started one
			if startedParsingToken {
				var parsedToken int
				var err error
				if tokenIsNumber {
					parsedToken, err = strconv.Atoi(string(tokenStore))
					if err != nil {
						panic("failed to parse numToken: " + err.Error())
					}
				} else if tokenIsText {
					parsedToken = int(TextToCubeColor[string(tokenStore)])
				}

				tokensInSet = append(tokensInSet, parsedToken)
				tokenStore = []rune{}
				startedParsingToken = false
				tokenIsNumber = false
				tokenIsText = false
			}
		} else {
			panic("unexpected token: " + string(char))
		}
	}

	// complete last token
	if len(tokenStore) > 0 && startedParsingToken && tokenIsText {
		parsedToken := int(TextToCubeColor[string(tokenStore)])
		tokensInSet = append(tokensInSet, parsedToken)
		tokens = append(tokens, tokensInSet)
	}

	// fmt.Println("tokenStore:", string(tokenStore))
	// fmt.Println("tokens:", tokens)

	// translate tokens into set structs
	return translateTokensToSets(tokens)
}

func translateTokensToSets(tokens [][]int) []*Set {
	sets := make([]*Set, len(tokens))
	for i, setTokens := range tokens {
		set := &Set{}
		var value int
		for j, token := range setTokens {
			if j%2 == 0 { // number
				value = token
			} else { // color ID
				switch token {
				case int(Green):
					set.GreensDrawn = value
				case int(Red):
					set.RedsDrawn = value
				case int(Blue):
					set.BluesDrawn = value
				}
			}
		}
		sets[i] = set
	}
	return sets
}

func ParseLineToGame(line string) *Game {
	gameFormatMatches := GameFormatRegex.FindStringSubmatch(line)

	if len(gameFormatMatches) != 3 {
		panic("No match found or bad input??")
	}

	// get game ID
	id, err := strconv.Atoi(gameFormatMatches[1])
	if err != nil {
		panic("bad game id conversion:" + err.Error())
	}

	// get game details
	sets := gameFormatMatches[2]
	game := &Game{
		ID:   id,
		Sets: ParseSets(sets),
	}

	return game
}

func (b *Bag) IsGamePossible(g *Game) bool {
	for _, set := range g.Sets {
		if (b.Blues - set.BluesDrawn) < 0 ||
		(b.Greens - set.GreensDrawn) < 0 ||
		(b.Reds - set.RedsDrawn) < 0 {
			return false
		}
	}
	return true
}

func GetSumIdsPossibleGames(doc *bufio.Scanner) int {
	possibleGameIdsSum := 0
	for doc.Scan() {
		line := strings.TrimSpace(doc.Text())
		game := ParseLineToGame(line)
		if StartingBag.IsGamePossible(game) {
			possibleGameIdsSum += game.ID
		}
	}
	// for line in doc:
	//    parse line into game
	//    if Bag.IsGamePossible: sum up ID
	// return sum
	return possibleGameIdsSum
}
