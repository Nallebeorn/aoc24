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

//go:embed example1.txt
var input string

func main() {
	fmt.Println("Day 16")
	partOne()
	partTwo()
}

type node struct {
	x, y   int
	facing int
	score  int
}

func partOne() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)
	width := len(lines[0])
	height := len(lines)
	maze := make([]bool, width*height)
	var sx, sy int

	for y, line := range lines {
		for x, c := range line {
			if c == 'E' {
				sx = x
				sy = y
			}
			maze[y*width+x] = c != '#'
		}
	}

	fmt.Println(sx, sy, maze)

}

func partTwo() {
	defer PrintTimeSince(time.Now())
}
