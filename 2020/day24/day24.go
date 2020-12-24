package day24

const (
	e int = iota
	se
	sw
	w
	nw
	ne
)

type tile struct {
	x int
	y int
}

// Tiles link together like this:
//    (-1, -1)   (0, -1)  (1, -1)
// (-1, 0)  (0, 0)   (1, 0)
//    (0, 1)    (1, 1)

func solvePartOne(input []string) int {
	sum := 0
	tiles := setup(input)

	for _, val := range tiles {
		if val {
			sum++
		}
	}
	return sum
}

func solvePartTwo(input []string) int {
	return -1
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

	for _, line := range input {
		instructions := parseInstructions(line)
		curr := reference
		for _, inst := range instructions {
			switch inst {
			case "e":
				curr = tile{x: curr.x + 1, y: curr.y}
				if _, ok := tiles[curr]; !ok {
					tiles[curr] = false
				}
				break

			case "se":
				curr = tile{x: curr.x + 1, y: curr.y + 1}
				if _, ok := tiles[curr]; !ok {
					tiles[curr] = false
				}
				break

			case "sw":
				curr = tile{x: curr.x, y: curr.y + 1}
				if _, ok := tiles[curr]; !ok {
					tiles[curr] = false
				}
				break

			case "w":
				curr = tile{x: curr.x - 1, y: curr.y}
				if _, ok := tiles[curr]; !ok {
					tiles[curr] = false
				}
				break

			case "nw":
				curr = tile{x: curr.x - 1, y: curr.y - 1}
				if _, ok := tiles[curr]; !ok {
					tiles[curr] = false
				}
				break

			case "ne":
				curr = tile{x: curr.x, y: curr.y - 1}
				if _, ok := tiles[curr]; !ok {
					tiles[curr] = false
				}
				break
			}
		}

		tiles[curr] = !tiles[curr]
	}

	return tiles
}
