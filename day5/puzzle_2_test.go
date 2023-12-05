package day5

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(5)

	almanac := parseAlmanac(input)

	result := solvePuzzle2(almanac)

	if result != 1928058 {
		t.Error(result)
	}
}
