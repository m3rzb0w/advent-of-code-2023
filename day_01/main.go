package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/1/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// part two
var mapNum = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var calibrationValuesPartOne int
var calibrationValuesPartTwo int

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	for _, v := range data {
		partOne := []string{}
		partTwo := []string{}
		for i := range v {
			currentChar := string(v[i])
			if isNumeric(currentChar) {
				partOne = append(partOne, currentChar)
				partTwo = append(partTwo, currentChar) //part two
			}
			//part two
			for key, val := range mapNum {
				if strings.HasPrefix(v[i:], key) {
					partTwo = append(partTwo, val)
				}
			}
		}
		partOneDigit, err := strconv.Atoi(partOne[0] + partOne[len(partOne)-1])
		if err != nil {
			log.Fatal(err)
		}
		partTwoDigit, err := strconv.Atoi(partTwo[0] + partTwo[len(partTwo)-1])
		if err != nil {
			log.Fatal(err)
		}
		calibrationValuesPartOne += partOneDigit
		calibrationValuesPartTwo += partTwoDigit
	}
	fmt.Println("Part One =>", calibrationValuesPartOne)
	fmt.Println("Part Two =>", calibrationValuesPartTwo)
}
