//=======================================================================
// Advent of Code 2015 Day 02 -- I Was Told There Would Be No Math
// http://adventofcode.com/2015/day/2
//=======================================================================

package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func part1() int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for _, line := range bytes.Split(input, []byte("\n")) {
		var x, y, z int
		fmt.Sscanf(string(line), "%dx%dx%d", &x, &y, &z)
		sides := []int{x, y, z}
		sort.Ints(sides)
		l, w, h := sides[0], sides[1], sides[2]
		total += 2*l*w + 2*w*h + 2*h*l + l*w
	}
	return total
}

func part2() int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for _, line := range bytes.Split(input, []byte("\n")) {
		var l, w, h int
		fmt.Sscanf(string(line), "%dx%dx%d", &l, &w, &h)
		sides := []int{l, w, h}
		sort.Ints(sides)
		total += 2*sides[0] + 2*sides[1] + l*w*h
	}
	return total
}

func main() {
	fmt.Println("Part 1:", part1()) // 1588178
	fmt.Println("Part 2:", part2()) // 3783758
}
