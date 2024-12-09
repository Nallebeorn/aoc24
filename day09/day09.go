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

type block struct {
	id   int
	size int
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	blocks := make([]block, 0)
	for i, c := range input {
		var id int
		if i%2 == 0 {
			id = i / 2
		} else {
			id = -1
		}

		size := StrToInt(string(c))
		blocks = append(blocks, block{id, size})
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		blk := blocks[i]

		if blk.id == -1 {
			continue
		}

		firstFree := slices.IndexFunc(blocks, func(b block) bool {
			return b.id == -1 && b.size >= blk.size
		})

		if firstFree != -1 && firstFree < i {
			remainingSpace := blocks[firstFree].size - blk.size
			blocks[i].id = -1
			blocks[firstFree] = blk
			if remainingSpace > 0 {
				blocks = slices.Insert(blocks, firstFree+1, block{-1, remainingSpace})
			}
		}
	}

	// fmt.Println(blocks)

	checksum := 0

	i := 0
	for _, block := range blocks {
		for j := 0; j < block.size; j++ {
			if block.id != -1 {
				checksum += i * block.id
			}
			i++
		}
	}

	fmt.Println(checksum)
}
