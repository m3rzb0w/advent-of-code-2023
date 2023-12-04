package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var url string = "https://adventofcode.com/2023/day/3/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

type Position struct {
	x, y int
}

type Symbol struct {
	Position
	Value string
}

type SymbolSlice []Symbol

func (symbols *SymbolSlice) FindNumbers(data []string) map[Symbol][]int {
	visited := make(map[Position]bool)
	symbolsNumbers := make(map[Symbol][]int)
	for _, s := range *symbols {
		for _, move := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}} { // 8 axis moves
			j, i := move[0], move[1]

			currentPos := Position{s.x + i, s.y + j}
			currentChar := data[currentPos.y][currentPos.x]

			if unicode.IsNumber(rune(currentChar)) {
				numberCollector := []string{}
				// move forward
				for i := currentPos.x; i < len(data[currentPos.y]); i++ {
					newPosForward := Position{currentPos.y, i}
					if unicode.IsNumber(rune(data[currentPos.y][i])) && !visited[newPosForward] {
						visited[newPosForward] = true
						numberCollector = append(numberCollector, string(data[currentPos.y][i]))
					} else {
						break
					}
				}
				// move backward
				for j := currentPos.x - 1; j >= 0; j-- {
					newPosBackward := Position{currentPos.y, j}
					if unicode.IsNumber(rune(data[currentPos.y][j])) && !visited[newPosBackward] {
						visited[newPosBackward] = true
						numberCollector = append([]string{string(data[currentPos.y][j])}, numberCollector...)
					} else {
						break
					}

				}
				if len(numberCollector) != 0 {
					num, err := strconv.Atoi(strings.Join(numberCollector, ""))
					if err != nil {
						log.Fatal(err)
					}
					symbolsNumbers[s] = append(symbolsNumbers[s], num)
				}
			}
		}
	}
	return symbolsNumbers
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	allSymbols := SymbolSlice{}
	allGears := SymbolSlice{} //part two

	for i, v := range data {
		for j, char := range v {
			if !unicode.IsNumber(char) && string(char) != "." {
				symbol := Symbol{Position: Position{x: j, y: i}, Value: string(char)}
				allSymbols = append(allSymbols, symbol)
			}
			//part two
			if string(char) == "*" {
				gear := Symbol{Position: Position{x: j, y: i}, Value: string(char)}
				allGears = append(allGears, gear)
			}
		}
	}

	allSymbolsNumbers := allSymbols.FindNumbers(data)
	allGearsNumbers := allGears.FindNumbers(data) // part two

	var ansPartOne, ansPartTwo int

	for _, s := range allSymbolsNumbers {
		for _, v := range s {
			ansPartOne += v
		}
	}

	for _, g := range allGearsNumbers { //part two
		if len(g) == 2 {
			ansPartTwo += (g[0] * g[1])
		}
	}
	fmt.Println("Part one =>", ansPartOne)
	fmt.Println("Part two =>", ansPartTwo)
}
