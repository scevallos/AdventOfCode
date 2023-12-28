package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	helpers "AdventOfCode"
)

var (
	lookupTable = map[string]HandType{}
)

type HandType int
type CardRank int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard

	A CardRank = iota
	K
	Q
	T
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	J
)

type Hand struct {
	Cards string // 32T3K
}

type Player struct {
	Hand Hand
	Bid  int
}

func (h Hand) GetType() HandType {
	if val, ok := lookupTable[h.Cards]; ok {
		return val
	}
	cards := map[string]int{}
	for i := 0; i < len(h.Cards); i++ {
		card := fmt.Sprintf("%c", h.Cards[i])
		val, ok := cards[card]
		if ok {
			cards[card] = val + 1
		} else {
			cards[card] = 1
		}
	}

	numJokers, containsJoker := cards["J"]

	var retType HandType

	switch len(cards) {
	case 1:
		retType = FiveOfAKind
	case 2:
		// either 4 of a kind or full house
		maxOccurrences := 0
		for _, numOccurrences := range cards {
			if numOccurrences > maxOccurrences {
				maxOccurrences = numOccurrences
			}
		}
		if maxOccurrences == 4 {
			retType = FourOfAKind
			if containsJoker {
				retType = FiveOfAKind
			}
		} else {
			retType = FullHouse
			if containsJoker {
				retType = FiveOfAKind
			}
		}
	case 3:
		// either 3 of a kind or two pairs
		maxOccurrences := 0
		numPairs := 0
		for _, numOccurrences := range cards {
			if numOccurrences == 2 {
				numPairs++
			}
			if numOccurrences > maxOccurrences {
				maxOccurrences = numOccurrences
			}
		}

		if maxOccurrences == 3 {
			retType = ThreeOfAKind
			if containsJoker {
				retType = FourOfAKind
			}
		} else if numPairs == 2 {
			retType = TwoPair
			if containsJoker {
				if numJokers == 2 {
					retType = FourOfAKind
				} else { // numJokers == 1
					retType = FullHouse
				}
			}
		}
	case 4:
		retType = OnePair
		if containsJoker {
			retType = ThreeOfAKind
		}
	case 5:
		retType = HighCard
		if containsJoker {
			retType = OnePair
		}
	}

	lookupTable[h.Cards] = retType
	return retType
}

func main() {
	doc, closeFile := helpers.GetDocFromFile("sampleInput.txt")
	defer closeFile()

	doc2, closeFile2 := helpers.GetDocFromFile("actualInput.txt")
	defer closeFile2()

	fmt.Println("GetTotalWinnings(sampleInput.txt) =", GetTotalWinnings(doc))
	fmt.Println("GetTotalWinnings(actualInput.txt) =", GetTotalWinnings(doc2))
}

func parseInput(doc *bufio.Scanner) []Player {
	players := []Player{}
	for doc.Scan() {
		handBid := strings.Split(strings.TrimSpace(doc.Text()), " ")
		hand, bid := handBid[0], handBid[1]
		bidInt, err := strconv.Atoi(bid)
		if err != nil {
			panic(err)
		}
		players = append(players, Player{
			Hand: Hand{Cards: hand},
			Bid:  bidInt,
		})
	}

	return players
}

func ConvertToCardRank(in string) CardRank {
	switch in {
	case "A":
		return A
	case "K":
		return K
	case "Q":
		return Q
	case "J":
		return J
	case "T":
		return T
	case "9":
		return Nine
	case "8":
		return Eight
	case "7":
		return Seven
	case "6":
		return Six
	case "5":
		return Five
	case "4":
		return Four
	case "3":
		return Three
	case "2":
		return Two
	default:
		panic("invalid cardRank: " + in)
	}
}

func OrderPlayers(players []Player) {
	sort.SliceStable(players, func(i, j int) bool {
		left := players[i].Hand.GetType()
		right := players[j].Hand.GetType()
		if left != right {
			return left < right
		} else {
			for a := 0; a < 5; a++ {
				leftCardRank := ConvertToCardRank(string(players[i].Hand.Cards[a]))
				rightCardRank := ConvertToCardRank(string(players[j].Hand.Cards[a]))
				if leftCardRank != rightCardRank {
					return leftCardRank < rightCardRank
				} else {
					continue
				}
			}
			panic("exactly the same two hands??")
		}
	})
}

func GetTotalWinnings(doc *bufio.Scanner) int {
	players := parseInput(doc)
	OrderPlayers(players)

	totalWinnings := 0
	for i, player := range players {
		rank := len(players) - i
		// fmt.Printf("hand: %s rank: %v bid: %v\n", player.Hand.Cards, rank, player.Bid)
		winnings := player.Bid * rank
		totalWinnings += winnings
	}
	return totalWinnings
}
