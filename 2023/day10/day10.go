package day10

import (
	"math"
)

type point struct {
	x int
	y int
}

var (
	up    = point{x: 0, y: -1}
	down  = point{x: 0, y: 1}
	left  = point{x: -1, y: 0}
	right = point{x: 1, y: 0}
)

func (p point) add(p2 point) point {
	return point{x: p.x + p2.x, y: p.y + p2.y}
}

func (p point) isValid(max point) bool {
	return p.x >= 0 && p.x < max.x && p.y >= 0 && p.y < max.y
}

type node struct {
	pos    point
	ch     rune
	weight int
	up     *node
	down   *node
	left   *node
	right  *node
}

// Parse input and returns starting node
func parseInput(input []string) *node {
	var start *node
	nodes := map[point]*node{}

	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				pos := point{x: x, y: y}
				nodes[pos] = &node{
					pos:    pos,
					ch:     char,
					weight: math.MaxInt32,
				}

				if char == '|' || char == 'J' || char == 'L' || char == 'S' {
					above := pos.add(up)
					if above.isValid(point{x: len(line), y: len(input)}) {
						if _, ok := nodes[above]; ok && checkConnectionUp(char, nodes[above].ch) {
							nodes[pos].up = nodes[above]
							nodes[above].down = nodes[pos]
						}
					}
				}
				if char == 'J' || char == '-' || char == '7' || char == 'S' {
					toLeft := pos.add(left)
					if toLeft.isValid(point{x: len(line), y: len(input)}) {
						if _, ok := nodes[toLeft]; ok && checkConnectionLeft(char, nodes[toLeft].ch) {
							nodes[pos].left = nodes[toLeft]
							nodes[toLeft].right = nodes[pos]
						}
					}
				}

				if char == 'S' {
					start = nodes[pos]
					start.weight = 0
				}
			}
		}
	}

	return start
}

func checkConnectionUp(from, to rune) bool {
	return (from == '|' || from == 'J' || from == 'L' || from == 'S') && (to == '|' || to == '7' || to == 'F' || to == 'S')
}

func checkConnectionLeft(from, to rune) bool {
	return (from == '-' || from == '7' || from == 'J' || from == 'S') && (to == '-' || to == 'F' || to == 'L' || to == 'S')
}

// Traverse the graph and return the max weight and map of nodes included in the path
func traverse(start *node) (int, map[point]*node) {
	stack := []*node{start}
	max := 0

	nodes := map[point]*node{}
	nodes[start.pos] = start

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if current.weight > max {
			max = current.weight
		}

		if current.up != nil && current.up.weight > current.weight+1 {
			current.up.weight = current.weight + 1
			stack = append(stack, current.up)
			nodes[current.up.pos] = current.up
		}
		if current.down != nil && current.down.weight > current.weight+1 {
			current.down.weight = current.weight + 1
			stack = append(stack, current.down)
			nodes[current.down.pos] = current.down
		}
		if current.left != nil && current.left.weight > current.weight+1 {
			current.left.weight = current.weight + 1
			stack = append(stack, current.left)
			nodes[current.left.pos] = current.left
		}
		if current.right != nil && current.right.weight > current.weight+1 {
			current.right.weight = current.weight + 1
			stack = append(stack, current.right)
			nodes[current.right.pos] = current.right
		}
	}

	return max, nodes
}

func solvePartOne(input []string) int {
	start := parseInput(input)
	max, _ := traverse(start)
	return max
}

func solvePartTwo(input []string) int {
	start := parseInput(input)
	_, nodes := traverse(start)

	bigNodes := map[point]rune{}

	// all nodes, but big
	for y := range input {
		for yMul := 0; yMul < 3; yMul++ {
			for x := range input[y] {
				for xMul := 0; xMul < 3; xMul++ {
					if node, ok := nodes[point{x: x, y: y}]; ok {
						if node.ch == '-' {
							if yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						} else if node.ch == '|' {
							if xMul == 1 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						} else if node.ch == '7' {
							if xMul == 0 && yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else if xMul == 1 && yMul == 1 {
								// fmt.Print("7")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '7'
							} else if xMul == 1 && yMul == 2 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						} else if node.ch == 'J' {
							if xMul == 1 && yMul == 0 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else if xMul == 1 && yMul == 1 {
								// fmt.Print("J")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = 'J'
							} else if xMul == 0 && yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						} else if node.ch == 'L' {
							if xMul == 1 && yMul == 0 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else if xMul == 1 && yMul == 1 {
								// fmt.Print("L")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = 'L'
							} else if xMul == 2 && yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						} else if node.ch == 'F' {
							if xMul == 1 && yMul == 1 {
								// fmt.Print("F")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = 'F'
							} else if xMul == 2 && yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else if xMul == 1 && yMul == 2 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						} else if node.ch == 'S' {
							if node.up != nil && xMul == 1 && yMul == 0 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else if node.left != nil && xMul == 0 && yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else if node.down != nil && xMul == 1 && yMul == 2 {
								// fmt.Print("|")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '|'
							} else if node.right != nil && xMul == 2 && yMul == 1 {
								// fmt.Print("-")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '-'
							} else if xMul == 1 && yMul == 1 {
								// fmt.Print("S")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = 'S'
							} else {
								// fmt.Print(" ")
								bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
							}
						}
					} else if xMul == 1 && yMul == 1 {
						// fmt.Print(".")
						bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = '.'
					} else {
						// fmt.Print(" ")
						bigNodes[point{x: x*3 + xMul, y: y*3 + yMul}] = ' '
					}
				}
			}
			// fmt.Printf("\n")
		}
	}

	flood(point{x: 0, y: 0}, bigNodes)

	sum := 0
	for _, n := range bigNodes {
		if n == '.' {
			sum++
		}
	}

	return sum
}

func flood(p point, nodes map[point]rune) {
	if nodes[p] == ' ' || nodes[p] == '.' {
		nodes[p] = 'X'
		flood(p.add(up), nodes)
		flood(p.add(down), nodes)
		flood(p.add(left), nodes)
		flood(p.add(right), nodes)
	}
}
