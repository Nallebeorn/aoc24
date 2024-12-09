package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	fmt.Println("Day 09")
	partOne()
	partTwo()
}

func partOne() {
	defer PrintTimeSince(time.Now())

	blocks := make([]int, 0)
	for i, c := range input {
		var id int
		if i%2 == 0 {
			id = i / 2
		} else {
			id = -1
		}

		for j := 0; j < StrToInt(string(c)); j++ {
			blocks = append(blocks, id)
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		last := blocks[i]
		blocks[i] = -1
		firstFree := slices.Index(blocks, -1)
		blocks[firstFree] = last
		if firstFree > i {
			break
		}
	}

	checksum := 0

	for i, id := range blocks {
		if id != -1 {
			checksum += i * id
		} else {
			break
		}
	}

	fmt.Println(checksum)
}

func partTwo() {
	defer PrintTimeSince(time.Now())
}
