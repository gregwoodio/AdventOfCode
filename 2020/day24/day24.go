package day24

type tile struct {
	x int
	y int
}

// Tiles link together like this:
//    (-1, -1)   (0, -1)  (1, -1)
// (-1, 0)  (0, 0)   (1, 0)
//    (0, 1)    (1, 1)

func solvePartOne(input []string) int {
	tiles := setup(input)

	sum := 0
	for _, val := range tiles {
		if val {
			sum++
		}
	}
	return sum
}

func solvePartTwo(input []string) int {
	tiles := setup(input)
	dirs := getDirections()

	for i := 0; i < 100; i++ {
		tmp := make(map[tile]bool)
		next := make(map[tile]bool)

		for t, f := range tiles {
			if f {
				tmp[t] = true
			}

			for _, d := range []string{"e", "se", "sw", "w", "nw", "ne"} {
				neighbour := tile{x: t.x + dirs[d].x, y: t.y + dirs[d].y}
				if _, found := tmp[neighbour]; !found {
					tmp[neighbour] = false
				}
			}
		}

		for t, flipped := range tmp {
			neighbours := 0
			for _, d := range []string{"e", "se", "sw", "w", "nw", "ne"} {
				newTile := tile{x: t.x + dirs[d].x, y: t.y + dirs[d].y}
				f := tiles[newTile]
				if f {
					neighbours++
				}
			}

			if flipped && (neighbours == 1 || neighbours == 2) {
				next[t] = true
			} else if !flipped && neighbours == 2 {
				next[t] = true
			}
		}

		tiles = next
	}

	sum := 0
	for _, val := range tiles {
		if val {
			sum++
		}
	}
	return sum
}

func parseInstructions(input string) []string {
	inst := []string{}

	for i := 0; i < len(input); i++ {
		if input[i] == 'n' || input[i] == 's' {
			inst = append(inst, input[i:i+2])
			i++
		} else {
			inst = append(inst, input[i:i+1])
		}
	}

	return inst
}

func setup(input []string) map[tile]bool {
	reference := tile{x: 0, y: 0}
	tiles := make(map[tile]bool)
	tiles[reference] = false
	dirs := getDirections()

	for _, line := range input {
		instructions := parseInstructions(line)
		curr := reference
		for _, inst := range instructions {
			curr = tile{x: curr.x + dirs[inst].x, y: curr.y + dirs[inst].y}
			if _, ok := tiles[curr]; !ok {
				tiles[curr] = false
			}
		}

		tiles[curr] = !tiles[curr]
	}

	// only track black tiles
	for t, v := range tiles {
		if !v {
			delete(tiles, t)
		}
	}

	return tiles
}

func getDirections() map[string]tile {
	return map[string]tile{
		"e":  {x: 1, y: 0},
		"se": {x: 1, y: 1},
		"sw": {x: 0, y: 1},
		"w":  {x: -1, y: 0},
		"nw": {x: -1, y: -1},
		"ne": {x: 0, y: -1},
	}
}
