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

type point struct {
	x, y int
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	walls := make(map[point]bool)
	boxes := make([]point, 0)
	movements := make([]rune, 0)

	hasLeftBox := func(x, y int) bool {
		for _, box := range boxes {
			if box.y == y && box.x == x {
				return true
			}
		}

		return false
	}

	findBox := func(x, y int) (int, bool) {
		for i, box := range boxes {
			if box.y == y && (box.x == x || box.x+1 == x) {
				return i, true
			}
		}

		return 0, false
	}

	var robotx, roboty int

	parseStep := 0
	lines := SplitByLines(input)
	width := len(lines[0]) * 2
	var height int

	for y, line := range lines {
		if line == "" {
			parseStep = 1
		} else if parseStep == 0 {
			height++
			for x, c := range line {
				if c == '@' {
					robotx = x * 2
					roboty = y
				} else if c == 'O' {
					boxes = append(boxes, point{x * 2, y})
				} else if c == '#' {
					walls[point{x * 2, y}] = true
					walls[point{x*2 + 1, y}] = true
				}
			}
		} else if parseStep == 1 {
			for _, c := range line {
				movements = append(movements, c)
			}
		}
	}

	printState := func() {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if x == robotx && y == roboty {
					fmt.Print("@")
				} else if walls[point{x, y}] {
					fmt.Print("#")
				} else if hasLeftBox(x, y) {
					fmt.Print("[")
				} else if hasLeftBox(x-1, y) {
					fmt.Print("]")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}

		fmt.Print("\n\n")
	}

	printState()

	fmt.Println("num boxes", len(boxes))

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

		new := point{robotx + dx, roboty + dy}
		pushed := make([]int, 0)

		if dy == 0 {
			for x := new.x; true; x += dx * 2 {
				boxId, isBox := findBox(x, new.y)
				if walls[point{x, new.y}] {
					break
				} else if isBox {
					if !slices.Contains(pushed, boxId) {
						pushed = append(pushed, boxId)
					}
				} else {
					robotx = new.x
					for _, b := range pushed {
						boxes[b].x += dx
					}
					break
				}
			}
		} else {
			if walls[new] {
				goto stop
			} else if boxId, isBox := findBox(new.x, new.y); isBox {
				pushed = append(pushed, boxId)

				for b := 0; b < len(pushed); b++ {
					box := boxes[pushed[b]]
					if walls[point{box.x, box.y + dy}] || walls[point{box.x + 1, box.y + dy}] {
						goto stop
					}

					if boxId, isBox = findBox(box.x, box.y+dy); isBox {
						if !slices.Contains(pushed, boxId) {
							pushed = append(pushed, boxId)
						}
					}
					if boxId, isBox = findBox(box.x+1, box.y+dy); isBox {
						if !slices.Contains(pushed, boxId) {
							pushed = append(pushed, boxId)
						}
					}
				}
			}

			roboty = new.y
			for _, b := range pushed {
				boxes[b].y += dy
			}
		}

	stop:
		// fmt.Println(i, string(move))
		// printState()

		for i, box := range boxes {
			if box.x < 2 || box.x > width-3 || box.y < 1 || box.y > height-2 {
				fmt.Println("oob box", i)
				panic("wtf")
			}
		}
	}

	sum := 0
	for _, box := range boxes {
		gps := box.y*100 + box.x
		sum += gps
	}

	fmt.Println(sum)
}
