package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"sort"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/7/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

type Hand struct {
	Cards string
	Bid   int
	Type  int
}

var cardValue = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardValuePartTwo = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 1, // part two joker is weakest
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func parseHands(data []string) []Hand {
	var h Hand
	var allHands []Hand
	for _, v := range data {
		lineSplit := strings.Split(v, " ")
		h.Cards = lineSplit[0]
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		h.Bid = bid
		allHands = append(allHands, h)
	}
	return allHands
}

func (h *Hand) getType(cardValue map[byte]int, part string) int {
	mapVal := map[rune]int{}
	for _, card := range h.Cards {
		mapVal[card]++
	}

	if part == "partTwo" {
		jokerSwap(mapVal)
	}

	val := sort.IntSlice{}
	for _, v := range mapVal {
		val = append(val, v)
	}
	sort.Sort(sort.Reverse(val))
	switch val[0] {
	case 5: // FiveOfAKind
		return 6
	case 4: // FourOfAKind
		return 5
	case 3:
		if val[1] == 2 {
			return 4 // FullHouse
		}
		return 3 // ThreeOfAKind
	case 2:
		if val[1] == 2 {
			return 2 // TwoPair
		}
		return 1 // OnePair
	case 1:
		return 0 // HighCard
	}
	log.Fatal("Err something went wroong")
	return -1
}

func jokerSwap(cardValue map[rune]int) {
	for cardValue['J'] > 0 {
		vMaxNotJoker := -1
		var card rune
		for k, v := range cardValue {
			if k == 'J' {
				continue
			}
			if v > vMaxNotJoker {
				vMaxNotJoker = v
				card = k
			}
		}
		cardValue['J']--
		cardValue[card]++
	}
	// fmt.Println(cardValue)
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	allHands := parseHands(data)
	allHandsTwo := parseHands(data)
	for i := 0; i < len(allHands); i++ {
		allHands[i].Type = allHands[i].getType(cardValue, "partOne")
		allHandsTwo[i].Type = allHands[i].getType(cardValuePartTwo, "partTwo")
	}

	sort.Slice(allHands, func(i, j int) bool {
		if allHands[i].Type < allHands[j].Type {
			return true
		}
		if allHands[i].Type > allHands[j].Type {
			return false
		}

		for c := 0; c < 5; c++ {
			cardOne, cardTwo := cardValue[allHands[i].Cards[c]], cardValue[allHands[j].Cards[c]]
			if cardOne < cardTwo {
				return true
			}
			if cardOne > cardTwo {
				return false
			}
		}
		return false
	})

	sort.Slice(allHandsTwo, func(i, j int) bool {
		if allHandsTwo[i].Type < allHandsTwo[j].Type {
			return true
		}
		if allHandsTwo[i].Type > allHandsTwo[j].Type {
			return false
		}

		for c := 0; c < 5; c++ {
			cardOne, cardTwo := cardValuePartTwo[allHandsTwo[i].Cards[c]], cardValuePartTwo[allHandsTwo[j].Cards[c]]
			if cardOne < cardTwo {
				return true
			}
			if cardOne > cardTwo {
				return false
			}
		}
		return false
	})

	ans := 0
	for i, h := range allHands {
		ans += (i + 1) * h.Bid
	}

	ansTwo := 0
	for i, h := range allHandsTwo {
		ansTwo += (i + 1) * h.Bid
	}

	fmt.Println(ans)
	fmt.Println(ansTwo)

}
