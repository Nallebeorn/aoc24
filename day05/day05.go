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
	fmt.Println("Day 05")
	partOne()
	partTwo()
}

func partOne() {
	lines := SplitByLines(input)

	pagesThatMustBeBefore := make(map[int][]int)

	section := 0

	result := 0

	for _, line := range lines {
		if line == "" {
			section = 1
			continue
		}

		if section == 0 {
			pages := strings.Split(line, "|")
			before := StrToInt(pages[0])
			after := StrToInt(pages[1])
			pagesThatMustBeBefore[after] = append(pagesThatMustBeBefore[after], before)

		} else if section == 1 {
			split := strings.Split(line, ",")
			pages := make([]int, len(split))
			for i, page := range split {
				pages[i] = StrToInt(page)
			}

			if func() bool {
				for index, page := range pages {
					for _, mustBeBefore := range pagesThatMustBeBefore[page] {
						if slices.Index(pages, mustBeBefore) > index {
							return false
						}
					}
				}

				return true
			}() {
				middle := pages[len(pages)/2]
				result += middle
			}
		}
	}

	fmt.Println(result)

}

func partTwo() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)

	pagesThatMustBeBefore := make(map[int][]int)

	section := 0

	result := 0

	for _, line := range lines {
		if line == "" {
			section = 1
			continue
		}

		if section == 0 {
			pages := strings.Split(line, "|")
			before := StrToInt(pages[0])
			after := StrToInt(pages[1])
			pagesThatMustBeBefore[after] = append(pagesThatMustBeBefore[after], before)

		} else if section == 1 {
			split := strings.Split(line, ",")
			pages := make([]int, len(split))
			for i, page := range split {
				pages[i] = StrToInt(page)
			}

			wasIncorrect := false
			for index := 0; index < len(pages); index++ {
				page := pages[index]
				for _, mustBeBefore := range pagesThatMustBeBefore[page] {
					beforeIndex := slices.Index(pages, mustBeBefore)
					if beforeIndex > index {
						pages = slices.Delete(pages, beforeIndex, beforeIndex+1)
						pages = slices.Insert(pages, index, mustBeBefore)
						index = -1
						wasIncorrect = true
						break
					}
				}
			}

			if wasIncorrect {
				middle := pages[len(pages)/2]
				result += middle
			}
		}
	}

	fmt.Println(result)
}
