package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
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

func mod(a, b int) int {
	return (a%b + b) % b
}

//go:embed input.txt
var input string

// const width = 11
// const height = 7

const width = 101
const height = 103

type robot struct {
	x, y   int
	vx, vy int
}

func main() {
	fmt.Println("Day 14")
	partOne()
	partTwo()
}

func printState(w io.Writer, iteration int, robots []robot) {
	fmt.Fprintln(w, iteration, "seconds:")
	for y := 0; y < height; y++ {
	LINE:
		for x := 0; x < width; x++ {
			for _, robot := range robots {
				if x == robot.x && y == robot.y {
					fmt.Fprint(w, "█")
					continue LINE
				}
			}
			fmt.Fprint(w, " ")
		}

		fmt.Fprint(w, "\n")
	}

	fmt.Fprint(w, "\n\n\n\n\n\n\n\n\n\n\n\n")
}

func countQuadrants(robots []robot) (int, int, int, int) {
	var q1, q2, q3, q4 int
	for _, robot := range robots {
		if robot.x < width/2 {
			if robot.y < height/2 {
				q1++
			} else if robot.y > height/2 {
				q2++
			}
		} else if robot.x > width/2 {
			if robot.y < height/2 {
				q3++
			} else if robot.y > height/2 {
				q4++
			}
		}
	}

	return q1, q2, q3, q4
}

func partOne() {
	defer PrintTimeSince(time.Now())

	robots := make([]robot, 0)
	for _, line := range SplitByLines(input) {
		fields := strings.Fields(line)
		p := strings.Split(fields[0][2:], ",")
		v := strings.Split(fields[1][2:], ",")
		robot := robot{
			x:  StrToInt(p[0]),
			y:  StrToInt(p[1]),
			vx: StrToInt(v[0]),
			vy: StrToInt(v[1]),
		}
		robots = append(robots, robot)
	}

	for i := 0; i < 100; i++ {
		for r := range robots {
			robots[r].x = mod(robots[r].x+robots[r].vx, width)
			robots[r].y = mod(robots[r].y+robots[r].vy, height)
		}

	}

	q1, q2, q3, q4 := countQuadrants(robots)
	safetyFactor := q1 * q2 * q3 * q4
	fmt.Println(safetyFactor)
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	robots := make([]robot, 0)
	for _, line := range SplitByLines(input) {
		fields := strings.Fields(line)
		p := strings.Split(fields[0][2:], ",")
		v := strings.Split(fields[1][2:], ",")
		robot := robot{
			x:  StrToInt(p[0]),
			y:  StrToInt(p[1]),
			vx: StrToInt(v[0]),
			vy: StrToInt(v[1]),
		}
		robots = append(robots, robot)
	}

	f, _ := os.Create("./day14-output.txt")
	w := bufio.NewWriter(f)

	for i := 0; i < 1000; i++ {
		printState(w, i, robots)

		for r := range robots {
			robots[r].x = mod(robots[r].x+robots[r].vx, width)
			robots[r].y = mod(robots[r].y+robots[r].vy, height)
		}

	}

	w.Flush()

	fmt.Println("See day14-output.txt (find the christmas tree!)")
}
