package main

import (
	"fmt"
	fetch "getdata"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/5/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func main() {

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	seedStr := strings.Split(lines[0], " ")[1:]
	seeds := make([]int, len(seedStr))
	for i, s := range seedStr {
		seeds[i], _ = strconv.Atoi(s)

	}

	fmt.Println(seeds)

	maps := make([][][]int, 0)
	i := 2
	for i < len(lines) {
		m := make([][]int, 0)

		i++
		for i < len(lines) && lines[i] != "" {
			rangeData := strings.Split(lines[i], " ")
			dstStart, _ := strconv.Atoi(rangeData[0])
			srcStart, _ := strconv.Atoi(rangeData[1])
			rangeLen, _ := strconv.Atoi(rangeData[2])
			m = append(m, []int{dstStart, srcStart, rangeLen})
			i++
		}

		maps = append(maps, m)
		i += 1
	}

	// fmt.Println("MAPPS =>>", maps)

	findLoc := func(seed int) int {
		curNum := seed

		for _, m := range maps {
			// fmt.Println("current M ==>", m)
			for _, rangeData := range m {
				// fmt.Println("current range data====>", rangeData)
				dstStart := rangeData[0]
				srcStart := rangeData[1]
				rangeLen := rangeData[2]
				// fmt.Println(dstStart, srcStart, rangeLen)
				if srcStart <= curNum && curNum < srcStart+rangeLen {
					curNum = dstStart + (curNum - srcStart)
					// fmt.Println("currr =====>", curNum)
					break
				}
			}
		}
		return curNum
	}

	locs := make([]int, 0)
	for _, seed := range seeds {
		loc := findLoc(seed)
		locs = append(locs, loc)
	}

	minLoc := locs[0]
	for _, loc := range locs {
		if loc < minLoc {
			minLoc = loc
		}
	}

	fmt.Println("MinLOCone =>", minLoc)

	var refLoc int
	// locsTwo := make([]int, 0)
	for i, seed := range seeds {
		if i%2 == 0 {
			currSeed := seed
			currLoop := seeds[i+1] + currSeed
			// fmt.Println(currSeed, currLoop)
			for currSeed < currLoop {
				// fmt.Println(currSeed)
				loc := findLoc(currSeed)
				// locsTwo = append(locsTwo, loc)
				if loc < refLoc || refLoc == 0 {
					refLoc = loc
				}
				currSeed++
			}
		}
	}
	fmt.Println(refLoc)
	// minLocTwo := locsTwo[0]
	// for _, loc := range locsTwo {
	// 	if loc < minLocTwo {
	// 		minLocTwo = loc
	// 	}
	// }
	// fmt.Println("mintwo=>", minLocTwo)

}
