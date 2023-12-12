package main

import (
	"fmt"
	fetch "getdata"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/11/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func getEmptyRows(data []string) []int {
	tmp := []int{}
	for i, v := range data {
		isEmpty := true
		for _, c := range v {
			if c != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			tmp = append(tmp, i)
		}
	}
	return tmp
}

func getEmptyCols(data []string) []int {
	tmp := []int{}
	for i := 0; i < len(data[0]); i++ {
		isEmpty := true
		for j := 0; j < len(data); j++ {
			if data[j][i] != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			tmp = append(tmp, i)
		}
	}
	return tmp
}

type Position struct {
	x, y int
}

func getGalaxies(data []string) []Position {
	tmp := []Position{}
	for y, v := range data {
		for x, k := range v {
			if k == '#' {
				p := Position{x: x, y: y}
				tmp = append(tmp, p)
			}
		}
	}
	return tmp
}

func contains(slice []int, v int) bool {
	for _, num := range slice {
		if num == v {
			return true
		}
	}
	return false
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	emptyRows := getEmptyRows(data)
	emptyCols := getEmptyCols(data)
	galaxies := getGalaxies(data)

	fmt.Println(emptyRows)
	fmt.Println(emptyCols)
	fmt.Println(galaxies)

	var ans int
	var ansTwo int

	for i := range galaxies {
		for _, v := range galaxies[:i] {
			minY := min(galaxies[i].y, v.y)
			maxY := max(galaxies[i].y, v.y)
			for i := minY; i < maxY; i++ {
				if contains(emptyRows, i) {
					ans += 2
					ansTwo += 1000000
				} else {
					ans++
					ansTwo++
				}
			}
			minX := min(galaxies[i].x, v.x)
			maxX := max(galaxies[i].x, v.x)
			for j := minX; j < maxX; j++ {
				if contains(emptyCols, j) {
					ans += 2
					ansTwo += 1000000
				} else {
					ans++
					ansTwo++
				}
			}
		}
	}
	fmt.Println(ans)
	fmt.Println(ansTwo)
}
