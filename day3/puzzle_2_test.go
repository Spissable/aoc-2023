package day3

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(3)
	result := solvePuzzle2(input)

	if result != 81709807 {
		t.Error(result)
	}
}
