package day12

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(12)

	springField := parseSpringfield(input)

	result := solvePuzzle1(springField)

	if result != 7670 {
		t.Error(result)
	}
}
