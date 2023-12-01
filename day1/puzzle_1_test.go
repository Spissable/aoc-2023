package day1

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(1)

	result := solvePuzzle1(input)

	if result != 54390 {
		t.Error(result)
	}
}
