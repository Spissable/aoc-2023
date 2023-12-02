package day2

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(2)

	game := parseGame(input)
	result := solvePuzzle2(game)

	if result != 78375 {
		t.Error(result)
	}
}
