package main

import (
	_ "embed"
	"fmt"
	"slices"
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

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 11")
	partOne()
	partTwo()
}

func partOne() {
	const numBlinks = 25

	defer PrintTimeSince(time.Now())

	fields := strings.Fields(input)
	stones := make([]int, len(fields))
	for i, str := range fields {
		stones[i] = StrToInt(str)
	}

	for blink := 0; blink < numBlinks; blink++ {
		fmt.Println(blink, len(stones))
		for i := 0; i < len(stones); i++ {
			x := stones[i]
			if x == 0 {
				stones[i] = 1
			} else if xstr := strconv.Itoa(x); len(xstr)%2 == 0 {
				left := StrToInt(xstr[:len(xstr)/2])
				right := StrToInt(xstr[len(xstr)/2:])
				stones[i] = left
				stones = slices.Insert(stones, i+1, right)
				i++
			} else {
				stones[i] = x * 2024
			}
		}
	}

	// fmt.Println(stones)
	fmt.Println(len(stones))
}

var memo = sync.Map{}

func recurse(x int, count int, depth int) int {
	if depth >= 75 {
		return count
	}

	if ret, ok := memo.Load([3]int{x, count, depth}); ok {
		return ret.(int)
	}

	var result int

	if x == 0 {
		result = recurse(1, count, depth+1)
	} else if xstr := strconv.Itoa(x); len(xstr)%2 == 0 {
		left := StrToInt(xstr[:len(xstr)/2])
		right := StrToInt(xstr[len(xstr)/2:])
		result = recurse(left, 1, depth+1) + recurse(right, 1, depth+1)
	} else {
		result = recurse(x*2024, count, depth+1)
	}

	memo.Store([3]int{x, count, depth}, result)

	return result
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	fields := strings.Fields(input)
	stones := make([]int, len(fields))
	for i, str := range fields {
		stones[i] = StrToInt(str)
	}

	var count atomic.Int64
	var wg sync.WaitGroup

	for _, engraving := range stones {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			count.Add(int64(recurse(x, 1, 0)))
		}(engraving)
	}

	wg.Wait()

	fmt.Println(count.Load())
}
