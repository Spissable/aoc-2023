package day2

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(2)

	game := parseGame(input)
	result := solvePuzzle1(game, 12, 13, 14)

	if result != 2406 {
		t.Error(result)
	}
}
