package day9

func solvePuzzle2(oasis Oasis) int {
	sum := 0

	for _, line := range oasis {
		sum += makeLeftPrediction(line)
	}

	return sum
}

func makeLeftPrediction(line []int) int {
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

	previousNum := lines[len(lines)-1][0]
	num := 0

	for i := len(lines) - 2; i >= 0; i-- {
		num = lines[i][0] - previousNum
		previousNum = num
	}

	return num
}
