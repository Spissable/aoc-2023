package day2

import (
	"strconv"
	"strings"
)

type Round struct {
	red   int
	green int
	blue  int
}

type Game struct {
	rounds []Round
}

func parseGame(input string) map[int]Game {
	lines := strings.Split(input, "\n")
	game := map[int]Game{}

	for _, line := range lines {
		mainParts := strings.Split(line, ":")

		gameId, _ := strconv.Atoi(strings.Split(mainParts[0], " ")[1])
		rounds := strings.Split(mainParts[1], ";")

		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			red := 0
			green := 0
			blue := 0

			for _, cube := range cubes {
				parts := strings.Split(strings.Trim(cube, " "), " ")
				num, _ := strconv.Atoi(parts[0])

				switch parts[1] {
				case "red":
					red = num

				case "green":
					green = num

				case "blue":
					blue = num
				}
			}

			round := Round{
				red,
				green,
				blue,
			}
			game[gameId] = Game{
				rounds: append(game[gameId].rounds, round),
			}
		}
	}

	return game
}

func solvePuzzle1(game map[int]Game, r, g, b int) int {
	sum := 0

	for k, v := range game {
		isPossible := true

		for _, round := range v.rounds {
			if round.red > r || round.green > g || round.blue > b {
				isPossible = false
			}
		}

		if isPossible {
			sum += k
		}
	}

	return sum
}
