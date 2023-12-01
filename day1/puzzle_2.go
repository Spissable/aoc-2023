package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solvePuzzle2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	stringNumbers := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	for _, line := range lines {
		firstDigit := 0
		firstDigitIndex := math.MaxInt
		lastDigit := 0
		lastDigitIndex := -1

		for i, char := range line {
			if char >= 48 && char <= 57 {
				val, _ := strconv.Atoi(string(char))
				if i < firstDigitIndex {
					firstDigit = val
					firstDigitIndex = i
				}
				if i > lastDigitIndex {
					lastDigit = val
					lastDigitIndex = i
				}
			}
		}

		for i, num := range stringNumbers {
			strFirstIndex := strings.Index(line, num)
			strLastIndexIndex := strings.LastIndex(line, num)

			if strFirstIndex != -1 && strFirstIndex < firstDigitIndex {
				firstDigit = i + 1
				firstDigitIndex = strFirstIndex
			}

			if strLastIndexIndex != -1 && strLastIndexIndex > lastDigitIndex {
				lastDigit = i + 1
				lastDigitIndex = strLastIndexIndex
			}
		}

		lineVal, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		sum += lineVal
	}

	return sum
}
