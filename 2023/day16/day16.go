package day16

type point struct {
	x, y int
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

type mirrorDir rune

const (
	verticalSplit   mirrorDir = '|'
	horizontalSplit mirrorDir = '-'
	backslash       mirrorDir = '\\'
	slash           mirrorDir = '/'
)

type square struct {
	location  point
	energized bool
	mirror    mirrorDir
}

type beam struct {
	facing   direction
	location point
}

func solvePartOne(input []string) int {
	board := parseInput(input)
	simulate(board)

	sum := 0
	for _, row := range board {
		for _, square := range row {
			if square.energized {
				sum++
			}
		}
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}

func parseInput(input []string) [][]square {
	board := [][]square{}
	for y, line := range input {
		row := []square{}
		for x, char := range line {
			var mirror mirrorDir
			if char != '.' {
				mirror = mirrorDir(char)
			}

			row = append(row, square{
				location: point{x, y},
				mirror:   mirror,
			})
		}

		board = append(board, row)
	}

	return board
}

func simulate(board [][]square) {

	beams := []beam{
		{
			facing:   east,
			location: point{0, 0},
		},
	}

	// Keep track of when we've hit a splitter and actually split (didn't continue in the
	// same direction). If we hit this again we can ignore it
	hitSplitters := map[point]bool{}

	for i := 0; i < len(beams); i++ {
		laser := beams[i]

		for laser.location.x >= 0 && laser.location.x < len(board[0]) && laser.location.y >= 0 && laser.location.y < len(board) {
			// if we're on a square that's not energized, energize it
			board[laser.location.y][laser.location.x].energized = true

			// if we're on a mirror, change direction. If we're on a splitter, then
			// check that we're not aligned with the splitter and split, else continue in
			// the same direction
			if board[laser.location.y][laser.location.x].mirror == verticalSplit {
				if laser.facing == north {
					laser.location.y--
				} else if laser.facing == south {
					laser.location.y++
				} else {
					if _, ok := hitSplitters[laser.location]; ok {
						break
					}

					// split
					hitSplitters[laser.location] = true
					laser2 := beam{
						facing:   north,
						location: point{laser.location.x, laser.location.y - 1},
					}
					beams = append(beams, laser2)

					laser.facing = south
					laser.location.y++
				}
			} else if board[laser.location.y][laser.location.x].mirror == horizontalSplit {
				if laser.facing == east {
					laser.location.x++
				} else if laser.facing == west {
					laser.location.x--
				} else {
					if _, ok := hitSplitters[laser.location]; ok {
						break
					}

					// split
					hitSplitters[laser.location] = true

					laser2 := beam{
						facing:   east,
						location: point{laser.location.x + 1, laser.location.y},
					}
					beams = append(beams, laser2)

					laser.facing = west
					laser.location.x--
				}
			} else if board[laser.location.y][laser.location.x].mirror == backslash {
				if laser.facing == east {
					laser.facing = south
					laser.location.y++
				} else if laser.facing == west {
					laser.facing = north
					laser.location.y--
				} else if laser.facing == north {
					laser.facing = west
					laser.location.x--
				} else if laser.facing == south {
					laser.facing = east
					laser.location.x++
				}
			} else if board[laser.location.y][laser.location.x].mirror == slash {
				if laser.facing == east {
					laser.facing = north
					laser.location.y--
				} else if laser.facing == west {
					laser.facing = south
					laser.location.y++
				} else if laser.facing == north {
					laser.facing = east
					laser.location.x++
				} else if laser.facing == south {
					laser.facing = west
					laser.location.x--
				}
			} else {
				// if we're facing east, we're going to go right
				if laser.facing == east {
					laser.location.x++
				} else if laser.facing == west {
					laser.location.x--
				} else if laser.facing == north {
					laser.location.y--
				} else if laser.facing == south {
					laser.location.y++
				}
			}
		}
	}
}
