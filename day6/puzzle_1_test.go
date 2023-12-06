package day6

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(6)

	races := parseInput(input)

	result := solvePuzzle1(races)

	if result != 505494 {
		t.Error(result)
	}
}
