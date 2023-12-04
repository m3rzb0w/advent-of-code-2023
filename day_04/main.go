package main

import (
	"fmt"
	fetch "getdata"
)

var url string = "https://adventofcode.com/2023/day/4/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func main() {
	fmt.Println("hey mum")
}
