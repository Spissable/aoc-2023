package day3

import (
	"fmt"
	"strconv"
	"strings"
)

const gear = 42

type Coords struct {
	line int
	pos  int
}

type Gears map[string][]int

func solvePuzzle2(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	numStr := ""
	from := -1
	to := -1
	gears := Gears{}

	for lineI, line := range lines {

		for charI := 0; charI < len(line); charI++ {
			r := line[charI]
			if r <= numMax && r >= numMin {
				numStr += string(r)
				if from == -1 {
					from = charI
				}
				to = charI

			} else if numStr != "" {
				coords := getAdjacentGears(lines, lineI, from, to)
				for _, coord := range coords {
					num, _ := strconv.Atoi(numStr)
					gearCoord := fmt.Sprintf("%d,%d", coord.line, coord.pos)
					gears[gearCoord] = append(gears[gearCoord], num)
				}

				numStr = ""
				from = -1
				to = -1
			}
		}
	}

	for _, parts := range gears {
		if len(parts) == 2 {
			sum += parts[0] * parts[1]
		}
	}

	return sum
}

func getAdjacentGears(lines []string, lineNum, from, to int) []Coords {
	numLen := to - from + 1
	var coords []Coords

	if lineNum > 0 {
		// Check above
		line := lineNum - 1
		for i := 0; i < numLen; i++ {
			pos := from + i
			if rune(lines[line][pos]) == gear {
				coords = append(coords, Coords{
					line: line,
					pos:  pos,
				})
			}
		}

		if from > 0 {
			// Check top left
			pos := from - 1
			if rune(lines[line][pos]) == gear {
				coords = append(coords, Coords{
					line: line,
					pos:  pos,
				})
			}
		}

		if to < len(lines[line])-1 {
			// Check top right
			pos := to + 1
			if rune(lines[line][pos]) == gear {
				coords = append(coords, Coords{
					line: line,
					pos:  pos,
				})
			}
		}
	}

	if lineNum < len(lines)-1 {
		// Check below
		line := lineNum + 1
		for i := 0; i < numLen; i++ {
			pos := from + i
			if rune(lines[line][pos]) == gear {
				coords = append(coords, Coords{
					line: line,
					pos:  pos,
				})
			}
		}

		if from > 0 {
			// Check bottom left
			pos := from - 1
			if rune(lines[line][pos]) == gear {
				coords = append(coords, Coords{
					line: line,
					pos:  pos,
				})
			}
		}

		if to < len(lines[lineNum])-1 {
			// Check bottom right
			pos := to + 1
			if rune(lines[line][pos]) == gear {
				coords = append(coords, Coords{
					line: line,
					pos:  pos,
				})
			}
		}
	}

	if from > 0 {
		// Check left
		pos := from - 1
		if rune(lines[lineNum][pos]) == gear {
			coords = append(coords, Coords{
				line: lineNum,
				pos:  pos,
			})
		}
	}

	if to < len(lines[lineNum])-1 {
		// Check right
		pos := to + 1
		if rune(lines[lineNum][pos]) == gear {
			coords = append(coords, Coords{
				line: lineNum,
				pos:  pos,
			})
		}
	}

	return coords
}
