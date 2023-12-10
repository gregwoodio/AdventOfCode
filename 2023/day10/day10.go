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

func solvePartOne(input []string) int {
	start := parseInput(input)

	stack := []*node{start}
	max := 0

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if current.weight > max {
			max = current.weight
		}

		if current.up != nil && current.up.weight > current.weight+1 {
			current.up.weight = current.weight + 1
			stack = append(stack, current.up)
		}
		if current.down != nil && current.down.weight > current.weight+1 {
			current.down.weight = current.weight + 1
			stack = append(stack, current.down)
		}
		if current.left != nil && current.left.weight > current.weight+1 {
			current.left.weight = current.weight + 1
			stack = append(stack, current.left)
		}
		if current.right != nil && current.right.weight > current.weight+1 {
			current.right.weight = current.weight + 1
			stack = append(stack, current.right)
		}
	}

	return max
}

func solvePartTwo(input []string) int {
	return -1
}
