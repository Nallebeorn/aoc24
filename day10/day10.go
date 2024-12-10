package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func PrintTimeSince(start time.Time) {
	fmt.Println(time.Since(start))
}

func SplitByLines(data string) []string {
	normalized := strings.ReplaceAll(data, "\r\n", "\n")
	return strings.Split(normalized, "\n")
}

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return num
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 10")
	partOne()
	partTwo()
}

func partOne() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)
	width := len(lines[0])
	height := len(lines)
	topology := make([]int, 0, width*height)

	for _, line := range lines {
		for _, c := range line {
			topology = append(topology, StrToInt(string(c)))
		}
	}

	score := make([]int, len(topology))
	visited := make([]bool, len(topology))

	for i, h := range topology {
		if h == 9 {
			x := i % width
			y := i / width

			clear(visited)

			var markReachable func(int, int, int)
			markReachable = func(x, y, fromHeight int) {
				idx := y*width + x
				if x < 0 || x >= width || y < 0 || y >= height {
					return
				}

				height := topology[idx]
				if visited[idx] || fromHeight != height+1 {
					return
				}

				visited[idx] = true
				score[idx]++
				markReachable(x-1, y, height)
				markReachable(x+1, y, height)
				markReachable(x, y-1, height)
				markReachable(x, y+1, height)
			}

			markReachable(x, y, 10)
		}
	}

	sum := 0
	for i := range topology {
		if topology[i] == 0 {
			sum += score[i]
		}
	}

	fmt.Println(sum)
}

func partTwo() {
	defer PrintTimeSince(time.Now())
}
