package main

import (
	"fmt"
	fetch "getdata"
	"regexp"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/4/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	global := 0

	for _, v := range data {
		lineSplit := strings.Split(v, "|")
		regCard := regexp.MustCompile(`(?:Card\s+)(\d+)`)
		cardId := regCard.FindStringSubmatch(lineSplit[0])[1]
		cardWinningNums := strings.Split(strings.Split(lineSplit[0], ":")[1], " ")
		playerNums := strings.Split(lineSplit[1], " ")
		fmt.Println(cardId, cardWinningNums, playerNums)
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
}
