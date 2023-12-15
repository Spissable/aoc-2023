package day12

import (
	"aoc-2023/util"
	"fmt"
	"regexp"
	"strings"
)

type SpringsLine struct {
	cgroups        []int
	springs        string
	unknownIndices []int
}

type SpringField []SpringsLine

func solvePuzzle1(springField SpringField) int {
	sum := 0
	for _, line := range springField {
		sum += getLinePossibilities(line)
	}

	return sum
}

func getLinePossibilities(line SpringsLine) int {
	possibilities := 0

	combinations := getAllCombinations(make([]bool, len(line.unknownIndices)))
	regex := regexp.MustCompile(createRegex(line.cgroups))

	for _, combination := range combinations {
		testLine := line.springs
		for i, unknownIndex := range line.unknownIndices {
			if combination[i] {
				testLine = testLine[:unknownIndex] + "." + testLine[unknownIndex+1:]
			} else {
				testLine = testLine[:unknownIndex] + "#" + testLine[unknownIndex+1:]
			}
		}

		if regex.MatchString(testLine) {
			possibilities++
		}
	}

	return possibilities
}

func generateCombinations(arr []bool, index int, combination []bool, result *[][]bool) {
	if index == len(arr) {
		temp := make([]bool, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
		return
	}

	// Include the current element in the combination
	combination[index] = true
	generateCombinations(arr, index+1, combination, result)

	// Exclude the current element from the combination
	combination[index] = false
	generateCombinations(arr, index+1, combination, result)
}

func getAllCombinations(arr []bool) [][]bool {
	result := make([][]bool, 0)
	combination := make([]bool, len(arr))
	generateCombinations(arr, 0, combination, &result)
	return result
}

func parseSpringfield(input string) SpringField {
	lines := strings.Split(input, "\n")
	var springField SpringField

	for _, line := range lines {
		parts := strings.Split(line, " ")

		var unknownIndices []int
		for i, spring := range parts[0] {
			if spring == '?' {
				unknownIndices = append(unknownIndices, i)
			}
		}

		cgroups := util.StringsToNums(strings.Split(parts[1], ","))
		springField = append(springField, SpringsLine{
			cgroups:        cgroups,
			springs:        parts[0],
			unknownIndices: unknownIndices,
		})
	}

	return springField
}

func createRegex(cgroups []int) string {
	regex := `^(\.)*`
	for i, cgroup := range cgroups {
		regex += fmt.Sprintf("#{%d}", cgroup)

		regex += `(\.)`

		if i == len(cgroups)-1 {
			regex += "*$"
		} else {
			regex += "+"
		}
	}

	return regex
}
