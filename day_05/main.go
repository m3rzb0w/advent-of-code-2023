package main

import (
	"fmt"
	fetch "getdata"
	"log"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/5/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	seeds := []int{}
	soilToFertilizer := [][]int{}
	fertilizerToWater := [][]int{}
	waterToLight := [][]int{}
	lightToTemperature := [][]int{}
	temperatureToHumidity := [][]int{}
	humidityToLocation := [][]int{}

	var currentCat string
	for _, v := range data {
		if strings.Contains(v, "seeds:") {
			s := strings.Split(v, " ")
			for _, v := range s {
				if !strings.Contains(v, "seeds:") {
					seednum, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					seeds = append(seeds, seednum)
				}
			}
		}

		if strings.Contains(v, "seed-to-soil map:") || currentCat == "seed-to-soil map:" {
			currentCat = "seed-to-soil map:"
			s := strings.Split(v, " ")
			tmpsts := []int{}
			for _, v := range s {
				if isNumeric(v) {
					sts, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					tmpsts = append(tmpsts, sts)
				}
			}
			if len(tmpsts) > 0 {
				soilToFertilizer = append(soilToFertilizer, tmpsts)
			}

		}

		if strings.Contains(v, "soil-to-fertilizer map:") || currentCat == "soil-to-fertilizer map:" {
			currentCat = "soil-to-fertilizer map:"
			s := strings.Split(v, " ")
			tmpsts := []int{}
			for _, v := range s {
				if isNumeric(v) {
					sts, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					tmpsts = append(tmpsts, sts)
				}
			}
			if len(tmpsts) > 0 {
				fertilizerToWater = append(fertilizerToWater, tmpsts)
			}

		}

		if strings.Contains(v, "water-to-light map:") || currentCat == "water-to-light map:" {
			currentCat = "water-to-light map:"
			s := strings.Split(v, " ")
			tmpsts := []int{}
			for _, v := range s {
				if isNumeric(v) {
					sts, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					tmpsts = append(tmpsts, sts)
				}
			}
			if len(tmpsts) > 0 {
				waterToLight = append(waterToLight, tmpsts)
			}

		}

		if strings.Contains(v, "light-to-temperature map:") || currentCat == "light-to-temperature map:" {
			currentCat = "light-to-temperature map:"
			s := strings.Split(v, " ")
			tmpsts := []int{}
			for _, v := range s {
				if isNumeric(v) {
					sts, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					tmpsts = append(tmpsts, sts)
				}
			}
			if len(tmpsts) > 0 {
				lightToTemperature = append(lightToTemperature, tmpsts)
			}

		}

		if strings.Contains(v, "temperature-to-humidity map:") || currentCat == "temperature-to-humidity map:" {
			currentCat = "temperature-to-humidity map:"
			s := strings.Split(v, " ")
			tmpsts := []int{}
			for _, v := range s {
				if isNumeric(v) {
					sts, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					tmpsts = append(tmpsts, sts)
				}
			}
			if len(tmpsts) > 0 {
				temperatureToHumidity = append(temperatureToHumidity, tmpsts)
			}

		}

		if strings.Contains(v, "humidity-to-location map:") || currentCat == "humidity-to-location map:" {
			currentCat = "humidity-to-location map:"
			s := strings.Split(v, " ")
			tmpsts := []int{}
			for _, v := range s {
				if isNumeric(v) {
					sts, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					tmpsts = append(tmpsts, sts)
				}
			}
			if len(tmpsts) > 0 {
				humidityToLocation = append(humidityToLocation, tmpsts)
			}

		}

		if len(v) == 0 {
			currentCat = ""
		}

	}

	fmt.Println(seeds)
	fmt.Println(soilToFertilizer)
	fmt.Println(fertilizerToWater)
	fmt.Println(waterToLight)
	fmt.Println(lightToTemperature)
	fmt.Println(temperatureToHumidity)
	fmt.Println(humidityToLocation)

}
