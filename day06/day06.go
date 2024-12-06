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
	fmt.Println("Day 06")
	partOne()
	partTwo()
}

func partOne() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)

	guardX := 0
	guardY := 0
	guardDirection := 0 // 0 ^, 1 >, 2 v, 3 <

findStartingPos:
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				guardX = x
				guardY = y
				break findStartingPos
			}
		}
	}

	visited := make(map[[2]int]bool)

	width := len(lines[0])
	height := len(lines)
	for {
		visited[[2]int{guardX, guardY}] = true

		newX, newY := guardX, guardY

		switch guardDirection {
		case 0:
			newY--
		case 1:
			newX++
		case 2:
			newY++
		case 3:
			newX--
		}

		if newX < 0 || newX > width-1 || newY < 0 || newY > height-1 {
			break
		}

		if lines[newY][newX] == '#' {
			guardDirection = (guardDirection + 1) % 4
		} else {
			guardX, guardY = newX, newY
		}
	}

	fmt.Println(len(visited))
}

func partTwo() {

}
