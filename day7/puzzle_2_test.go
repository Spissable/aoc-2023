package day7

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(7)

	sortedPlayers := parseGameP2(input)

	result := solvePuzzle2(sortedPlayers)

	if result != 254412181 {
		t.Error(result)
	}
}
