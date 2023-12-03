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

type Pos struct {
	x, y int
}

type Symbol struct {
	x, y int
	val  string
}

var numsCollected = []string{}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	sympolsPos := []Symbol{}
	gearsPos := []Symbol{}
	for i, v := range data {
		for j, char := range v {
			if !unicode.IsNumber(char) && string(char) != "." {
				symbol := Symbol{x: j, y: i, val: string(char)}
				sympolsPos = append(sympolsPos, symbol)
			}
			//part two
			if string(char) == "*" {
				gear := Symbol{x: j, y: i, val: string(char)}
				gearsPos = append(gearsPos, gear)
			}
		}
	}
	visited := make(map[Pos]bool)
	for _, s := range sympolsPos {
		for _, move := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}} { // down left up right..etc
			j, i := move[1], move[0]
			currentPos := Pos{s.x + j, s.y + i}
			// fmt.Println(currentPos, s.val, s.x, s.y, string(data[currentPos.y][currentPos.x]))
			currentChar := data[currentPos.y][currentPos.x]
			if unicode.IsNumber(rune(currentChar)) {
				tmpDigitCollect := []string{}
				//visited[currentPos] = true
				// fmt.Println(currentPos, s.val, s.x, s.y, string(data[currentPos.y][currentPos.x]))
				for i := currentPos.x; i < len(data[currentPos.y]); i++ {
					newPos := Pos{currentPos.y, i}
					// fmt.Println(string(data[currentPos.y][i]))
					if unicode.IsNumber(rune(data[currentPos.y][i])) && !visited[newPos] {
						visited[newPos] = true
						tmpDigitCollect = append(tmpDigitCollect, string(data[currentPos.y][i]))
					} else {
						break
					}
				}

				for j := currentPos.x - 1; j >= 0; j-- {
					newPosBack := Pos{currentPos.y, j}
					// fmt.Println("TESTBACKWARD===>", j, string(data[currentPos.y][j]), currentPos.x)
					// fmt.Println(currentPos, s.val, s.x, s.y, string(data[currentPos.y][currentPos.x]))
					if unicode.IsNumber(rune(data[currentPos.y][j])) && !visited[newPosBack] {
						visited[newPosBack] = true
						tmpDigitCollect = append([]string{string(data[currentPos.y][j])}, tmpDigitCollect...)
					} else {
						break
					}

				}
				// fmt.Println("tmpdigicolelct", tmpDigitCollect, len(tmpDigitCollect))
				if len(tmpDigitCollect) > 0 {
					fmt.Println(tmpDigitCollect)
					num := strings.Join(tmpDigitCollect, "")
					numsCollected = append(numsCollected, num)
				}
			}
		}
	}
	// fmt.Println(numsCollected)
	sum := 0
	for _, v := range numsCollected {
		// fmt.Println(v)
		if len(v) != 0 {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			sum += n
		}
	}
	fmt.Println(sum)
	// fmt.Println(string(data[0][3]))
	//hith => 533763
	fmt.Println(gearsPos)
	numsGearCollected := []string{}
	visitedGear := make(map[Pos]bool)
	gearCount := make(map[Symbol][]string)
	for _, g := range gearsPos {
		for _, move := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}} { // down left up right..etc
			j, i := move[1], move[0]
			currentPos := Pos{g.x + j, g.y + i}
			// fmt.Println(currentPos, s.val, s.x, s.y, string(data[currentPos.y][currentPos.x]))
			currentChar := data[currentPos.y][currentPos.x]
			if unicode.IsNumber(rune(currentChar)) {
				tmpGearCollect := []string{}
				// fmt.Println(foundGear)
				//visited[currentPos] = true
				// fmt.Println(currentPos, s.val, s.x, s.y, string(data[currentPos.y][currentPos.x]))
				for i := currentPos.x; i < len(data[currentPos.y]); i++ {
					newPos := Pos{currentPos.y, i}
					// fmt.Println(string(data[currentPos.y][i]))
					if unicode.IsNumber(rune(data[currentPos.y][i])) && !visitedGear[newPos] {
						visitedGear[newPos] = true
						tmpGearCollect = append(tmpGearCollect, string(data[currentPos.y][i]))
					} else {
						break
					}
				}

				for j := currentPos.x - 1; j >= 0; j-- {
					newPosBack := Pos{currentPos.y, j}
					// fmt.Println("TESTBACKWARD===>", j, string(data[currentPos.y][j]), currentPos.x)
					// fmt.Println(currentPos, s.val, s.x, s.y, string(data[currentPos.y][currentPos.x]))
					if unicode.IsNumber(rune(data[currentPos.y][j])) && !visitedGear[newPosBack] {
						visitedGear[newPosBack] = true
						tmpGearCollect = append([]string{string(data[currentPos.y][j])}, tmpGearCollect...)
					} else {
						break
					}

				}
				// fmt.Println("tmpdigicolelct", tmpDigitCollect, len(tmpDigitCollect))
				if len(tmpGearCollect) > 0 {
					numGear := strings.Join(tmpGearCollect, "")
					numsGearCollected = append(numsGearCollected, numGear)
					gearCount[g] = append(gearCount[g], numGear)
					fmt.Println(g, numGear)
				}

			}
		}
	}
	fmt.Println(numsGearCollected)
	fmt.Println(gearCount)
	ans := 0
	for _, gear := range gearCount {
		fmt.Println(gear)
		if len(gear) == 2 {
			n1, _ := strconv.Atoi(gear[0])
			n2, _ := strconv.Atoi(gear[1])
			ans += (n1 * n2)
		}
	}
	fmt.Println(ans)
}
