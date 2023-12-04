package day4

func solvePuzzle2(games []Game) int {
	sum := 0
	cards := map[int]int{}

	for i, game := range games {
		cards[i]++
		winners := 0
		for _, own := range game.own {
			if isWinner(game.winners, own) {
				winners++
				copyIndex := i + winners

				if copyIndex < len(games) {
					cards[copyIndex] += cards[i]
				}
			}
		}

		sum += cards[i]
	}

	return sum
}
