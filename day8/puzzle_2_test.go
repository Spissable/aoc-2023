package day8

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(8)

	leMap := parseMap(input)

	result := solvePuzzle2(leMap)

	if result != 18625484023687 {
		t.Error(result)
	}
}
