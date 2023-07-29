//=======================================================================
// Advent of Code 2015 Day 01 -- Not Quite Lisp
//=======================================================================

package main

import (
	"fmt"
	"os"
)

func part1() int {
	var input, err = os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var floor = 0
	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func part2() int {
	var input, err = os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var floor = 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("Part 1:", part1()) // 74
	fmt.Println("Part 2:", part2()) // 1795
}
