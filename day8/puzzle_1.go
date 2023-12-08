package day8

import (
	"regexp"
	"strings"
)

type Direction string

const (
	L Direction = "L"
	R Direction = "R"
)

type Directions struct {
	l string
	r string
}

type DMap map[string]Directions

type Map struct {
	instr []Direction
	dMap  DMap
}

func solvePuzzle1(leMap Map) int {
	node := "AAA"
	counter := 0
	for {
		for _, direction := range leMap.instr {
			counter++
			if direction == L {
				node = leMap.dMap[node].l
			} else {
				node = leMap.dMap[node].r
			}

			if node == "ZZZ" {
				return counter
			}
		}
	}
}

func parseDirection(char rune) Direction {
	if string(char) == "R" {
		return R
	}
	if string(char) == "L" {
		return L
	}

	panic("Invalid direction")
}

func addDMapLine(line string, leMap Map) Map {
	re := regexp.MustCompile(`\b\w{3}\b`)

	blocks := re.FindAllString(line, -1)
	node := blocks[0]
	l := blocks[1]
	r := blocks[2]

	leMap.dMap[node] = Directions{
		l: l,
		r: r,
	}

	return leMap
}

func parseMap(input string) Map {
	lines := strings.Split(input, "\n")

	leMap := Map{
		dMap: DMap{},
	}

	for i, line := range lines {
		if i == 0 {
			for _, char := range line {
				leMap.instr = append(leMap.instr, parseDirection(char))
			}
		} else if len(line) > 0 {
			leMap = addDMapLine(line, leMap)
		}
	}

	return leMap
}
