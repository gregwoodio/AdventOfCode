package day10

import (
	"fmt"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func solvePartOne(input []string) int {
	asteroids := []aocutil.Coord{}
	grid := [][]bool{}
	counts := [][]int{}
	var bestLocation aocutil.Coord

	// build grid
	for y, line := range input {
		grid = append(grid, make([]bool, len(line)))
		counts = append(counts, make([]int, len(line)))

		for x, ch := range line {
			if ch == '#' {
				asteroids = append(asteroids, aocutil.Coord{
					X: x,
					Y: y,
				})
				grid[y][x] = true
			}
		}
	}

	// Find slope of each asteroid to every point along the outer edge
	// of the grid. Count the other asteroids the current one can see
	max := 0
	for _, ast := range asteroids {
		count := 0
		counted := [][]bool{}

		for _, line := range input {
			counted = append(counted, make([]bool, len(line)))
		}

		for x := 0; x < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				if x == ast.X && y == ast.Y {
					continue
				}

				target := aocutil.Coord{
					X: x,
					Y: y,
				}
				count += countNeighbours(&ast, &target, &grid, &counted, &counts)
			}
		}

		if count > max {
			max = count
			bestLocation = ast
		}
	}

	fmt.Printf("Best station location: %d, %d\n", bestLocation.X, bestLocation.Y)
	return max
}

func countNeighbours(current, target *aocutil.Coord, grid *[][]bool, counted *[][]bool, counts *[][]int) int {
	dx, dy := target.GetSlope(current)
	tx := current.X
	ty := current.Y
	count := 0

	if dx == 0 && dy == 0 {
		return count
	}

	for {
		tx += dx
		ty += dy

		// out of bounds
		if tx >= len((*grid)[0]) || ty >= len(*grid) || tx < 0 || ty < 0 {
			break
		}

		if (*grid)[ty][tx] && !(*counted)[ty][tx] {
			// square has an asteroid that hasn't been counted yet
			count++
			(*counts)[current.Y][current.X]++
			(*counted)[ty][tx] = true
			break
		} else if (*grid)[ty][tx] {
			// square has an asteroid that has been counted, so we're done searching
			break
		}
		// else keep searching
	}

	return count
}
