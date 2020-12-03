package day03

type point struct {
	x, y int
}

func solvePartOne(input []string) int {
	return toboggan(3, 1, input)
}

func solvePartTwo(input []string) int {
	prod := toboggan(1, 1, input)
	prod *= toboggan(3, 1, input)
	prod *= toboggan(5, 1, input)
	prod *= toboggan(7, 1, input)
	prod *= toboggan(1, 2, input)

	return prod
}

func toboggan(right, down int, input []string) int {
	height := len(input)
	width := len(input[0])
	location := point{
		x: 0,
		y: 0,
	}
	hill := makeHill(input)

	// slope is right 3, down 1
	treeCount := 0
	for location.y < height {
		location.y += down
		location.x = (location.x + right) % width

		if hill[location] {
			treeCount++
		}
	}

	return treeCount
}

func makeHill(inputs []string) map[point]bool {
	hill := make(map[point]bool)

	for row, input := range inputs {
		for col, ch := range input {
			p := point{
				x: col,
				y: row,
			}

			// true for trees
			hill[p] = ch == '#'
		}
	}

	return hill
}
