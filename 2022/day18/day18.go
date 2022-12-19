package day18

import (
	"strconv"
	"strings"
)

type coord struct {
	x, y, z int
}

func solvePartOne(input []string) int {
	coords := parse(input)

	count := 0

	for c := range coords {
		// check sides
		comp := coord{x: c.x + 1, y: c.y, z: c.z}
		if _, ok := coords[comp]; !ok {
			count++
		}
		comp = coord{x: c.x - 1, y: c.y, z: c.z}
		if _, ok := coords[comp]; !ok {
			count++
		}
		comp = coord{x: c.x, y: c.y + 1, z: c.z}
		if _, ok := coords[comp]; !ok {
			count++
		}
		comp = coord{x: c.x, y: c.y - 1, z: c.z}
		if _, ok := coords[comp]; !ok {
			count++
		}
		comp = coord{x: c.x, y: c.y, z: c.z + 1}
		if _, ok := coords[comp]; !ok {
			count++
		}
		comp = coord{x: c.x, y: c.y, z: c.z - 1}
		if _, ok := coords[comp]; !ok {
			count++
		}
	}

	return count
}

type fill int

const (
	water fill = iota
	lava
	empty
)

func solvePartTwo(input []string) int {
	coords := parse(input)
	maxX, maxY, maxZ := 0, 0, 0

	for c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.z > maxZ {
			maxZ = c.z
		}
	}

	maxX++
	maxY++
	maxZ++

	// make a 3d array of either empty or lava
	area := [][][]fill{}
	for x := 0; x <= maxX; x++ {
		area = append(area, [][]fill{})
		for y := 0; y <= maxY; y++ {
			area[x] = append(area[x], []fill{})
			for z := 0; z <= maxZ; z++ {
				comp := coord{x: x, y: y, z: z}
				if _, ok := coords[comp]; ok {
					area[x][y] = append(area[x][y], lava)
				} else {
					area[x][y] = append(area[x][y], empty)
				}
			}
		}
	}

	// set the highest value to water then flood fill
	flood(&area, maxX, maxY, maxZ)

	// for x := 0; x < len(area); x++ {
	// 	fmt.Printf("\n=== X %d ===\n", x)

	// 	for y := 0; y < len(area[x]); y++ {
	// 		for z := 0; z < len(area[x][y]); z++ {
	// 			if area[x][y][z] == water {
	// 				fmt.Printf("~")
	// 			} else if area[x][y][z] == lava {
	// 				fmt.Printf("#")
	// 			} else {
	// 				fmt.Printf(".")
	// 			}
	// 		}
	// 		fmt.Printf("\n")
	// 	}
	// }

	count := 0

	for x := 0; x < len(area); x++ {
		for y := 0; y < len(area[x]); y++ {
			for z := 0; z < len(area[x][y]); z++ {
				if area[x][y][z] == lava {

					if x-1 < 0 || area[x-1][y][z] == water {
						count++
					}
					if x+1 >= len(area) || area[x+1][y][z] == water {
						count++
					}
					if y-1 < 0 || area[x][y-1][z] == water {
						count++
					}
					if y+1 >= len(area[x]) || area[x][y+1][z] == water {
						count++
					}
					if z-1 < 0 || area[x][y][z-1] == water {
						count++
					}
					if z+1 >= len(area[x][y]) || area[x][y][z+1] == water {
						count++
					}
				}
			}
		}
	}

	return count
}

func flood(area *[][][]fill, x, y, z int) {
	if x < 0 || y < 0 || z < 0 {
		return
	}

	if x >= len(*area) || y >= len((*area)[x]) || z >= len((*area)[x][y]) {
		return
	}

	if (*area)[x][y][z] != empty {
		// already filled with lava or water
		return
	}

	if (*area)[x][y][z] == empty {
		(*area)[x][y][z] = water
	}

	flood(area, x-1, y, z)
	flood(area, x+1, y, z)
	flood(area, x, y-1, z)
	flood(area, x, y+1, z)
	flood(area, x, y, z-1)
	flood(area, x, y, z+1)
}

func parse(input []string) map[coord]bool {
	spaces := map[coord]bool{}

	for _, line := range input {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err.Error())
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err.Error())
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err.Error())
		}

		c := coord{
			x: x,
			y: y,
			z: z,
		}
		spaces[c] = true
	}

	return spaces
}
