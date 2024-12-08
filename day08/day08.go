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
	fmt.Println("Day 08")
	partOne()
	partTwo()
}

func partOne() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)

	width := len(lines[0])
	height := len(lines)

	antennas := make(map[byte][][2]int)

	for y, line := range lines {
		for x, c := range []byte(line) {
			if c != '.' {
				antennas[c] = append(antennas[c], [2]int{x, y})
			}
		}
	}

	antinodes := make(map[[2]int]bool, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for freq, positions := range antennas {
				for _, position := range positions {
					dx := position[0] - x
					dy := position[1] - y

					if dx == 0 && dy == 0 {
						continue
					}

					otherX := x + dx*2
					otherY := y + dy*2

					if otherX >= 0 && otherX < width && otherY >= 0 && otherY < height {
						if lines[otherY][otherX] == freq {
							antinodes[[2]int{x, y}] = true
						}
					}
				}
			}
		}
	}

	for y, line := range lines {
		for x, c := range line {
			if antinodes[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Print("\n")
	}

	fmt.Println(len(antinodes))

}

func partTwo() {
}
