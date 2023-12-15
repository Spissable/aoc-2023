package day10

func solvePuzzle2(maze Maze) int {
	loop := getLoop(maze)
	// Shoelace formula
	trailing := 0
	for i := 0; i < len(loop); i++ {
		if i == len(loop)-1 {
			trailing += (loop[i].x + loop[0].x) * (loop[i].y - loop[0].y)
		} else {
			trailing += (loop[i].x + loop[i+1].x) * (loop[i].y - loop[i+1].y)
		}
	}
	area := trailing / 2

	// Pick's theorem
	return area - (len(loop) / 2) + 1
}
