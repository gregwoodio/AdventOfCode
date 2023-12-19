package day18

import (
	"strconv"
	"strings"
)

type fillType int

const (
	unknown fillType = iota
	filled
	cleared
)

type point struct {
	x, y int
}

func solvePartOne(input []string) int {
	return dig(input)
}

func solvePartTwo(input []string) int {
	return -1
}

func dig(input []string) int {
	lowestX := 0
	highestX := 0
	lowestY := 0
	highestY := 0

	curr := point{0, 0}
	boardMap := map[point]bool{}
	boardMap[curr] = true

	for _, line := range input {
		parts := strings.Split(line, " ")
		var move point
		switch parts[0] {
		case "R":
			move = point{1, 0}
		case "L":
			move = point{-1, 0}
		case "U":
			move = point{0, -1}
		case "D":
			move = point{0, 1}
		}

		value, _ := strconv.Atoi(parts[1])
		for i := 0; i < value; i++ {
			curr = point{curr.x + move.x, curr.y + move.y}
			boardMap[curr] = true

			if curr.x < lowestX {
				lowestX = curr.x
			}

			if curr.x > highestX {
				highestX = curr.x
			}

			if curr.y < lowestY {
				lowestY = curr.y
			}

			if curr.y > highestY {
				highestY = curr.y
			}
		}
	}

	yLen := abs(highestY) + abs(lowestY)
	xLen := abs(highestX) + abs(lowestX)
	board := make([][]fillType, yLen+1)

	for y := 0; y <= yLen; y++ {
		board[y] = make([]fillType, xLen+1)
		for x := 0; x <= xLen; x++ {
			if boardMap[point{x + lowestX, y + lowestY}] {
				// fmt.Printf("#")
				board[y][x] = cleared
			} else {
				board[y][x] = unknown
			}
		}
	}

	// flood fill from every exterior point
	for y := 0; y <= yLen; y++ {
		floodFill(board, point{0, y})
		floodFill(board, point{xLen, y})
	}
	for x := 0; x <= xLen; x++ {
		floodFill(board, point{x, 0})
		floodFill(board, point{x, yLen})
	}

	sum := 0
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == filled {
				// fmt.Printf(".")
			}
			if board[y][x] == cleared {
				// fmt.Printf("#")
				sum++
			}
			if board[y][x] == unknown {
				// fmt.Printf("?")
				sum++
			}
		}
		// fmt.Println()
	}

	return sum
}

// Flood fill all the unknowns with filled.
func floodFill(board [][]fillType, p point) {
	if board[p.y][p.x] == cleared || board[p.y][p.x] == filled {
		return
	}

	board[p.y][p.x] = filled

	if p.x > 0 {
		floodFill(board, point{p.x - 1, p.y})
	}
	if p.x < len(board[0])-1 {
		floodFill(board, point{p.x + 1, p.y})
	}
	if p.y > 0 {
		floodFill(board, point{p.x, p.y - 1})
	}
	if p.y < len(board)-1 {
		floodFill(board, point{p.x, p.y + 1})
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
