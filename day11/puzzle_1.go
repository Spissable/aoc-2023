package day11

import (
	"strings"
)

type Universe [][]bool
type Coords struct {
	x int
	y int
}
type Galaxies []Coords

func solvePuzzle1(galaxies Galaxies) int {
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for i2 := i + 1; i2 < len(galaxies); i2++ {
			b := galaxies[i2]
			sum += calcPairingDistance(a, b)
		}
	}

	return sum
}

func calcPairingDistance(a, b Coords) int {
	steps := 0
	if a.x > b.x {
		steps += a.x - b.x
	} else {
		steps += b.x - a.x
	}

	if a.y > b.y {
		steps += a.y - b.y
	} else {
		steps += b.y - a.y
	}

	return steps
}

func parseUniverse(input string, expansionSize int) Galaxies {
	lines := strings.Split(input, "\n")

	var universe Universe
	var galaxies Galaxies

	for _, line := range lines {
		rowGalaxy := false
		var row []bool

		for x, char := range line {
			isGalaxy := false
			if char == '#' {
				rowGalaxy = true
				isGalaxy = true
				galaxies = append(galaxies, Coords{x: len(row), y: len(universe)})
			}

			row = append(row, isGalaxy)

			if !columnContainsGalaxy(lines, x) {
				for i := 0; i < expansionSize; i++ {
					row = append(row, false)
				}
			}
		}

		universe = append(universe, row)

		if !rowGalaxy {
			for i := 0; i < expansionSize; i++ {
				universe = append(universe, row)
			}
		}
	}

	return galaxies
}

func columnContainsGalaxy(lines []string, column int) bool {
	for _, line := range lines {
		if line[column] == '#' {
			return true
		}
	}

	return false
}
