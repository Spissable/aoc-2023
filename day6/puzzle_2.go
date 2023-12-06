package day6

import (
	"aoc-2023/util"
	"regexp"
	"strings"
)

func solvePuzzle2(race Race) int {
	possibilitiesToWin := 0

	minHold := 0
	maxHold := 0
	for i := 1; i <= race.time/2; i++ {
		if minHold > 0 && maxHold > 0 {
			winsForThisRace := maxHold - minHold + 1
			if possibilitiesToWin == 0 {
				possibilitiesToWin = winsForThisRace
			} else {
				possibilitiesToWin *= winsForThisRace
			}

			break
		}

		if minHold == 0 {
			speed := i
			remainingTime := race.time - i
			if speed*remainingTime > race.distance {
				minHold = speed
			}
		}
		if maxHold == 0 {
			speed := race.time - i
			remainingTime := i
			if speed*remainingTime > race.distance {
				maxHold = speed
			}
		}
	}

	return possibilitiesToWin
}

func parseSingleRace(input string) Race {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d+`)

	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	return Race{
		time:     util.StringToNum(strings.Join(times, "")),
		distance: util.StringToNum(strings.Join(distances, "")),
	}
}
