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

func StrToInt(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return num
}

type Operator byte

const (
	Start    Operator = '_'
	Add      Operator = '+'
	Multiply Operator = '*'
	Concat   Operator = '|'
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 07")
	partOne()
	partTwo()
}

func verifyEquation(testValue int64, operands []int64, operator Operator, accumulator int64) bool {
	switch operator {
	case Add:
		accumulator += operands[0]
	case Multiply:
		accumulator *= operands[0]
	case Start:
		accumulator = operands[0]
	default:
		panic(fmt.Sprintln("Unknown operator", operator))
	}

	if len(operands) == 1 {
		return accumulator == testValue
	}

	if accumulator > testValue {
		return false
	}

	if verifyEquation(testValue, operands[1:], Add, accumulator) ||
		verifyEquation(testValue, operands[1:], Multiply, accumulator) {
		return true
	}

	return false
}

func partOne() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)

	var result int64

	for _, line := range lines {
		split := strings.Split(line, ":")
		testValue := StrToInt(split[0])
		operands := make([]int64, 0, 16)
		for _, v := range strings.Fields(split[1]) {
			operands = append(operands, StrToInt(v))
		}

		if verifyEquation(testValue, operands, Start, 0) {
			result += testValue
		}
	}

	fmt.Println(result)

}

func verifyEquationWithConcat(testValue int64, operands []int64, operator Operator, accumulator int64) bool {
	switch operator {
	case Add:
		accumulator += operands[0]
	case Multiply:
		accumulator *= operands[0]
	case Start:
		accumulator = operands[0]
	case Concat:
		accumulator = StrToInt(strconv.FormatInt(accumulator, 10) + strconv.FormatInt(operands[0], 10))
	default:
		panic(fmt.Sprintln("Unknown operator", operator))
	}

	if len(operands) == 1 {
		return accumulator == testValue
	}

	if accumulator > testValue {
		return false
	}

	if verifyEquationWithConcat(testValue, operands[1:], Add, accumulator) ||
		verifyEquationWithConcat(testValue, operands[1:], Multiply, accumulator) ||
		verifyEquationWithConcat(testValue, operands[1:], Concat, accumulator) {
		return true
	}

	return false
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	lines := SplitByLines(input)

	var result int64

	for _, line := range lines {
		split := strings.Split(line, ":")
		testValue := StrToInt(split[0])
		operands := make([]int64, 0, 16)
		for _, v := range strings.Fields(split[1]) {
			operands = append(operands, StrToInt(v))
		}

		if verifyEquationWithConcat(testValue, operands, Start, 0) {
			result += testValue
		}
	}

	fmt.Println(result)
}
