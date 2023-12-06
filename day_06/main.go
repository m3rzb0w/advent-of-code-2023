package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/6/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func main() {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	re := regexp.MustCompile(`\d+`)
	timesStr := re.FindAllString(lines[0], -1)
	DistStr := re.FindAllString(lines[1], -1)

	var times, distances []int

	//part two
	var timeP2str string
	var distP2str string

	for i, num := range timesStr {
		timeNum, err := strconv.Atoi(num)
		if err == nil {
			times = append(times, timeNum)
		}
		distNum, err := strconv.Atoi(DistStr[i])
		if err == nil {
			distances = append(distances, distNum)
		}

		//part two
		timeP2str += num
		distP2str += DistStr[i]

	}

	fmt.Println(times, distances)

	// t = 7
	// d = 9
	//t-i * i > d
	//7-0 * 0 > 9 false
	//(7-1) * 1 = 6 > 9 false
	//(7-2) * 2 = 10 > 9 true
	//...etc

	computeTime := func(t int, d int) int {
		var count int
		for i := 0; i < t; i++ {
			if (t-i)*i > d {
				count++
			}
		}
		return count
	}

	var ans int = 1

	for i, t := range times {
		ans *= computeTime(t, distances[i])

	}
	fmt.Println("Pat one =>", ans)

	timeNumP2, err := strconv.Atoi(timeP2str)
	if err != nil {
		log.Fatal(err)
	}
	distNumP2, err := strconv.Atoi(distP2str)
	if err != nil {
		log.Fatal(err)
	}

	ansP2 := computeTime(timeNumP2, distNumP2)
	fmt.Println("ParTwo =>", ansP2)

}
