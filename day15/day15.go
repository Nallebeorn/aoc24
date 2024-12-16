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

//go:embed example2.txt
var input string

func main() {
	fmt.Println("Day 15")
	partOne()
	partTwo()
}

func partOne() {
	defer PrintTimeSince(time.Now())

	warehouse := make([][]rune, 0)
	movements := make([]rune, 0)

	var robotx, roboty int

	parseStep := 0

	for y, line := range SplitByLines(input) {
		if line == "" {
			parseStep = 1
		} else if parseStep == 0 {
			row := make([]rune, 0, len(line))
			for x, c := range line {
				if c == '@' {
					robotx = x
					roboty = y
					row = append(row, '.')
				} else {
					row = append(row, c)
				}
			}
			warehouse = append(warehouse, row)
		} else if parseStep == 1 {
			for _, c := range line {
				movements = append(movements, c)
			}
		}
	}

	for _, move := range movements {
		var dx, dy int
		switch move {
		case '>':
			dx = 1
		case '^':
			dy = -1
		case '<':
			dx = -1
		case 'v':
			dy = 1
		default:
			panic("Invalid move " + string(move))
		}

		var pushBox func(x, y int) bool
		pushBox = func(x, y int) bool {
			switch warehouse[y+dy][x+dx] {
			case '#':
				return false
			case 'O':
				return pushBox(x+dx, y+dy)
			case '.':
				warehouse[y+dy][x+dx] = 'O'
				return true
			}

			panic("Invalid object found")
		}

		switch warehouse[roboty+dy][robotx+dx] {
		case 'O':
			if pushBox(robotx+dx, roboty+dy) {
				robotx += dx
				roboty += dy
				warehouse[roboty][robotx] = '.'
			}
		case '.':
			robotx += dx
			roboty += dy
		case '#':
			continue
		default:
			panic("Invalid object found")
		}
	}

	sum := 0
	for y, row := range warehouse {
		for x, c := range row {
			if c == 'O' {
				gps := y*100 + x
				sum += gps
			}
		}
	}

	fmt.Println(sum)
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	warehouse := make([][]rune, 0)
	movements := make([]rune, 0)

	var robotx, roboty int

	printState := func() {
		for y, row := range warehouse {
			for x, c := range row {
				if x == robotx && y == roboty {
					fmt.Print("@")
				} else {
					fmt.Print(string(c))
				}
			}
			fmt.Print("\n")
		}

		fmt.Print("\n\n")
	}

	removeBox := func(x, y int, c rune) {
		warehouse[y][x] = '.'
		if c == ']' {
			warehouse[y][x-1] = '.'
		} else if c == '[' {
			warehouse[y][x+1] = '.'
		} else {
			panic("Not a box to remove " + string(c))
		}
	}

	putBox := func(x, y int, c rune) {
		warehouse[y][x] = c
		if c == ']' {
			warehouse[y][x-1] = '['
		} else if c == '[' {
			warehouse[y][x+1] = ']'
		} else {
			panic("Not a box to put " + string(c))
		}
	}

	parseStep := 0

	for y, line := range SplitByLines(input) {
		if line == "" {
			parseStep = 1
		} else if parseStep == 0 {
			row := make([]rune, 0, len(line))
			for x, c := range line {
				if c == '@' {
					robotx = x * 2
					roboty = y
					row = append(row, '.', '.')
				} else if c == 'O' {
					row = append(row, '[', ']')
				} else {
					row = append(row, c, c)
				}
			}
			warehouse = append(warehouse, row)
		} else if parseStep == 1 {
			for _, c := range line {
				movements = append(movements, c)
			}
		}
	}

	printState()

	for _, move := range movements {
		var dx, dy int
		switch move {
		case '>':
			dx = 1
		case '^':
			dy = -1
		case '<':
			dx = -1
		case 'v':
			dy = 1
		default:
			panic("Invalid move " + string(move))
		}

		atNew := warehouse[roboty+dy][robotx+dx]

		var pushBox func(x, y int) bool
		pushBox = func(x, y int) bool {
			atHere := warehouse[y][x]
			atNext := warehouse[y+dy][x+dx]
			// fmt.Println("Push", x, y, string(atHere), string(atNext))
			if atNext == '#' {
				return false
			} else if atNext == '.' {
				// fmt.Println("EOl")
				removeBox(x, y, atHere)
				putBox(x+dx, y+dy, atHere)
				return true
			} else {
				fmt.Println("next", string(atNext))
				if dy == 0 {
					if pushBox(x+dx, y+dy) {
						// fmt.Println("pushy push")
						removeBox(x, y, atHere)
						putBox(x+dx, y+dy, atHere)
						return true
					} else {
						return false
					}
				} else {
					a := pushBox(x, y+dy)
					var b bool
					if warehouse[y][x] == '[' {
						b = pushBox(x+1, y+dy)
					} else if warehouse[y][x] == ']' {
						b = pushBox(x-1, y+dy)
					} else {
						panic("AAAAAaaa")
					}
					if a && b {
						// fmt.Println("double push")
						removeBox(x, y, atHere)
						putBox(x+dx, y+dy, atHere)
						return true
					}
					return false
				}
			}
		}

		switch atNew {
		case '[', ']':
			if pushBox(robotx+dx, roboty+dy) {
				robotx += dx
				roboty += dy
				warehouse[roboty][robotx] = '.'
				if dy != 0 {
					if atNew == '[' {
						warehouse[roboty][robotx+1] = '.'
					} else if atNew == ']' {
						warehouse[roboty][robotx-1] = '.'
					} else {
						panic("Aaaasdadad")
					}
				}
			}
		case '.':
			robotx += dx
			roboty += dy
		case '#':
			continue
		default:
			panic("Invalid object found")
		}

		fmt.Println(string(move))
		printState()
	}

	sum := 0
	for y, row := range warehouse {
		for x, c := range row {
			if c == '[' {
				gps := y*100 + x
				sum += gps
			}
		}
	}

	fmt.Println(sum)
}
