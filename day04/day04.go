package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func SplitByLines(data string) []string {
	normalized := strings.ReplaceAll(data, "\r\n", "\n")
	return strings.Split(normalized, "\n")
}

//go:embed input.txt
var input string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := SplitByLines(input)
	count := 0

	height := len(lines)
	width := len(lines[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var east, south, southEast, northWest string
			if x+3 < width {
				east = lines[y][x : x+4]
			}
			if y+3 < height {
				south = string([]byte{
					lines[y+0][x],
					lines[y+1][x],
					lines[y+2][x],
					lines[y+3][x],
				})
			}
			if x+3 < width && y+3 < height {
				southEast = string([]byte{
					lines[y+0][x+0],
					lines[y+1][x+1],
					lines[y+2][x+2],
					lines[y+3][x+3],
				})
				northWest = string([]byte{
					lines[y+3][x+0],
					lines[y+2][x+1],
					lines[y+1][x+2],
					lines[y+0][x+3],
				})
			}

			for _, word := range []string{east, south, southEast, northWest} {
				if word == "XMAS" || word == "SAMX" {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func partTwo() {
	lines := SplitByLines(input)
	count := 0

	height := len(lines)
	width := len(lines[0])

	for y := 0; y <= height-3; y++ {
		for x := 0; x <= width-3; x++ {
			c := lines[y+1][x+1]
			if c != 'A' {
				continue
			}

			nw := lines[y][x]
			ne := lines[y][x+2]
			sw := lines[y+2][x]
			se := lines[y+2][x+2]

			if ((nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')) &&
				((sw == 'M' && ne == 'S') || (sw == 'S' && ne == 'M')) {
				count++
			}
		}
	}

	fmt.Println(count)
}

//
// OLD:
//
// const xmas = "XMAS"
// const samx = "SAMX"
// const length = len(xmas)

// func FindXmas(lines []string) int {
// 	count := 0
// 	for _, line := range lines {
// 		for i := 0; i <= len(line)-length; i++ {
// 			word := line[i : i+length]
// 			if word == xmas || word == samx {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

// func GetColumns(lines []string) []string {
// 	columns := make([]string, 0, len(lines[0]))
// 	height := len(lines)
// 	for x := range lines[0] {
// 		column := make([]byte, 0, height)
// 		for y := 0; y < height; y++ {
// 			column = append(column, lines[y][x])
// 		}
// 		columns = append(columns, string(column))
// 	}

// 	return columns
// }

// func GetDiagonals(lines []string) []string {
// 	width := len(lines[0])
// 	height := len(lines)
// 	diagonals := make([]string, 0, width+height)
// 	for x := 0; x < width; x++ {
// 		diagonal1 := make([]byte, 0, height)
// 		diagonal2 := make([]byte, 0, height)
// 		for i := 0; i < height; i++ {
// 			if x+i >= width {
// 				break
// 			}
// 			diagonal1 = append(diagonal1, lines[i][x+i])
// 			diagonal2 = append(diagonal2, lines[height-1-i][x+i])
// 		}
// 		diagonals = append(diagonals, string(diagonal1))
// 		diagonals = append(diagonals, string(diagonal2))
// 	}

// 	for y := 1; y < height; y++ {
// 		diagonal1 := make([]byte, 0, width)
// 		diagonal2 := make([]byte, 0, width)
// 		for i := 0; i < width; i++ {
// 			if y+i >= height {
// 				break
// 			}
// 			diagonal1 = append(diagonal1, lines[y+i][i])
// 			diagonal2 = append(diagonal2, lines[height-1-(y+i)][i])
// 		}
// 		diagonals = append(diagonals, string(diagonal1))
// 		diagonals = append(diagonals, string(diagonal2))
// 	}

// 	return diagonals
// }

// func partOne() {
// 	lines := SplitByLines(input)
// 	count := 0

// 	columns := GetColumns(lines)
// 	diagonals := GetDiagonals(lines)

// 	count += FindXmas(lines)
// 	count += FindXmas(columns)
// 	count += FindXmas(diagonals)

// 	fmt.Println(count)
// }
