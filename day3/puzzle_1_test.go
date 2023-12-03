package day3

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(3)
	result := solvePuzzle1(input)

	if result != 538046 {
		t.Error(result)
	}
}
