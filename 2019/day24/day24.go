package day24

import "fmt"

func solvePartOne(input []string) int {
	bio := make(map[int]bool)
	grid := makeGrid(input)

	for {
		if _, ok := bio[grid]; ok {
			return grid
		}
		bio[grid] = true

		grid = iterate(grid)
	}
}

func makeGrid(input []string) int {
	grid := 0

	for i := range input {

		for j := range input[i] {
			if input[i][j] == '#' {
				grid |= 1 << uint(i*5+j)
			}
		}
	}

	return grid
}

func iterate(grid int) int {
	newGrid := 0

	for i := 0; i < 25; i++ {
		sum := 0

		// check surrounding
		if i-5 >= 0 && (1<<uint(i-5))&grid > 0 {
			sum++
		}
		if i+5 < 25 && (1<<uint(i+5))&grid > 0 {
			sum++
		}
		if ((i-1)%5) < (i%5) && (1<<uint(i-1))&grid > 0 {
			sum++
		}
		if ((i+1)%5) > (i%5) && (1<<uint(i+1))&grid > 0 {
			sum++
		}

		if grid&(1<<uint(i)) > 0 && sum == 1 {
			newGrid |= (1 << uint(i))
		} else if grid&(1<<uint(i)) == 0 && (sum == 1 || sum == 2) {
			newGrid |= (1 << uint(i))
		}
	}

	return newGrid
}

func printGrid(grid int) {
	for i := 0; i < 25; i++ {
		if grid&(1<<uint(i)) > 0 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if i%5 == 4 {
			fmt.Print("\n")
		}
	}
}
