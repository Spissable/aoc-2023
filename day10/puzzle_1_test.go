package day10

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(10)

	maze := parseMaze(input)

	result := solvePuzzle1(maze)

	if result != 7097 {
		t.Error(result)
	}
}
