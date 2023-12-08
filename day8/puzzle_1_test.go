package day8

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(8)

	leMap := parseMap(input)

	result := solvePuzzle1(leMap)

	if result != 17287 {
		t.Error(result)
	}
}
