package day11

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(11)

	universe := parseUniverse(input)

	result := solvePuzzle1(universe)

	if result != 9965032 {
		t.Error(result)
	}
}
