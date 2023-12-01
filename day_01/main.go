package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/1/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

// part two
var numDict = map[string]string{
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

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	calibrationValuePartOne := 0
	calibrationValuePartTwo := 0
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	regTwo, err := regexp.Compile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		tmp := strings.Split(reg.ReplaceAllString(v, ""), "")
		tmpPartTwo := regTwo.FindAllString(v, -1)
		if len(tmp) > 1 {
			firstDigit := tmp[0]
			lastDigit := tmp[len(tmp)-1]
			digit, _ := strconv.Atoi(firstDigit + lastDigit)
			calibrationValuePartOne += digit
		} else {
			firstDigit := tmp[0]
			digit, _ := strconv.Atoi(firstDigit + firstDigit)
			calibrationValuePartOne += digit
		}
		//par two
		if len(tmpPartTwo) > 1 {
			firstDigitPartTwoStr := tmpPartTwo[0]
			lastDigitPartTwoStr := tmpPartTwo[len(tmpPartTwo)-1]

			if val, ok := numDict[firstDigitPartTwoStr]; ok {
				firstDigitPartTwoStr = val
			}

			if val, ok := numDict[lastDigitPartTwoStr]; ok {
				lastDigitPartTwoStr = val
			}
			digit, _ := strconv.Atoi(firstDigitPartTwoStr + lastDigitPartTwoStr)
			calibrationValuePartTwo += digit
		} else {
			firstDigitPartTwoStr := tmpPartTwo[0]
			digit, _ := strconv.Atoi(firstDigitPartTwoStr + firstDigitPartTwoStr)
			calibrationValuePartTwo += digit
		}
		fmt.Println(v, tmpPartTwo)
	}
	//part two

	fmt.Println("Part One =>", calibrationValuePartOne)
	fmt.Println("Part Two =>", calibrationValuePartTwo)
}
