package main

import (
	"fmt"
	fetch "getdata"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/1/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func main() {
	data := strings.Split(input, "\n")
	fmt.Println(data)
}
