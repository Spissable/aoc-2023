package day4

import (
	"strconv"
	"strings"
)

type Game struct {
	winners []int
	own     []int
}

func solvePuzzle1(games []Game) int {
	sum := 0

	for _, game := range games {
		currentGameSum := 0

		for _, own := range game.own {
			if isWinner(game.winners, own) {
				if currentGameSum == 0 {
					currentGameSum = 1
				} else {
					currentGameSum *= 2
				}
			}
		}

		sum += currentGameSum
	}

	return sum
}

func isWinner(winners []int, num int) bool {
	for _, winner := range winners {
		if num == winner {
			return true
		}
	}

	return false
}

func parseCards(input string) []Game {
	var game []Game
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		piles := strings.Split(strings.Split(line, ":")[1], "|")
		var winners []int
		var own []int

		winnerStrs := strings.Split(strings.Trim(piles[0], " "), " ")
		for _, char := range winnerStrs {
			num, _ := strconv.Atoi(char)
			if num != 0 {
				winners = append(winners, num)
			}
		}

		ownStrs := strings.Split(strings.Trim(piles[1], " "), " ")
		for _, char := range ownStrs {
			num, _ := strconv.Atoi(char)
			if num != 0 {
				own = append(own, num)
			}
		}

		game = append(game, Game{
			winners: winners,
			own:     own,
		})
	}

	return game
}
