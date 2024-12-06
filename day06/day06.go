package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

func ReplaceAt(str string, index int, new rune) string {
	return str[:index] + string(new) + str[index+1:]
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

func isLoop(lines []string, guardX int, guardY int, obsX int, obsY int) bool {
	visitedWithDirections := make(map[[2]int]int)

	guardDirection := 0 // 0 ^, 1 >, 2 v, 3 <
	width := len(lines[0])
	height := len(lines)
	for i := 0; i < 100000000; i++ {
		if dir, beenHere := visitedWithDirections[[2]int{guardX, guardY}]; beenHere && dir&(1<<guardDirection) != 0 {
			return true
		}

		visitedWithDirections[[2]int{guardX, guardY}] |= 1 << guardDirection

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
			return false
		}

		if lines[newY][newX] == '#' || newX == obsX && newY == obsY {
			guardDirection = (guardDirection + 1) % 4
		} else {
			guardX, guardY = newX, newY
		}
	}

	fmt.Println("Time paradox detected when placing obstacle at", obsX, obsY)
	return false
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)

	guardX := 0
	guardY := 0

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

	var numPossibilities atomic.Uint64
	var wg sync.WaitGroup

	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				wg.Add(1)
				go func(x, y int) {
					defer wg.Done()

					if isLoop(lines, guardX, guardY, x, y) {
						numPossibilities.Add(1)
					}
				}(x, y)
			}
		}
	}

	wg.Wait()

	fmt.Println(numPossibilities.Load())
}
