package day10

import (
	"aoc-2023/util"
	"strings"
)

type Position struct {
	x int
	y int
}

type Maze struct {
	start Position
	m     [][]rune
}

type Direction string

func solvePuzzle1(maze Maze) int {
	var visited = make(map[Position]bool)

	r := checkPosition(maze, visited, Position{maze.start.x + 1, maze.start.y}, 0)
	visited = make(map[Position]bool)
	l := checkPosition(maze, visited, Position{maze.start.x - 1, maze.start.y}, 0)
	visited = make(map[Position]bool)
	d := checkPosition(maze, visited, Position{maze.start.x, maze.start.y + 1}, 0)
	visited = make(map[Position]bool)
	u := checkPosition(maze, visited, Position{maze.start.x, maze.start.y - 1}, 0)
	return util.MaxInt(r, l, d, u) / 2
}

func checkPosition(maze Maze, visited map[Position]bool, p Position, step int) int {
	if visited[p] || p.x < 0 || p.y < 0 || p.x >= len(maze.m[0]) || p.y >= len(maze.m) || maze.m[p.y][p.x] == '.' {
		return step
	}
	step++
	visited[p] = true

	switch maze.m[p.y][p.x] {
	case '|':
		d := checkPosition(maze, visited, Position{p.x, p.y + 1}, step)
		u := checkPosition(maze, visited, Position{p.x, p.y - 1}, step)
		return util.MaxInt(d, u)
	case '-':
		r := checkPosition(maze, visited, Position{p.x + 1, p.y}, step)
		l := checkPosition(maze, visited, Position{p.x - 1, p.y}, step)
		return util.MaxInt(r, l)
	case 'F':
		r := checkPosition(maze, visited, Position{p.x + 1, p.y}, step)
		d := checkPosition(maze, visited, Position{p.x, p.y + 1}, step)
		return util.MaxInt(r, d)
	case '7':
		l := checkPosition(maze, visited, Position{p.x - 1, p.y}, step)
		d := checkPosition(maze, visited, Position{p.x, p.y + 1}, step)
		return util.MaxInt(l, d)
	case 'J':
		l := checkPosition(maze, visited, Position{p.x - 1, p.y}, step)
		u := checkPosition(maze, visited, Position{p.x, p.y - 1}, step)
		return util.MaxInt(l, u)
	case 'L':
		r := checkPosition(maze, visited, Position{p.x + 1, p.y}, step)
		u := checkPosition(maze, visited, Position{p.x, p.y - 1}, step)
		return util.MaxInt(r, u)
	case 'S':
		break
	}

	return step
}

func parseMaze(input string) Maze {
	lines := strings.Split(input, "\n")
	var maze Maze

	for y, line := range lines {
		var l []rune
		for x, char := range line {
			l = append(l, char)
			if char == 'S' {
				maze.start = Position{
					x: x,
					y: y,
				}
			}
		}

		maze.m = append(maze.m, l)
	}

	return maze
}
