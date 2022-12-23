package day22

import (
	"strconv"
)

type facing int

const (
	right facing = iota
	down
	left
	up
)

type terrain int

const (
	none  terrain = iota
	empty         // walkable space
	wall
)

func solvePartOne(input []string) int {
	height := len(input) - 2
	width := 0
	for i := 0; i < len(input)-2; i++ {
		if len(input[i]) > width {
			width = len(input[i])
		}
	}

	x, y := -1, 0
	f := right

	board := make([][]terrain, height)
	for i := 0; i < len(board); i++ {
		board[i] = make([]terrain, width)

		var j int
		for j = 0; j < len(input[i]); j++ {
			if input[i][j] == ' ' {
				board[i][j] = none
			} else if input[i][j] == '.' {
				if i == 0 && x == -1 {
					// set starting position if not yet set
					x = j
				}
				board[i][j] = empty
			} else if input[i][j] == '#' {
				board[i][j] = wall
			}
		}

		// fill in rest with none squares
		for ; j < width; j++ {
			board[i][j] = none
		}
	}

	instructions := splitInstructions(input[len(input)-1])

	for _, inst := range instructions {
		switch inst {
		case "R":
			f = facing((int(f) + 1) % 4)

		case "L":
			f = facing((int(f) + 3) % 4)
		default:
			val, err := strconv.Atoi(inst)
			if err != nil {
				panic(err.Error())
			}

			// continue val spaces in f direction
			dx, dy := x, y
			switch f {
			case right:
				for step := 1; step <= val; step++ {
					if board[dy][(dx+1)%len(board[dy])] == empty {
						dx = (dx + 1) % len(board[dy])
					} else if board[dy][(dx+1)%len(board[dy])] == wall {
						break
					} else {
						// none spaces, continue until wall or empty
						tx := (dx + 1) % len(board[dy])
						for board[dy][tx] == none {
							tx = (tx + 1) % len(board[dy])
						}

						if board[dy][tx] == empty {
							dx = tx
						}
					}
				}
				x = dx

			case down:
				for step := 1; step <= val; step++ {
					if board[(dy+1)%len(board)][dx] == empty {
						dy = (dy + 1) % len(board)
					} else if board[(dy+1)%len(board)][dx] == wall {
						break
					} else {
						// none spaces, continue until wall or empty
						ty := (dy + 1) % len(board)
						for board[ty][dx] == none {
							ty = (ty + 1 + len(board)) % len(board)
						}

						if board[ty][dx] == empty {
							dy = ty
						}
					}
				}
				y = dy

			case left:
				for step := 1; step <= val; step++ {
					if board[dy][(dx-1+len(board[dy]))%len(board[dy])] == empty {
						dx = (dx - 1 + len(board[dy])) % len(board[dy])
					} else if board[dy][(dx-1+len(board[dy]))%len(board[dy])] == wall {
						break
					} else {
						// none spaces, continue until wall or empty
						tx := (dx - 1 + len(board[dy])) % len(board[dy])
						for board[dy][tx] == none {
							tx = (tx - 1 + len(board[dy])) % len(board[dy])
						}

						if board[dy][tx] == empty {
							dx = tx
						}
					}
				}
				x = dx + len(board[dy])%len(board[dy])

			case up:
				for step := 1; step <= val; step++ {
					if board[(dy-1+len(board))%len(board)][dx] == empty {
						dy = (dy - 1 + len(board)) % len(board)
					} else if board[(dy-1+len(board))%len(board)][dx] == wall {
						break
					} else {
						// none spaces, continue until wall or empty
						ty := (dy - 1 + len(board)) % len(board)
						for board[ty][dx] == none {
							ty = (ty - 1 + len(board)) % len(board)
						}

						if board[ty][dx] == empty {
							dy = ty
						}
					}
				}
				y = dy + len(board)%len(board)
			}
		}

	}

	return (y+1)*1000 + (x+1)*4 + int(f)

}

func solvePartTwo(input []string) int {
	return -1
}

func splitInstructions(inst string) []string {
	instructions := []string{}
	i, j := 0, 0

	for ; i < len(inst); i, j = i+1, i+1 {
		if inst[i] == 'R' || inst[i] == 'L' {
			instructions = append(instructions, inst[i:i+1])

		} else {
			for ; i < len(inst) && inst[i] != 'R' && inst[i] != 'L'; i++ {
			}

			instructions = append(instructions, inst[j:i])
			i--
		}
	}

	return instructions
}
