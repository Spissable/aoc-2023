package day2

func solvePuzzle2(game map[int]Game) int {
	power := 0

	for _, v := range game {
		minR := 0
		minG := 0
		minB := 0

		for _, round := range v.rounds {
			r := round.red
			g := round.green
			b := round.blue

			if r > minR {
				minR = r
			}
			if g > minG {
				minG = g
			}
			if b > minB {
				minB = b
			}
		}

		power += minR * minG * minB
	}

	return power
}
