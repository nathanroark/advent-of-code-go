//==========================================================================
// Advent of Code 2015 *** Day 3: Perfectly Spherical Houses in a Vacuum ***
//==========================================================================

package main

import (
	"fmt"
	"os"
)

func main() {
	// Part 1
	fmt.Println("Part 1: ", part1())

	// Part 2
	fmt.Println("Part 2: ", part2())
}

func part1() int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	houses := make(map[string]int)
	x, y := 0, 0
	houses["0,0"] = 1
	for _, v := range input {
		switch v {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		houses[fmt.Sprintf("%d,%d", x, y)]++
	}
	return len(houses)
}

func part2() int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	houses := make(map[string]int)
	sx, sy, rx, ry := 0, 0, 0, 0
	houses["0,0"] = 2
	for i, v := range input {
		switch v {
		case '^':
			if i%2 == 0 {
				sy++
			} else {
				ry++
			}
		case 'v':
			if i%2 == 0 {
				sy--
			} else {
				ry--
			}
		case '>':
			if i%2 == 0 {
				sx++
			} else {
				rx++
			}
		case '<':
			if i%2 == 0 {
				sx--
			} else {
				rx--
			}
		}
		houses[fmt.Sprintf("%d,%d", sx, sy)]++
		houses[fmt.Sprintf("%d,%d", rx, ry)]++
	}
	return len(houses)
}
