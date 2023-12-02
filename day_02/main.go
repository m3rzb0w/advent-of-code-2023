package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/2/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var countPartOne, countPartTwo int

// part 1
const maxRedCubesPartOne int = 12
const maxGreenCubesPartOne int = 13
const maxBlueCubesPartOne int = 14

// part 2
var fewRed, fewGreen, fewBlue int

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	for _, v := range data {
		currentGame := strings.Split(v, ";")
		re := regexp.MustCompile(`(?:Game\s)(\d+)`)
		reRed := regexp.MustCompile(`(\d+)(?:\sred)`)
		reGreen := regexp.MustCompile(`(\d+)(?:\sgreen)`)
		reBlue := regexp.MustCompile(`(\d+)(?:\sblue)`)
		gameId, err := strconv.Atoi(re.FindStringSubmatch(currentGame[0])[1])
		if err != nil {
			log.Fatal(err)
		}

		for _, set := range currentGame {
			var currentRed int
			var currentGreen int
			var currentBlue int
			if reRed.MatchString(set) {
				currentRed, _ = strconv.Atoi(reRed.FindStringSubmatch(set)[1])
			}
			if reGreen.MatchString(set) {
				currentGreen, _ = strconv.Atoi(reGreen.FindStringSubmatch(set)[1])
			}
			if reBlue.MatchString(set) {
				currentBlue, _ = strconv.Atoi(reBlue.FindStringSubmatch(set)[1])
			}

			if currentRed > maxRedCubesPartOne || currentGreen > maxGreenCubesPartOne || currentBlue > maxBlueCubesPartOne {
				gameId = 0

			}
			//part two
			if currentRed > fewRed {
				fewRed = currentRed
			}

			if currentBlue > fewBlue {
				fewBlue = currentBlue
			}

			if currentGreen > fewGreen {
				fewGreen = currentGreen
			}
		}
		countPartOne += gameId
		countPartTwo += fewRed * fewBlue * fewGreen
		fewRed = 0
		fewBlue = 0
		fewGreen = 0
	}
	fmt.Println("Part one =>", countPartOne)
	fmt.Println("Part two =>", countPartTwo)
}
