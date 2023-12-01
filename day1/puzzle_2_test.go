package day1

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(1)

	result := solvePuzzle2(input)

	if result != 54277 {
		t.Error(result)
	}
}
