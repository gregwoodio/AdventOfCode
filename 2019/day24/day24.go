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

func solvePartTwo(input []string, time int) int {
	depth := make(map[int]int)
	depth[0] = makeGrid(input)

	for tick := 0; tick < time; tick++ {
		newDepth := make(map[int]int)

		//check inner
		for d := 0; ; d++ {
			if _, ok := depth[d]; !ok {
				// Layer doesn't exist yet, check if we need to create or break.
				// Bugs come alive when next to one or two live bugs. If the
				// 7, 11, 13 or 17th bit is flipped in the current layer, create it.
				if depth[d-1]&(1<<7) > 0 ||
					depth[d-1]&(1<<11) > 0 ||
					depth[d-1]&(1<<13) > 0 ||
					depth[d-1]&(1<<17) > 0 {
					newDepth[d] = iteratePartTwo(d, &depth)
				} else {
					break
				}
			} else {
				// layer exists, process
				newDepth[d] = iteratePartTwo(d, &depth)
			}
		}

		//check outer
		for d := -1; ; d-- {
			if _, ok := depth[d]; !ok {
				// Layer doesn't exist yet, check if we need to create or break.
				// Here we need to check all the outer bits to see if we need to
				// create a new layer.
				if depth[d+1]&(1<<0) > 0 ||
					depth[d+1]&(1<<1) > 0 ||
					depth[d+1]&(1<<2) > 0 ||
					depth[d+1]&(1<<3) > 0 ||
					depth[d+1]&(1<<4) > 0 ||
					depth[d+1]&(1<<5) > 0 ||
					depth[d+1]&(1<<9) > 0 ||
					depth[d+1]&(1<<10) > 0 ||
					depth[d+1]&(1<<14) > 0 ||
					depth[d+1]&(1<<15) > 0 ||
					depth[d+1]&(1<<19) > 0 ||
					depth[d+1]&(1<<20) > 0 ||
					depth[d+1]&(1<<21) > 0 ||
					depth[d+1]&(1<<22) > 0 ||
					depth[d+1]&(1<<23) > 0 ||
					depth[d+1]&(1<<24) > 0 {
					newDepth[d] = iteratePartTwo(d, &depth)
				} else {
					break
				}
			} else {
				newDepth[d] = iteratePartTwo(d, &depth)
			}
		}

		depth = newDepth
	}

	// count bugs
	sum := 0
	for d := 0; ; d-- {
		if _, ok := depth[d]; !ok {
			break
		}

		for i := 0; i < 25; i++ {
			if depth[d]&(1<<uint(i)) > 0 {
				sum++
			}
		}
	}

	for d := 1; ; d++ {
		if _, ok := depth[d]; !ok {
			break
		}

		for i := 0; i < 25; i++ {
			if depth[d]&(1<<uint(i)) > 0 {
				sum++
			}
		}
	}

	return sum
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

func iteratePartTwo(index int, layers *map[int]int) int {
	newGrid := 0
	below := (*layers)[index+1]
	curr := (*layers)[index]
	above := (*layers)[index-1]

	for i := 0; i < 25; i++ {
		sum := 0

		if i == 12 {
			continue
		}

		// check above
		if i < 5 {
			if above&(1<<7) > 0 {
				sum++
			}
		} else if i == 17 {
			// check inner bottom row
			if below&(1<<20) > 0 {
				sum++
			}
			if below&(1<<21) > 0 {
				sum++
			}
			if below&(1<<22) > 0 {
				sum++
			}
			if below&(1<<23) > 0 {
				sum++
			}
			if below&(1<<24) > 0 {
				sum++
			}
		} else {
			if curr&(1<<uint(i-5)) > 0 {
				sum++
			}
		}

		// check right
		if i%5 == 4 {
			if above&(1<<13) > 0 {
				sum++
			}
		} else if i == 11 {
			// check inner left column
			if below&(1<<0) > 0 {
				sum++
			}
			if below&(1<<5) > 0 {
				sum++
			}
			if below&(1<<10) > 0 {
				sum++
			}
			if below&(1<<15) > 0 {
				sum++
			}
			if below&(1<<20) > 0 {
				sum++
			}
		} else {
			if curr&(1<<uint(i+1)) > 0 {
				sum++
			}
		}

		// check down
		if i >= 20 {
			if above&(1<<17) > 0 {
				sum++
			}
		} else if i == 7 {
			// check inner top row
			if below&(1<<0) > 0 {
				sum++
			}
			if below&(1<<1) > 0 {
				sum++
			}
			if below&(1<<2) > 0 {
				sum++
			}
			if below&(1<<3) > 0 {
				sum++
			}
			if below&(1<<4) > 0 {
				sum++
			}
		} else {
			if curr&(1<<uint(i+5)) > 0 {
				sum++
			}
		}

		// check left
		if i%5 == 0 {
			if above&(1<<11) > 0 {
				sum++
			}
		} else if i == 13 {
			// check inner right column
			if below&(1<<4) > 0 {
				sum++
			}
			if below&(1<<9) > 0 {
				sum++
			}
			if below&(1<<14) > 0 {
				sum++
			}
			if below&(1<<19) > 0 {
				sum++
			}
			if below&(1<<24) > 0 {
				sum++
			}
		} else {
			if curr&(1<<uint(i-1)) > 0 {
				sum++
			}
		}

		if curr&(1<<uint(i)) > 0 && sum == 1 {
			// bug dies unless there is exactly one bug beside it
			newGrid |= (1 << uint(i))
		} else if curr&(1<<uint(i)) == 0 && (sum == 1 || sum == 2) {
			// new bug appears if the space is empty and there are one or
			// two bugs beside it
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
