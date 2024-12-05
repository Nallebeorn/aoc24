package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
	partOne()
	partTwo()
}

func partOne() {
	instructionMatcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	instructions := instructionMatcher.FindAllString(input, -1)

	result := 0
	for _, instruction := range instructions {
		args := instruction[4 : len(instruction)-1]
		commaIndex := strings.IndexByte(args, ',')
		arg1 := StrToInt(args[:commaIndex])
		arg2 := StrToInt(args[commaIndex+1:])

		result += arg1 * arg2
	}

	fmt.Println(result)
}

func partTwo() {
	instructionMatcher := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	instructions := instructionMatcher.FindAllString(input, -1)

	result := 0
	enabled := true
	for _, instruction := range instructions {
		bracketIndex := strings.IndexByte(instruction, '(')
		operator := instruction[:bracketIndex]
		if operator == "do" {
			enabled = true
		} else if operator == "don't" {
			enabled = false
		} else if operator == "mul" && enabled {
			args := instruction[bracketIndex+1 : len(instruction)-1]
			commaIndex := strings.Index(args, ",")
			arg1 := StrToInt(args[:commaIndex])
			arg2 := StrToInt(args[commaIndex+1:])

			result += arg1 * arg2
		}
	}

	fmt.Println(result)
}
