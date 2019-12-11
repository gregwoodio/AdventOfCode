package day10

import "github.com/gregwoodio/aocutil"

func solvePartOne(input []string) int {
	asteroids := []aocutil.Coord{}
	grid := [][]bool{}
	counts := [][]int{}

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

		for x := 0; x < len(grid[0])-1; x++ {
			top := aocutil.Coord{
				X: x,
				Y: 0,
			}
			dx, dy := top.GetSlope(&ast)
			tx := ast.X
			ty := ast.Y

			if dx == 0 && dy == 0 {
				continue
			}

			for {
				tx += dx
				ty += dy

				// out of bounds
				if tx >= len(grid[0]) || ty >= len(grid) || tx < 0 || ty < 0 {
					break
				}

				if grid[ty][tx] && !counted[ty][tx] {
					count++
					counts[ast.Y][ast.X]++
					counted[ty][tx] = true
					break
				}
			}
		}

		for x := 0; x < len(grid[0])-1; x++ {
			bottom := aocutil.Coord{
				X: x,
				Y: len(grid) - 1,
			}
			dx, dy := bottom.GetSlope(&ast)
			tx := ast.X
			ty := ast.Y

			if dx == 0 && dy == 0 {
				continue
			}

			for {
				tx += dx
				ty += dy

				// out of bounds
				if tx >= len(grid[0]) || ty >= len(grid) || tx < 0 || ty < 0 {
					break
				}

				if grid[ty][tx] && !counted[ty][tx] {
					count++
					counts[ast.Y][ast.X]++
					counted[ty][tx] = true
					break
				}
			}
		}

		for y := 0; y < len(grid)-1; y++ {
			left := aocutil.Coord{
				X: 0,
				Y: y,
			}
			dx, dy := left.GetSlope(&ast)
			tx := ast.X
			ty := ast.Y

			if dx == 0 && dy == 0 {
				continue
			}

			for {
				tx += dx
				ty += dy

				// out of bounds
				if tx >= len(grid[0]) || ty >= len(grid) || tx < 0 || ty < 0 {
					break
				}

				if grid[ty][tx] && !counted[ty][tx] {
					count++
					counts[ast.Y][ast.X]++
					counted[ty][tx] = true
					break
				}
			}
		}

		for y := 0; y < len(grid)-1; y++ {
			bottom := aocutil.Coord{
				X: len(grid[0]) - 1,
				Y: y,
			}
			dx, dy := bottom.GetSlope(&ast)
			tx := ast.X
			ty := ast.Y

			if dx == 0 && dy == 0 {
				continue
			}

			for {
				tx += dx
				ty += dy

				// out of bounds
				if tx >= len(grid[0]) || ty >= len(grid) || tx < 0 || ty < 0 {
					break
				}

				if grid[ty][tx] && !counted[ty][tx] {
					count++
					counts[ast.Y][ast.X]++
					counted[ty][tx] = true
					break
				}
			}
		}

		if count > max {
			max = count
		}
	}

	return max
}
