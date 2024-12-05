package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

//go:embed input.txt
var input string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := SplitByLines(input)
	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	for _, line := range lines {
		split := strings.Fields(line)
		leftNums = append(leftNums, StrToInt(split[0]))
		rightNums = append(rightNums, StrToInt(split[1]))
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	totalDistance := 0

	for i := range leftNums {
		totalDistance += Abs(leftNums[i] - rightNums[i])
	}

	fmt.Println("Part one", totalDistance)
}

func Count(list []int, searchFor int) int {
	count := 0

	for _, x := range list {
		if x == searchFor {
			count++
		}
	}

	return count
}

func partTwo() {
	lines := SplitByLines(input)
	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	for _, line := range lines {
		split := strings.Fields(line)
		leftNums = append(leftNums, StrToInt(split[0]))
		rightNums = append(rightNums, StrToInt(split[1]))
	}

	similarityScore := 0
	for _, num := range leftNums {
		similarityScore += num * Count(rightNums, num)
	}

	fmt.Println("Part two", similarityScore)
}
