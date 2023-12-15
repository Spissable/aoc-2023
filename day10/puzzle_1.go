package day10

import (
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

type Path []Position

type Direction string

func solvePuzzle1(maze Maze) int {
	loop := getLoop(maze)
	return len(loop) / 2
}

func getLoop(maze Maze) Path {
	var visited = make(map[Position]bool)

	r := checkPosition(maze, visited, Position{maze.start.x + 1, maze.start.y}, Path{})
	visited = make(map[Position]bool)
	l := checkPosition(maze, visited, Position{maze.start.x - 1, maze.start.y}, Path{})
	visited = make(map[Position]bool)
	d := checkPosition(maze, visited, Position{maze.start.x, maze.start.y + 1}, Path{})
	visited = make(map[Position]bool)
	u := checkPosition(maze, visited, Position{maze.start.x, maze.start.y - 1}, Path{})

	return getLongestPath(r, l, d, u)
}

func checkPosition(maze Maze, visited map[Position]bool, p Position, path Path) Path {
	if visited[p] || p.x < 0 || p.y < 0 || p.x >= len(maze.m[0]) || p.y >= len(maze.m) || maze.m[p.y][p.x] == '.' {
		return path
	}

	path = append(path, p)
	visited[p] = true

	switch maze.m[p.y][p.x] {
	case '|':
		d := checkPosition(maze, visited, Position{p.x, p.y + 1}, path)
		u := checkPosition(maze, visited, Position{p.x, p.y - 1}, path)
		return getLongestPath(d, u)
	case '-':
		r := checkPosition(maze, visited, Position{p.x + 1, p.y}, path)
		l := checkPosition(maze, visited, Position{p.x - 1, p.y}, path)
		return getLongestPath(r, l)
	case 'F':
		r := checkPosition(maze, visited, Position{p.x + 1, p.y}, path)
		d := checkPosition(maze, visited, Position{p.x, p.y + 1}, path)
		return getLongestPath(r, d)
	case '7':
		l := checkPosition(maze, visited, Position{p.x - 1, p.y}, path)
		d := checkPosition(maze, visited, Position{p.x, p.y + 1}, path)
		return getLongestPath(l, d)
	case 'J':
		l := checkPosition(maze, visited, Position{p.x - 1, p.y}, path)
		u := checkPosition(maze, visited, Position{p.x, p.y - 1}, path)
		return getLongestPath(l, u)
	case 'L':
		r := checkPosition(maze, visited, Position{p.x + 1, p.y}, path)
		u := checkPosition(maze, visited, Position{p.x, p.y - 1}, path)
		return getLongestPath(r, u)
	case 'S':
		break
	}

	return path
}

func getLongestPath(path1, path2 Path, paths ...Path) Path {
	longestPath := path1

	if len(path2) > len(path1) {
		longestPath = path2
	}

	for _, path := range paths {
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}

	return longestPath
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
