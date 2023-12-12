package day11

type point struct {
	x, y int
}

type galaxy struct {
	initial  point
	expanded point
}

func solvePartOne(input []string) int {
	// h, w := len(input), len(input[0])
	initial := [][]bool{}
	galaxies := []galaxy{}

	for y, row := range input {
		initial = append(initial, []bool{})
		for x, char := range row {
			initial[y] = append(initial[y], char == '#')

			if char == '#' {
				galaxies = append(galaxies, galaxy{
					initial:  point{x, y},
					expanded: point{x, y},
				})
			}
		}
	}

	rowsToAdd := []int{}
	for i, row := range initial {
		add := true
		for _, x := range row {
			if x {
				add = false
			}
		}

		if add {
			rowsToAdd = append(rowsToAdd, i)
		}
	}

	colsToAdd := []int{}
	for i := 0; i < len(initial[0]); i++ {
		add := true
		for _, row := range initial {
			if row[i] {
				add = false
			}
		}

		if add {
			colsToAdd = append(colsToAdd, i)
		}
	}

	// expand universe
	for _, col := range colsToAdd {
		for i, galaxy := range galaxies {
			if galaxy.initial.x > col {
				galaxies[i].expanded.x++
			}
		}
	}
	for _, row := range rowsToAdd {
		for i, galaxy := range galaxies {
			if galaxy.initial.y > row {
				galaxies[i].expanded.y++
			}
		}
	}

	pairs := getGalaxyPairs(galaxies)
	sum := 0
	for _, pair := range pairs {
		sum += getDistance(pair[0], pair[1])
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}

// takes a list of galaxies and returns all pairs of galaxies
func getGalaxyPairs(galaxies []galaxy) [][]galaxy {
	pairs := [][]galaxy{}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, []galaxy{galaxies[i], galaxies[j]})
		}
	}

	return pairs
}

// gets taxicab distance between two galaxies expanded points
func getDistance(g1, g2 galaxy) int {
	return abs(g1.expanded.x-g2.expanded.x) + abs(g1.expanded.y-g2.expanded.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
