package day7

import (
	"aoc-2023/util"
	"slices"
	"strings"
)

type Hand int

type Card int

const (
	One   Card = 1
	Two   Card = 2
	Three Card = 3
	Four  Card = 4
	Five  Card = 5
	Six   Card = 6
	Seven Card = 7
	Eight Card = 8
	Nine  Card = 9
	T     Card = 10
	J     Card = 11
	Q     Card = 12
	K     Card = 13
	A     Card = 14
)

const (
	HighCard  Hand = 1
	OnePair   Hand = 2
	TwoPair   Hand = 3
	ThreeKind Hand = 4
	FullHouse Hand = 5
	FourKind  Hand = 6
	FiveKind  Hand = 7
)

type Player struct {
	cards []Card
	bid   int
	hand  Hand
}

func solvePuzzle1(sortedPlayers []Player) int {
	sum := 0

	for i, player := range sortedPlayers {
		sum += (i + 1) * player.bid
	}

	return sum
}

func getHandValue(cards []Card) Hand {
	groups := map[Card]int{}

	for _, card := range cards {
		groups[card]++
	}

	groupCount := len(groups)

	if groupCount == 1 {
		return FiveKind
	}
	if groupCount == 2 {
		for _, count := range groups {
			if count == 1 || count == 4 {
				return FourKind
			}
			return FullHouse
		}
	}
	if groupCount == 3 {
		for _, count := range groups {
			if count == 3 {
				return ThreeKind
			}
			if count == 2 {
				return TwoPair
			}
		}
	}
	if groupCount == 4 {
		return OnePair
	}

	return HighCard
}

func insertPlayerSorted(player Player, players []Player) []Player {
	for i, p := range players {
		if player.hand < p.hand {
			// Hand has smaller value
			return slices.Insert(players, i, player)
		}
		if player.hand == p.hand {
			for i2, card := range player.cards {
				if card < p.cards[i2] {
					// Hand has same value but lesser card
					return slices.Insert(players, i, player)
				}
				if card > p.cards[i2] {
					// Hand has same value but higher card, continue outer loop to check with next player
					break
				}
				// Hand has same value and current card is same value, continue inner loop and check next card
			}
		}
	}

	return append(players, player)
}

func stringToCard(str string) Card {
	switch str {
	case "1":
		return One
	case "2":
		return Two
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "T":
		return T
	case "J":
		return J
	case "Q":
		return Q
	case "K":
		return K
	case "A":
		return A
	}

	panic("Unknown Card")
}

func parseGame(input string) []Player {
	lines := strings.Split(input, "\n")
	var players []Player

	for _, line := range lines {
		parts := strings.Split(line, " ")
		bid := util.StringToNum(parts[1])
		var cards []Card

		for _, card := range parts[0] {
			cards = append(cards, stringToCard(string(card)))
		}

		player := Player{
			cards: cards,
			bid:   bid,
			hand:  getHandValue(cards),
		}

		players = insertPlayerSorted(player, players)
	}

	return players
}
