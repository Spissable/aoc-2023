package day10

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(10)

	maze := parseMaze(input)

	result := solvePuzzle2(maze)

	if result != 355 {
		t.Error(result)
	}
}
