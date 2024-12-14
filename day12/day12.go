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

type plot struct {
	left, right, up, down *plot
	plant                 rune
	x, y                  int
	visited               bool
}

func main() {
	fmt.Println("Day 12")
	partOne()
	partTwo()
}

func getPrice(plot *plot) int {
	if plot.visited {
		return 0
	}

	area, perimeter := getAreaAndPerimeter(plot)
	return area * perimeter
}

func getAreaAndPerimeter(p *plot) (int, int) {
	if p.visited {
		return 0, 0
	}

	p.visited = true

	area := 1

	perimeter := 0
	for _, neighbour := range []*plot{p.left, p.right, p.up, p.down} {
		if neighbour == nil || neighbour.plant != p.plant {
			perimeter++
		}

		if neighbour != nil && !neighbour.visited && neighbour.plant == p.plant {
			// fmt.Println("call", neighbour.x, neighbour.y, neighbour.visited)
			na, np := getAreaAndPerimeter(neighbour)
			area += na
			perimeter += np
		}
	}

	return area, perimeter
}

func partOne() {
	defer PrintTimeSince(time.Now())

	garden := make([]plot, 0)
	lines := SplitByLines(input)
	width := len(lines[0])
	height := len(lines)
	for y, line := range lines {
		for x, plant := range line {
			garden = append(garden, plot{plant: plant, x: x, y: y})
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			plot := &garden[y*width+x]
			if x > 0 {
				plot.left = &garden[y*width+x-1]
			}
			if x < width-1 {
				plot.right = &garden[y*width+x+1]
			}
			if y > 0 {
				plot.up = &garden[y*width+x-width]
			}
			if y < height-1 {
				plot.down = &garden[y*width+x+width]
			}
		}
	}

	sum := 0
	for i := range garden {
		sum += getPrice(&garden[i])
	}

	// sum = getPrice(&garden[0])

	fmt.Println(sum)
}

func getPrice2(plot *plot) int {
	if plot.visited {
		return 0
	}

	area, perimeter := getAreaAndSides(plot)
	// fmt.Println(string(plot.plant), ":", plot.x, plot.y, area, "*", perimeter, "=", area*perimeter)
	return area * perimeter
}

func getAreaAndSides(p *plot) (int, int) {
	if p.visited {
		return 0, 0
	}

	p.visited = true

	area := 1

	isEdge := func(other *plot) bool {
		return other == nil || other.plant != p.plant
	}

	up := isEdge(p.up)
	down := isEdge(p.down)
	right := isEdge(p.right)
	left := isEdge(p.left)

	corners := 0

	if up && right {
		corners++
	}
	if right && down {
		corners++
	}
	if down && left {
		corners++
	}
	if left && up {
		corners++
	}

	if !up && !right && p.up.right.plant != p.plant {
		corners++
	}
	if !right && !down && p.right.down.plant != p.plant {
		corners++
	}
	if !down && !left && p.down.left.plant != p.plant {
		corners++
	}
	if !left && !up && p.left.up.plant != p.plant {
		corners++
	}

	// if corners != 0 && p.plant == 'C' {
	// 	fmt.Println(string(p.plant), "corners!", p.x, p.y, corners)
	// }

	for _, neighbour := range []*plot{p.left, p.right, p.up, p.down} {
		if !isEdge(neighbour) && !neighbour.visited {
			na, nc := getAreaAndSides(neighbour)
			area += na
			corners += nc
		}
	}

	return area, corners
}

func partTwo() {
	defer PrintTimeSince(time.Now())

	garden := make([]plot, 0)
	lines := SplitByLines(input)
	width := len(lines[0])
	height := len(lines)
	for y, line := range lines {
		for x, plant := range line {
			garden = append(garden, plot{plant: plant, x: x, y: y})
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			plot := &garden[y*width+x]
			if x > 0 {
				plot.left = &garden[y*width+x-1]
			}
			if x < width-1 {
				plot.right = &garden[y*width+x+1]
			}
			if y > 0 {
				plot.up = &garden[y*width+x-width]
			}
			if y < height-1 {
				plot.down = &garden[y*width+x+width]
			}
		}
	}

	sum := 0
	for i := range garden {
		sum += getPrice2(&garden[i])
	}

	// sum = getPrice(&garden[0])

	fmt.Println(sum)
}
