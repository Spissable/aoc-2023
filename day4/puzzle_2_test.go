package day4

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(4)

	game := parseCards(input)
	result := solvePuzzle2(game)

	if result != 8467762 {
		t.Error(result)
	}
}
