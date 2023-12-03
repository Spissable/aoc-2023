package day3

import (
	"strconv"
	"strings"
)

const period = 46
const numMin = 48
const numMax = 57

func solvePuzzle1(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	numStr := ""
	from := -1
	to := -1

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
				if hasAdjacentSymbol(lines, lineI, from, to) {
					num, _ := strconv.Atoi(numStr)
					sum += num
				}

				numStr = ""
				from = -1
				to = -1
			}
		}
	}

	return sum
}

func isSymbol(char rune) bool {
	return char > numMax || (char < numMin && char != period)
}

func hasAdjacentSymbol(lines []string, lineNum, from, to int) bool {
	numLen := to - from + 1

	if lineNum > 0 {
		// Check above
		for i := 0; i < numLen; i++ {
			if isSymbol(rune(lines[lineNum-1][from+i])) {
				return true
			}
		}

		if from > 0 {
			// Check top left
			if isSymbol(rune(lines[lineNum-1][from-1])) {
				return true
			}
		}

		if to < len(lines[lineNum])-1 {
			// Check top right
			if isSymbol(rune(lines[lineNum-1][to+1])) {
				return true
			}
		}
	}

	if lineNum < len(lines)-1 {
		// Check below
		for i := 0; i < numLen; i++ {
			if isSymbol(rune(lines[lineNum+1][from+i])) {
				return true
			}
		}

		if from > 0 {
			// Check bottom left
			if isSymbol(rune(lines[lineNum+1][from-1])) {
				return true
			}
		}

		if to < len(lines[lineNum])-1 {
			// Check bottom right
			if isSymbol(rune(lines[lineNum+1][to+1])) {
				return true
			}
		}
	}

	if from > 0 {
		// Check left
		if isSymbol(rune(lines[lineNum][from-1])) {
			return true
		}
	}

	if to < len(lines[lineNum])-1 {
		// Check right
		if isSymbol(rune(lines[lineNum][to+1])) {
			return true
		}
	}

	return false
}
