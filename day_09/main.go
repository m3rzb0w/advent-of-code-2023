package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/9/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func parseNumber(data []string) [][]int {
	allNumbers := [][]int{}
	for _, v := range data {
		nums := strings.Fields(v)
		lineNumbers := []int{}
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			lineNumbers = append(lineNumbers, num)
		}
		allNumbers = append(allNumbers, lineNumbers)
	}
	return allNumbers
}

func isFullZeros(slice []int) bool {
	zeroSlice := make([]int, len(slice))
	return reflect.DeepEqual(slice, zeroSlice)
}

func numSeq(line []int) [][]int {
	var res []int
	for i := 1; i < len(line); i++ {
		substractNum := line[i] - line[i-1]
		res = append(res, substractNum)

	}
	// fmt.Println(res)
	lineSeq = append(lineSeq, res)
	if !isFullZeros(res) {
		numSeq(res)
	}
	return lineSeq
}

func findDigit(sequence [][]int) int {
	var num int
	for i := len(sequence) - 1; i >= 0; i-- {
		lastNum := sequence[i][len(sequence[i])-1]
		num += lastNum
	}
	return num
}

func findDigitPartTwo(sequence [][]int) int {
	var num int
	for i := len(sequence) - 1; i >= 0; i-- {
		firstNum := sequence[i][0]
		num = firstNum - num

	}
	return num
}

var lineSeq [][]int
var allSequences [][][]int

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	allNumbers := parseNumber(data)
	for _, line := range allNumbers {
		lineSeq = append(lineSeq, line)
		tmp := numSeq(line)
		allSequences = append(allSequences, tmp)
		lineSeq = [][]int{}
	}

	var ans int
	var ansTwo int
	for _, v := range allSequences {
		ans += findDigit(v)
		ansTwo += findDigitPartTwo(v)
	}
	fmt.Println("Part one =>", ans)
	fmt.Println("Part two =>", ansTwo)
}
