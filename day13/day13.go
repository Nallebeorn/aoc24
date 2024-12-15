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

//go:embed input.txt
var input string

type machine struct {
	ax, ay         int
	bx, by         int
	prizex, prizey int
}

func main() {
	fmt.Println("Day 13")
	partOne()
	partTwo()
}

func getCost(machine machine) int {
	maxApresses := min(machine.prizex/machine.ax, machine.prizey/machine.ay, 100)
	maxBpresses := min(machine.prizex/machine.bx, machine.prizey/machine.by, 100)

	best := -1
	for a := 0; a <= maxApresses; a++ {
		for b := 0; b <= maxBpresses; b++ {
			x := machine.ax*a + machine.bx*b
			y := machine.ay*a + machine.by*b
			if x == machine.prizex && y == machine.prizey {
				cost := a*3 + b
				if best == -1 || cost < best {
					best = cost
				}
			} else if x > machine.prizex || y > machine.prizey {
				break
			}
		}
	}

	return best
}

func partOne() {
	defer PrintTimeSince(time.Now())

	machines := make([]machine, 0)

	lines := SplitByLines(input)
	for i := 0; i < len(lines); i += 4 {
		var machine machine
		var fields []string
		var field string

		fields = strings.Fields(lines[i])
		field = fields[2]
		machine.ax = StrToInt(field[2 : len(field)-1])
		machine.ay = StrToInt(fields[3][2:])

		fields = strings.Fields(lines[i+1])
		field = fields[2]
		machine.bx = StrToInt(field[2 : len(field)-1])
		machine.by = StrToInt(fields[3][2:])

		fields = strings.Fields(lines[i+2])
		field = fields[1]
		machine.prizex = StrToInt(field[2 : len(field)-1])
		machine.prizey = StrToInt(fields[2][2:])

		machines = append(machines, machine)
	}

	sum := 0
	for _, machine := range machines {
		cost := getCost(machine)
		if cost >= 0 {
			sum += cost
		}
	}

	fmt.Println(sum)
}

func getCost2(machine machine) (int, bool) {
	a := (machine.bx*machine.prizey - machine.prizex*machine.by) / (machine.bx*machine.ay - machine.ax*machine.by)
	b := (machine.prizex*machine.ay - machine.ax*machine.prizey) / (machine.bx*machine.ay - machine.ax*machine.by)

	x := machine.ax*a + machine.bx*b
	y := machine.ay*a + machine.by*b

	if x == machine.prizex && y == machine.prizey {
		return a*3 + b, true
	}

	return 0, false
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	machines := make([]machine, 0)

	lines := SplitByLines(input)
	for i := 0; i < len(lines); i += 4 {
		var machine machine
		var fields []string
		var field string

		fields = strings.Fields(lines[i])
		field = fields[2]
		machine.ax = StrToInt(field[2 : len(field)-1])
		machine.ay = StrToInt(fields[3][2:])

		fields = strings.Fields(lines[i+1])
		field = fields[2]
		machine.bx = StrToInt(field[2 : len(field)-1])
		machine.by = StrToInt(fields[3][2:])

		fields = strings.Fields(lines[i+2])
		field = fields[1]
		machine.prizex = 10_000_000_000_000 + StrToInt(field[2:len(field)-1])
		machine.prizey = 10_000_000_000_000 + StrToInt(fields[2][2:])

		machines = append(machines, machine)
	}

	sum := 0
	for _, machine := range machines {
		if cost, ok := getCost2(machine); ok {
			sum += cost
		}
	}

	fmt.Println(sum)
}
