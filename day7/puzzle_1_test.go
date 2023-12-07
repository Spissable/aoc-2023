package day7

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(7)

	sortedPlayers := parseGame(input)

	result := solvePuzzle1(sortedPlayers)

	if result != 256448566 {
		t.Error(result)
	}
}
