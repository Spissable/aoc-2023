package day11

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(11)

	universe := parseUniverse(input, 999999)

	result := solvePuzzle1(universe)

	if result != 550358864332 {
		t.Error(result)
	}
}
