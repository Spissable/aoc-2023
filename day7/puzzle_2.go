package day7

import (
	"aoc-2023/util"
	"slices"
	"strings"
)

type HandP2 int

type CardP2 int

const (
	JP2     CardP2 = 0
	OneP2   CardP2 = 1
	TwoP2   CardP2 = 2
	ThreeP2 CardP2 = 3
	FourP2  CardP2 = 4
	FiveP2  CardP2 = 5
	SixP2   CardP2 = 6
	SevenP2 CardP2 = 7
	EightP2 CardP2 = 8
	NineP2  CardP2 = 9
	TP2     CardP2 = 10
	QP2     CardP2 = 12
	KP2     CardP2 = 13
	AP2     CardP2 = 14
)

const (
	HighCardP2  HandP2 = 1
	OnePairP2   HandP2 = 2
	TwoPairP2   HandP2 = 3
	ThreeKindP2 HandP2 = 4
	FullHouseP2 HandP2 = 5
	FourKindP2  HandP2 = 6
	FiveKindP2  HandP2 = 7
)

type PlayerP2 struct {
	cards []CardP2
	bid   int
	hand  HandP2
}

func solvePuzzle2(sortedPlayerP2s []PlayerP2) int {
	sum := 0

	for i, player := range sortedPlayerP2s {
		sum += (i + 1) * player.bid
	}

	return sum
}

func getHandP2Value(cards []CardP2) HandP2 {
	groups := map[CardP2]int{}

	biggestGroupLength := 0
	jokerCount := 0
	var biggestGroupKey CardP2

	for _, card := range cards {
		if card != JP2 {
			groups[card]++

			groupLength := groups[card]
			if groupLength > biggestGroupLength {
				biggestGroupLength = groupLength
				biggestGroupKey = card
			}
		} else {
			jokerCount++
		}
	}

	// Add the Joker to the biggest group for maximum effect
	groups[biggestGroupKey] += jokerCount

	groupCount := len(groups)

	if groupCount == 1 {
		return FiveKindP2
	}
	if groupCount == 2 {
		for _, count := range groups {
			if count == 1 || count == 4 {
				return FourKindP2
			}
			return FullHouseP2
		}
	}
	if groupCount == 3 {
		for _, count := range groups {
			if count == 3 {
				return ThreeKindP2
			}
			if count == 2 {
				return TwoPairP2
			}
		}
	}
	if groupCount == 4 {
		return OnePairP2
	}

	return HighCardP2
}

func insertPlayerP2SortedP2(player PlayerP2, players []PlayerP2) []PlayerP2 {
	for i, p := range players {
		if player.hand < p.hand {
			// HandP2 has smaller value
			return slices.Insert(players, i, player)
		}
		if player.hand == p.hand {
			for i2, card := range player.cards {
				if card < p.cards[i2] {
					// HandP2 has same value but lesser card
					return slices.Insert(players, i, player)
				}
				if card > p.cards[i2] {
					// HandP2 has same value but higher card, continue outer loop to check with next player
					break
				}
				// HandP2 has same value and current card is same value, continue inner loop and check next card
			}
		}
	}

	return append(players, player)
}

func stringToCardP2(str string) CardP2 {
	switch str {
	case "1":
		return OneP2
	case "2":
		return TwoP2
	case "3":
		return ThreeP2
	case "4":
		return FourP2
	case "5":
		return FiveP2
	case "6":
		return SixP2
	case "7":
		return SevenP2
	case "8":
		return EightP2
	case "9":
		return NineP2
	case "T":
		return TP2
	case "J":
		return JP2
	case "Q":
		return QP2
	case "K":
		return KP2
	case "A":
		return AP2
	}

	panic("Unknown CardP2")
}

func parseGameP2(input string) []PlayerP2 {
	lines := strings.Split(input, "\n")
	var players []PlayerP2

	for _, line := range lines {
		parts := strings.Split(line, " ")
		bid := util.StringToNum(parts[1])
		var cards []CardP2

		for _, card := range parts[0] {
			cards = append(cards, stringToCardP2(string(card)))
		}

		player := PlayerP2{
			cards: cards,
			bid:   bid,
			hand:  getHandP2Value(cards),
		}

		players = insertPlayerP2SortedP2(player, players)
	}

	return players
}
