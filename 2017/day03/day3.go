package main

import (
	"fmt"
	"math"
)

// Starting direction values
const (
	Right int = 1
	Up    int = 3
	Left  int = 5
	Down  int = 7
)

func main() {
	input := 361527
	partOne(input)
	partTwo(input)
}

func partOne(input int) {
	var right = countSteps(input, Right)
	var up = countSteps(input, Up)
	var left = countSteps(input, Left)
	var down = countSteps(input, Down)

	var val1 = int(math.Min(float64(right), float64(up)))
	var val2 = int(math.Min(float64(left), float64(down)))
	var val3 = int(math.Min(float64(val1), float64(val2)))

	fmt.Printf("\nSolution: %d\n", val3)
}

func partTwo(input int) {
	board := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// starting position
	x, y := 2, 2
	rightUp, downLeft := 1, 2
	steps := 0

	for {
		for i := 0; i < rightUp; i++ {
			y++
			steps++
			board[x][y] = checkAdjacent(x, y, board)
			if checkDone(board[x][y], input) {
				fmt.Printf("Step %d, value %d\n", steps, board[x][y])
				return
			}
		}

		for i := 0; i < rightUp; i++ {
			x--
			steps++
			board[x][y] = checkAdjacent(x, y, board)
			if checkDone(board[x][y], input) {
				fmt.Printf("Step %d, value %d\n", steps, board[x][y])
				return
			}
		}
		rightUp += 2

		for i := 0; i < downLeft; i++ {
			y--
			steps++
			board[x][y] = checkAdjacent(x, y, board)
			if checkDone(board[x][y], input) {
				fmt.Printf("Step %d, value %d\n", steps, board[x][y])
				return
			}
		}

		for i := 0; i < downLeft; i++ {
			x++
			steps++
			board[x][y] = checkAdjacent(x, y, board)
			if checkDone(board[x][y], input) {
				fmt.Printf("Step %d, value %d\n", steps, board[x][y])
				return
			}
		}
		downLeft += 2
		board = expand(board)
		x++
		y++
	}
}

func countSteps(target int, direction int) int {
	var steps = 0
	var num = 1
	var last, lastSteps int
	var diff = 8

	for increment := direction; num < target; steps++ {
		last = num
		num += increment
		increment += diff
		lastSteps = steps
	}

	for ; last < target; last++ {
		lastSteps++
	}

	for ; num > target; num-- {
		steps++
	}

	return int(math.Min(float64(steps), float64(lastSteps)))
}

func checkAdjacent(x, y int, board [][]int) int {
	sum := 0
	sum += board[x-1][y-1]
	sum += board[x-1][y]
	sum += board[x-1][y+1]
	sum += board[x][y-1]
	sum += board[x][y+1]
	sum += board[x+1][y-1]
	sum += board[x+1][y]
	sum += board[x+1][y+1]

	return sum
}

func expand(board [][]int) [][]int {
	// expand middle rows by appending 0 to front and back
	for i := 0; i < len(board); i++ {
		row := board[i]
		row = append([]int{0}, row...)
		row = append(row, 0)
		board[i] = row
	}

	// expand top and bottom by appending equal sized rows to top and bottom
	top := make([]int, len(board[0]))
	bottom := make([]int, len(board[0]))

	newBoard := make([][]int, 0)
	newBoard = append(newBoard, top)
	newBoard = append(newBoard, board...)
	newBoard = append(newBoard, bottom)

	return newBoard
}

func checkDone(num, target int) bool {
	return num > target
}
