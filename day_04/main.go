package main

import (
	"fmt"
	fetch "getdata"
	"regexp"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/4/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

type Card struct {
	ID         int
	WinningNum []string
	PlayerNum  []string
}

type CardSlice []Card

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	global := 0

	var allCards CardSlice

	matchingNumber := make(map[int]int)

	for _, v := range data {
		var card Card
		lineSplit := strings.Split(v, "|")
		regCard := regexp.MustCompile(`(?:Card\s+)(\d+)`)
		cardId := regCard.FindStringSubmatch(lineSplit[0])[1]
		cardWinningNums := strings.Split(strings.Split(lineSplit[0], ":")[1], " ")
		playerNums := strings.Split(lineSplit[1], " ")
		fmt.Println(cardId, cardWinningNums, playerNums)

		cardIdNum, _ := strconv.Atoi(cardId)

		card.ID = cardIdNum
		card.WinningNum = cardWinningNums
		card.PlayerNum = playerNums
		allCards = append(allCards, card)

		count := []string{}
		tmpcount := 0
		check := make(map[string]bool)
		for _, v := range cardWinningNums {
			for _, n := range playerNums {
				if len(v) != 0 && len(n) != 0 && v == n && check[v] != true {
					fmt.Println("we have a match, ", v, n)
					check[v] = true
					count = append(count, v)
				}
			}
		}
		matchingNumber[cardIdNum]++
		for i := range count {
			matchingNumber[cardIdNum+i+1]++
		}

		fmt.Println("count=>", count)
		if len(count) == 1 {
			tmpcount++
		} else {
			ref := 0
			for i := range count {
				fmt.Println("index ==>", i)
				if ref == 0 {
					ref += i + 1
				} else {
					ref *= 2
				}

			}
			fmt.Println("REFF==>", ref)
			tmpcount += ref
			ref = 0
		}
		// fmt.Println(tmpcount)
		global += tmpcount
		count = []string{}
		check = make(map[string]bool)
	}
	fmt.Println("part1 =>", global)
	fmt.Println(allCards)
	fmt.Println(matchingNumber)
	for _, card := range allCards {
		if matchingNumber[card.ID] != 0 && matchingNumber[card.ID] != 1 {
			loop := 1
			for loop < matchingNumber[card.ID] {

				count := []string{}
				check := make(map[string]bool)
				for _, v := range card.WinningNum {
					for _, n := range card.PlayerNum {
						if len(v) != 0 && len(n) != 0 && v == n && check[v] != true {
							// fmt.Println("we have a match, ", v, n)
							check[v] = true
							count = append(count, v)
						}
					}
				}

				for i := range count {
					matchingNumber[card.ID+i+1]++
				}

				loop++
			}
		}
	}
	fmt.Println(matchingNumber)
	ans := 0
	for _, val := range matchingNumber {
		ans += val

	}
	fmt.Println(ans)
}
