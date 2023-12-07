package puzzle2

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(7)

	sortedPlayers := parseGame(input)

	result := solvePuzzle2(sortedPlayers)

	if result != 505494 {
		t.Error(result)
	}
}
