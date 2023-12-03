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
	for i, v := range data {
		for j, char := range v {
			if !unicode.IsNumber(char) && string(char) != "." {
				symbol := Symbol{x: j, y: i, val: string(char)}
				sympolsPos = append(sympolsPos, symbol)
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
}
