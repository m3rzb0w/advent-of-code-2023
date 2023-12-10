package main

import (
	"fmt"
	fetch "getdata"
	"strings"
)

var url string = "https://adventofcode.com/2023/day/10/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

func parseMaze(data []string) PipesMaze {
	maze := PipesMaze{}
	for y, v := range data {
		pipes := strings.Split(v, "")
		linePipes := []Pipe{}
		for x, p := range pipes {
			var tmp = Pipe{
				Position: Position{x: x, y: y}, Char: p,
			}
			linePipes = append(linePipes, tmp)
		}
		maze = append(maze, linePipes)
	}
	return maze
}

type Position struct {
	x, y int
}

type Pipe struct {
	Position
	Char string
}

type PipesMaze [][]Pipe

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	pipesMaze := parseMaze(data)
	visited := make(map[Position]bool)
	var startPosition Position
	for _, c := range pipesMaze {
		for _, r := range c {
			if r.Char == "S" {
				startPosition = r.Position
			}
		}
	}
	fmt.Println(startPosition, visited)
	visited[startPosition] = true
	q := []Position{startPosition}
	for len(q) != 0 {
		// fmt.Println("Queue size =>", len(q))
		current := q[0]
		char := pipesMaze[current.y][current.x].Char
		q = q[1:] //dequeu
		// fmt.Println(current, char)
		//up
		if current.y > 0 && strings.Contains("S|JL", char) && strings.Contains("|7F", pipesMaze[current.y-1][current.x].Char) && !visited[Position{current.x, current.y - 1}] {
			visited[Position{current.x, current.y - 1}] = true
			q = append(q, Position{current.x, current.y - 1})
		}
		//down
		if current.y < len(pipesMaze)-1 && strings.Contains("S|7F", char) && strings.Contains("|JL", pipesMaze[current.y+1][current.x].Char) && !visited[Position{current.x, current.y + 1}] {
			visited[Position{current.x, current.y + 1}] = true
			q = append(q, Position{current.x, current.y + 1})
		}
		//left
		if current.x > 0 && strings.Contains("S-J7", char) && strings.Contains("-LF", pipesMaze[current.y][current.x-1].Char) && !visited[Position{current.x - 1, current.y}] {
			visited[Position{current.x - 1, current.y}] = true
			q = append(q, Position{current.x - 1, current.y})
		}
		//right
		if current.x < len(pipesMaze[0]) && strings.Contains("S-LF", char) && strings.Contains("-J7", pipesMaze[current.y][current.x+1].Char) && !visited[Position{current.x + 1, current.y}] {
			visited[Position{current.x + 1, current.y}] = true
			q = append(q, Position{current.x + 1, current.y})
		}
	}
	fmt.Println(len(visited) / 2)

}
