package day9

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(9)

	oasis := parseOasis(input)

	result := solvePuzzle2(oasis)

	if result != 1131 {
		t.Error(result)
	}
}
