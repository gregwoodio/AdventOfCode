package day17

import (
	"math"
)

type point struct {
	x, y, z, w int
}

func solvePartOne(input []string) int {
	points := parseInput(input)

	for i := 1; i <= 6; i++ {
		next := make(map[point]bool)
		lowX, highX, lowY, highY, lowZ, highZ, _, _ := findLowsAndHighs(points)

		// fmt.Printf("After %d cycles:\n", i)

		for z := lowZ - 1; z <= highZ+1; z++ {
			// fmt.Printf("z=%d\n", z)
			for y := lowY - 1; y <= highY+1; y++ {
				// fmt.Printf("y=%5d ", y)

				for x := lowX - 1; x <= highX+1; x++ {
					p := point{
						x: x,
						y: y,
						z: z,
					}
					neighbours := countNeighbours(p, points)

					if neighbours == 3 && !points[p] {
						next[p] = true
					} else if (neighbours == 2 || neighbours == 3) && points[p] {
						next[p] = true
					}

					// if next[p] {
					// 	fmt.Print("#")
					// } else {
					// 	fmt.Print(".")
					// }
				}

				// fmt.Println()
			}

			// fmt.Println()
		}

		points = next
	}

	return len(points)
}

func solvePartTwo(input []string) int {
	points := parseInput(input)

	for i := 1; i <= 6; i++ {
		next := make(map[point]bool)
		lowX, highX, lowY, highY, lowZ, highZ, lowW, highW := findLowsAndHighs(points)

		for w := lowW - 1; w <= highW+1; w++ {
			for z := lowZ - 1; z <= highZ+1; z++ {
				for y := lowY - 1; y <= highY+1; y++ {
					for x := lowX - 1; x <= highX+1; x++ {
						p := point{
							x: x,
							y: y,
							z: z,
							w: w,
						}
						neighbours := countNeighbours4D(p, points)

						if neighbours == 3 && !points[p] {
							next[p] = true
						} else if (neighbours == 2 || neighbours == 3) && points[p] {
							next[p] = true
						}
					}
				}
			}
		}

		points = next
	}

	return len(points)
}

func parseInput(input []string) map[point]bool {
	points := make(map[point]bool)

	for y, line := range input {
		for x, ch := range line {
			if ch == '#' {
				p := point{
					x: x,
					y: y,
					z: 0,
					w: 0,
				}

				points[p] = true
			}
		}
	}

	return points
}

func findLowsAndHighs(points map[point]bool) (int, int, int, int, int, int, int, int) {
	lowX, highX, lowY, highY, lowZ, highZ, lowW, highW := math.MaxInt32, 0, math.MaxInt32, 0, math.MaxInt32, 0, math.MaxInt32, 0

	for k := range points {
		if k.x < lowX {
			lowX = k.x
		}

		if k.x > highX {
			highX = k.x
		}

		if k.y < lowY {
			lowY = k.y
		}

		if k.y > highY {
			highY = k.y
		}

		if k.z < lowZ {
			lowZ = k.z
		}

		if k.z > highZ {
			highZ = k.z
		}

		if k.w < lowW {
			lowW = k.w
		}

		if k.w > highW {
			highW = k.w
		}
	}

	return lowX, highX, lowY, highY, lowZ, highZ, lowW, highW
}

func countNeighbours(p point, points map[point]bool) int {
	neighbours := 0

	for z := p.z - 1; z <= p.z+1; z++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for x := p.x - 1; x <= p.x+1; x++ {
				target := point{
					z: z,
					y: y,
					x: x,
				}

				if p.x == x && p.y == y && p.z == z {
					continue
				}

				if points[target] {
					neighbours++
				}
			}
		}
	}

	return neighbours
}

func countNeighbours4D(p point, points map[point]bool) int {
	neighbours := 0

	for w := p.w - 1; w <= p.w+1; w++ {
		for z := p.z - 1; z <= p.z+1; z++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for x := p.x - 1; x <= p.x+1; x++ {
					target := point{
						w: w,
						z: z,
						y: y,
						x: x,
					}

					if p.x == x && p.y == y && p.z == z && p.w == w {
						continue
					}

					if points[target] {
						neighbours++
					}
				}
			}
		}
	}

	return neighbours
}
