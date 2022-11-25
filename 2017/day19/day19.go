package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func main() {
	input := aocutil.ReadStringsFromFile("input.txt")
	fmt.Println(solve(input))
}

const (
	UP    = iota
	DOWN  = iota
	LEFT  = iota
	RIGHT = iota
)

func solve(maze []string) (string, int) {
	x := strings.Index(maze[0], "|")
	y := 0
	var lastX, lastY int
	currDir := DOWN
	output := ""
	inMaze := true
	steps := 0

	// keep going until we cannot proceed in any direction
	for inMaze {
		lastX = x
		lastY = y

		switch currDir {
		case UP:
			// if we can keep going our current direction do so
			if y-1 >= 0 {
				inMaze = canContinue(maze[y-1][x])
				y--
				steps++
			}

			if unicode.IsLetter(rune(maze[y][x])) { // add letter to output for part 1
				output = output + string(maze[y][x])
			} else if maze[y][x] == '+' { // or change directions if at an intersection
				currDir, _ = changeDirection(maze, currDir, x, y)
			}
		case DOWN:
			if y+1 < len(maze) {
				inMaze = canContinue(maze[y+1][x])
				y++
				steps++
			}
			if unicode.IsLetter(rune(maze[y][x])) {
				output = output + string(maze[y][x])
			} else if maze[y][x] == '+' {
				currDir, _ = changeDirection(maze, currDir, x, y)
			}
		case LEFT:
			if x-1 >= 0 {
				inMaze = canContinue(maze[y][x-1])
				x--
				steps++
			}
			if unicode.IsLetter(rune(maze[y][x])) {
				output = output + string(maze[y][x])
			} else if maze[y][x] == '+' {
				currDir, _ = changeDirection(maze, currDir, x, y)
			}
		case RIGHT:
			if x+1 < len(maze[y]) {
				inMaze = canContinue(maze[y][x+1])
				x++
				steps++
			}
			if unicode.IsLetter(rune(maze[y][x])) {
				output = output + string(maze[y][x])
			} else if maze[y][x] == '+' {
				currDir, _ = changeDirection(maze, currDir, x, y)
			}
		}

		// for debugging, stop if we haven't moved
		if lastX == x && lastY == y {
			return "stuck", -1
		}
	}

	return output, steps
}

func changeDirection(maze []string, currDir int, x int, y int) (int, error) {
	switch currDir {
	case UP:
		fallthrough
	case DOWN:
		if x+1 < len(maze[y]) && (maze[y][x+1] == '-' || unicode.IsLetter(rune(maze[y][x+1]))) {
			return RIGHT, nil
		} else if x-1 >= 0 && (maze[y][x-1] == '-' || unicode.IsLetter(rune(maze[y][x-1]))) {
			return LEFT, nil
		}
	case LEFT:
		fallthrough
	case RIGHT:
		if y+1 < len(maze) && (maze[y+1][x] == '|' || unicode.IsLetter(rune(maze[y+1][x]))) {
			return DOWN, nil
		} else if y-1 >= 0 && (maze[y-1][x] == '|') || unicode.IsLetter(rune(maze[y-1][x])) {
			return UP, nil
		}
	}
	return 0, errors.New("Invalid direction")
}

func canContinue(char byte) bool {
	if char == '|' ||
		char == '-' ||
		char == '+' ||
		unicode.IsLetter(rune(char)) {
		return true
	}
	return false
}
