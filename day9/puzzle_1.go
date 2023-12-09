package day9

import (
	"aoc-2023/util"
	"regexp"
	"strings"
)

type Oasis [][]int

func solvePuzzle1(oasis Oasis) int {
	sum := 0

	for _, line := range oasis {
		sum += makePrediction(line)
	}

	return sum
}

func makePrediction(line []int) int {
	lines := [][]int{line}
	allNulls := false
	for lineI := 0; allNulls == false; lineI++ {
		allNulls = true
		currentLine := lines[lineI]
		var newLine []int

		for numI := 0; numI < len(currentLine)-1; numI++ {
			num := currentLine[numI+1] - currentLine[numI]
			newLine = append(newLine, num)
			if num != 0 {
				allNulls = false
			}
		}

		lines = append(lines, newLine)
	}

	previousNum := lines[len(lines)-1][len(lines[len(lines)-1])-1]
	num := 0

	for i := len(lines) - 2; i >= 0; i-- {
		num = previousNum + lines[i][len(lines[i])-1]
		previousNum = num
	}

	return num
}

func parseOasis(input string) Oasis {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`-?\d+`)
	oasis := make(Oasis, len(lines))

	for i, line := range lines {
		nums := util.StringsToNums(re.FindAllString(line, -1))
		oasis[i] = nums
	}

	return oasis
}
