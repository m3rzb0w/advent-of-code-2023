package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"regexp"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/8/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	if len(integers) < 2 {
		log.Fatal("Slice should contains two integers")
	}
	a := integers[0]
	b := integers[1]
	// fmt.Println(a, b)
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	moves := strings.Split(data[0], "")
	fmt.Println(moves)
	locations := make(map[string][]string)
	regexPattern := `(\w+)\s*=\s*\((\w+),\s*(\w+)\)`
	regex := regexp.MustCompile(regexPattern)
	for _, v := range data {
		if strings.Contains(v, "=") {
			matches := regex.FindStringSubmatch(v)
			k := matches[1]
			l := matches[2]
			r := matches[3]
			locations[k] = append(locations[k], l, r)
		}
	}
	// fmt.Println(locations)
	// fmt.Println(len(moves))
	var nextLoc = "AAA" //start position part one
	var ans int
	if len(locations[nextLoc]) != 0 {
		for i := 0; i < len(moves); i++ {
			// fmt.Println(i)
			var nextMoveVal int
			if moves[i] == "L" {
				nextMoveVal = 0
			} else if moves[i] == "R" {
				nextMoveVal = 1
			}
			nextLoc = locations[nextLoc][nextMoveVal]
			// fmt.Println(nextLoc)
			ans++
			if nextLoc == "ZZZ" {
				break
			}
			// fmt.Println("current i", i)
			if i == len(moves)-1 {
				// fmt.Println("current i", i)
				i = -1
			}
		}
	}

	fmt.Println("Part One =>", ans)
	// fmt.Println(locations)

	partTwoStartPos := []string{}
	for k := range locations { // <== this induces random start position in the slice therefore each associated cycles are changing position in their slice every run. It doesn't affect the result
		if k[2] == 'A' {
			partTwoStartPos = append(partTwoStartPos, k)
		}
	}
	fmt.Println(partTwoStartPos)
	cycles := make([]int, len(partTwoStartPos))
	fmt.Println(cycles)

	for j, pos := range partTwoStartPos {
		nextLoc := pos
		for i := 0; i < len(moves); i++ {

			var nextMoveVal int
			if moves[i] == "L" {
				nextMoveVal = 0
			} else if moves[i] == "R" {
				nextMoveVal = 1
			}
			nextLoc = locations[nextLoc][nextMoveVal]
			// fmt.Println(pos, moves[i], nextMoveVal, nextLoc)
			cycles[j]++
			if nextLoc[2] == 'Z' {
				break
			}
			if i == len(moves)-1 {
				i = -1
			}
		}

	}
	fmt.Println("cycles", cycles)
	fmt.Println("Part Two =>", LCM(cycles...))
}
