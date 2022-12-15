package day14

import (
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func solvePartOne(input []string) int {
	board, maxY := drawBoard(input)
	return solve(board, maxY, false)
}

func solvePartTwo(input []string) int {
	board, maxY := drawBoard(input)
	return solve(board, maxY, true)
}

func solve(board map[coord]bool, maxY int, isPartTwo bool) int {
	count := 0
	start := coord{
		x: 500,
		y: 0,
	}

outer:
	for {
		sand := start
	inner:
		for {
			if sand.y == maxY {
				if isPartTwo {
					board[sand] = true
					count++
					break inner
				} else {
					break outer
				}
			}

			if isPartTwo {
				if _, ok := board[start]; ok {
					break outer
				}
			}

			next := coord{
				x: sand.x,
				y: sand.y + 1,
			}

			if _, ok := board[next]; ok {
				// blocked, check left
				next = coord{
					x: sand.x - 1,
					y: sand.y + 1,
				}

				if _, ok = board[next]; ok {
					// blocked, check right
					next = coord{
						x: sand.x + 1,
						y: sand.y + 1,
					}

					if _, ok = board[next]; ok {
						// blocked, add sand position to board
						board[sand] = true
						count++
						break inner
					} else {
						sand = next
					}
				} else {
					sand = next
				}
			} else {
				sand = next
			}
		}
	}

	return count
}

func drawBoard(input []string) (map[coord]bool, int) {
	coords := map[coord]bool{}
	maxY := 0
	for _, line := range input {
		parts := strings.Split(line, " -> ")

		for i := 1; i < len(parts); i++ {
			prev := toCoord(parts[i-1])
			coords[prev] = true

			curr := toCoord(parts[i])

			for prev != curr {
				if prev.x == curr.x {
					if prev.y < curr.y {
						prev = coord{
							y: prev.y + 1,
							x: prev.x,
						}

						if prev.y > maxY {
							maxY = prev.y
						}
					} else {
						prev = coord{
							y: prev.y - 1,
							x: prev.x,
						}
					}
				} else if prev.y == curr.y {
					if prev.x < curr.x {
						prev = coord{
							y: prev.y,
							x: prev.x + 1,
						}
					} else {
						prev = coord{
							y: prev.y,
							x: prev.x - 1,
						}
					}
				}

				coords[prev] = true
			}
		}
	}

	return coords, maxY + 1
}

func toCoord(input string) coord {
	vals := strings.Split(input, ",")
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err.Error())
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err.Error())
	}
	return coord{
		x: x,
		y: y,
	}
}
