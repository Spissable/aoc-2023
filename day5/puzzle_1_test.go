package day5

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(5)

	almanac := parseAlmanac(input)

	result := solvePuzzle1(almanac)

	if result != 165788812 {
		t.Error(result)
	}
}
