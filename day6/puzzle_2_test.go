package day6

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(6)

	races := parseSingleRace(input)

	result := solvePuzzle2(races)

	if result != 23632299 {
		t.Error(result)
	}
}
