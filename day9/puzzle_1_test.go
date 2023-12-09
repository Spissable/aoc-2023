package day9

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(9)

	oasis := parseOasis(input)

	result := solvePuzzle1(oasis)

	if result != 1955513104 {
		t.Error(result)
	}
}
