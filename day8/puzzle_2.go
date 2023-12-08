package day8

func solvePuzzle2(leMap Map) int {
	nodes := getStartingNodes(leMap)
	var steps []int

	for _, node := range nodes {
		steps = append(steps, stepsToZ(leMap, node))
	}

	return lcm(steps[0], steps[1], steps[2:]...)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func stepsToZ(leMap Map, startingNode string) int {
	counter := 0
	node := startingNode

	for {
		for _, direction := range leMap.instr {
			counter++
			if direction == L {
				node = leMap.dMap[node].l
			} else {
				node = leMap.dMap[node].r
			}

			if string(node[2]) == "Z" {
				return counter
			}
		}
	}
}

func getStartingNodes(leMap Map) []string {
	var nodes []string

	for node := range leMap.dMap {
		if string(node[2]) == "A" {
			nodes = append(nodes, node)
		}
	}

	return nodes
}
