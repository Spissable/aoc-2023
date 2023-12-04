package day4

import (
	"aoc-2023/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(4)

	game := parseCards(input)
	result := solvePuzzle1(game)

	if result != 26346 {
		t.Error(result)
	}
}
