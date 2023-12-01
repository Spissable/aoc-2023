package day1

import (
	"strconv"
	"strings"
)

func solvePuzzle1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		calibrationValue := ""

		for _, char := range line {
			if char >= 48 && char <= 57 {
				calibrationValue += string(char)
				break
			}
		}

		for i, _ := range line {
			char := line[len(line)-1-i]
			if char >= 48 && char <= 57 {
				calibrationValue += string(char)
				break
			}
		}

		lineVal, _ := strconv.Atoi(calibrationValue)
		sum += lineVal
	}

	return sum
}
