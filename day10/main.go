package day10

import (
	"fmt"
	"math"
	"strings"
)

var input = "noop\nnoop\nnoop\naddx 6\nnoop\naddx 30\naddx -26\nnoop\naddx 5\nnoop\nnoop\nnoop\nnoop\naddx 5\naddx -5\naddx 6\naddx 5\naddx -1\naddx 5\nnoop\nnoop\naddx -14\naddx -18\naddx 39\naddx -39\naddx 25\naddx -22\naddx 2\naddx 5\naddx 2\naddx 3\naddx -2\naddx 2\nnoop\naddx 3\naddx 2\naddx 2\nnoop\naddx 3\nnoop\naddx 3\naddx 2\naddx 5\naddx 4\naddx -18\naddx 17\naddx -38\naddx 5\naddx 2\naddx -5\naddx 27\naddx -19\nnoop\naddx 3\naddx 4\nnoop\nnoop\naddx 5\naddx -1\nnoop\nnoop\naddx 4\naddx 5\naddx 2\naddx -4\naddx 5\nnoop\naddx -11\naddx 16\naddx -36\nnoop\naddx 5\nnoop\naddx 28\naddx -23\nnoop\nnoop\nnoop\naddx 21\naddx -18\nnoop\naddx 3\naddx 2\naddx 2\naddx 5\naddx 1\nnoop\nnoop\naddx 4\nnoop\nnoop\nnoop\nnoop\nnoop\naddx 8\naddx -40\nnoop\naddx 7\nnoop\naddx -2\naddx 5\naddx 2\naddx 25\naddx -31\naddx 9\naddx 5\naddx 2\naddx 2\naddx 3\naddx -2\nnoop\naddx 3\naddx 2\nnoop\naddx 7\naddx -2\naddx 5\naddx -40\naddx 20\naddx -12\nnoop\nnoop\nnoop\naddx -5\naddx 7\naddx 7\nnoop\naddx -1\naddx 1\naddx 5\naddx 3\naddx -2\naddx 2\nnoop\naddx 3\naddx 2\nnoop\nnoop\nnoop\nnoop\naddx 7\nnoop\nnoop\nnoop\nnoop"

var memory, sum int
var value, cycle = 1, 1

func cyc() {
	if math.Mod(float64(cycle)/20, 2) == 1 {
		sum += cycle * value
	}
	cycle++
}

func Part1() int {
	for _, op := range strings.Split(input, "\n") {
		if strings.HasPrefix(op, "add") {
			_, _ = fmt.Sscanf(op, "addx %d", &memory)
			cyc()
			cyc()
			value += memory
		} else {
			cyc()
		}
	}
	return sum
}

// Part 2 start
var screen = make(map[int][]string, 0)

func cycD() {
	row := (cycle / 40) + 1
	if cycle%40 == 0 {
		row = row - 1
	}
	pos := 39 - ((row * 40) - cycle)
	var pixel = "."
	if value == pos || value-1 == pos || value+1 == pos {
		pixel = "#"
	}
	screen[row] = append(screen[row], pixel)

	cyc()
}

func Part2() map[int][]string {
	for _, op := range strings.Split(input, "\n") {
		if strings.HasPrefix(op, "add") {
			_, _ = fmt.Sscanf(op, "addx %d", &memory)
			cycD()
			cycD()
			value += memory
		} else {
			cycD()
		}
	}
	// Returns {row_number, horizontal pixels}, output by:
	//  for i := 1; i < 7; i++ {
	//  	fmt.Println(screen[i])
	//  }
	return screen
}
