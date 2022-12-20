package day17

import "fmt"

type tetromino struct {
	points []int
	width  int
	height int
}

func (t tetromino) copy() tetromino {
	points := []int{}
	for _, p := range t.points {
		points = append(points, p)
	}

	return tetromino{
		points: points,
		width:  t.width,
		height: t.height,
	}
}

var horizontalLine = tetromino{
	points: []int{
		0b1111,
	},
	width:  4,
	height: 1,
}
var plus = tetromino{
	points: []int{
		0b010,
		0b111,
		0b010,
	},
	width:  3,
	height: 3,
}
var backwardsL = tetromino{
	points: []int{
		0b111,
		0b001,
		0b001,
	},
	width:  3,
	height: 3,
}
var verticalLine = tetromino{
	points: []int{
		0b1,
		0b1,
		0b1,
		0b1,
	},
	width:  1,
	height: 4,
}
var square = tetromino{
	points: []int{
		0b11,
		0b11,
	},
	width:  2,
	height: 2,
}

func solvePartOne(input string) int {
	rocks := []tetromino{horizontalLine, plus, backwardsL, verticalLine, square}
	windIndex := 0
	board := []int{}

	for r := 0; r < 2022; r++ {
		curr := rocks[r%len(rocks)].copy()
		// highest is the index of the top-most row of the piece
		highest := len(board) + 2 + curr.height
		left := 5

		for i, p := range curr.points {
			curr.points[i] = p << (5 - curr.width)
		}

		for {
			wind := input[windIndex]

			// wind blows first
			temp := curr.copy()
			switch wind {
			case '<':
				if left < 7 {
					for i, p := range curr.points {
						temp.points[i] = p << 1
					}
					if !checkOverlap(temp, board, highest) {
						curr = temp
						left++
					}
				}

				break
			case '>':
				if left-temp.width > 0 {
					for i, p := range curr.points {
						temp.points[i] = p >> 1
					}
					if !checkOverlap(temp, board, highest) {
						curr = temp
						left--
					}
				}
				break
			}
			windIndex++
			if windIndex >= len(input) {
				windIndex = 0
			}

			// fall down one
			if highest-1 < 0 || checkOverlap(curr, board, highest-1) {
				// blocked! add the rock to the board
				for i, p := highest-curr.height+1, 0; p < len(curr.points); i, p = i+1, p+1 {
					if len(board) == 0 || i < 0 {
						board = append(board, curr.points[p])
					} else if i >= len(board) {
						board = append(board, curr.points[p])
					} else {
						board[i] |= curr.points[p]
					}
				}
				// printBoard(board, r)
				break
				// } else if highest == 0 {
				// 	// blocked! add the rock to the board
				// 	for i, p := highest, 0; p < len(curr.points); i, p = i+1, p+1 {
				// 		if len(board) > i {
				// 			board[i] |= curr.points[p]
				// 		} else {
				// 			board = append(board, curr.points[p])
				// 		}
				// 	}
				// 	printBoard(board, r)

				// 	break
			}

			highest--
		}

	}

	return len(board)
}

func solvePartTwo(input string) int {
	return -1
}

// returns true if the piece overlaps with the board
func checkOverlap(curr tetromino, board []int, top int) bool {
	for i := 0; i < len(curr.points); i++ {
		// if top-(len(curr.points)-i-1) < 0 {
		// 	return true
		// }
		if top-i >= len(board) {
			// beyond the top of the board, so it's not overlapping
			continue
		}
		if top-i < 0 || board[top-i]&curr.points[len(curr.points)-1-i] > 0 {
			return true
		}
	}

	return false
}

func printBoard(board []int, rock int) {
	fmt.Printf("Rock #%d\n", rock)

	for r := len(board) - 1; r >= 0; r-- {
		row := board[r]
		for i := 6; i >= 0; i-- {
			if row>>i&1 == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("- %d\n", r)
	}
	fmt.Printf("\n")

}
