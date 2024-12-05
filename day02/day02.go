package main

import (
	_ "embed"
	"fmt"
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

func Sign(x int) int {
	if x >= 0 {
		return 1
	} else {
		return -1
	}
}

//go:embed input.txt
var input string

func main() {
	partOne()
	partTwo()
}

func IsSafe(report []int) bool {
	sign := Sign(report[1] - report[0])
	for i, level := range report[1:] {
		diff := level - report[i]
		if diff*sign < 1 || diff*sign > 3 {
			return false
		}
	}

	return true
}

func partOne() {
	reports := make([][]int, 0)

	for _, line := range SplitByLines(input) {
		report := make([]int, 0)
		for _, level := range strings.Fields(line) {
			report = append(report, StrToInt(level))
		}
		reports = append(reports, report)
	}

	numSafe := 0
	for _, report := range reports {
		if IsSafe(report) {
			numSafe++
		}
	}

	fmt.Println(numSafe)
}

func IsSafeSkipping(source []int, skip int) bool {
	report := make([]int, len(source))
	copy(report, source)
	report = append(report[:skip], report[skip+1:]...)
	return IsSafe(report)
}

func IsSafeDampened(report []int) bool {
	for skip := range report {
		if IsSafeSkipping(report, skip) {
			return true
		}
	}

	return false
}

func partTwo() {
	reports := make([][]int, 0)

	for _, line := range SplitByLines(input) {
		report := make([]int, 0)
		for _, level := range strings.Fields(line) {
			report = append(report, StrToInt(level))
		}
		reports = append(reports, report)
	}

	numSafe := 0
	for _, report := range reports {
		isSafe := IsSafeDampened(report)
		if isSafe {
			numSafe++
		}
	}

	fmt.Println(numSafe)
}
