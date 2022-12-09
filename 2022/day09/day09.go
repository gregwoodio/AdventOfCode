package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func (p *pos) toString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (a *pos) isTouching(b *pos) bool {
	return a.x == b.x && a.y == b.y ||
		a.x == b.x+1 && a.y == b.y ||
		a.x == b.x-1 && a.y == b.y ||
		a.x == b.x && a.y == b.y+1 ||
		a.x == b.x && a.y == b.y-1 ||
		a.x == b.x+1 && a.y == b.y+1 ||
		a.x == b.x+1 && a.y == b.y-1 ||
		a.x == b.x-1 && a.y == b.y+1 ||
		a.x == b.x-1 && a.y == b.y-1
}

func (tail *pos) catchUp(head *pos) {
	if head.x == tail.x {
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
	} else if head.y == tail.y {
		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
	} else {
		if head.x > tail.x && head.y > tail.y {
			tail.x++
			tail.y++
		} else if head.x < tail.x && head.y > tail.y {
			tail.x--
			tail.y++
		} else if head.x < tail.x && head.y < tail.y {
			tail.x--
			tail.y--
		} else if head.x > tail.x && head.y < tail.y {
			tail.x++
			tail.y--
		}
	}
}

func solvePartOne(input []string) int {
	rope := []*pos{
		{x: 0, y: 0},
		{x: 0, y: 0},
	}

	return solve(input, rope)
}

func solvePartTwo(input []string) int {
	rope := []*pos{
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
	}

	return solve(input, rope)
}

func solve(input []string, rope []*pos) int {
	visited := map[string]bool{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Could not parse value")
		}

		for i := 0; i < steps; i++ {
			switch parts[0] {
			case "R":
				rope[0].x++
				break
			case "L":
				rope[0].x--
				break
			case "U":
				rope[0].y--
				break
			case "D":
				rope[0].y++
				break
			}
			for i := 1; i < len(rope); i++ {
				if !rope[i].isTouching(rope[i-1]) {
					rope[i].catchUp(rope[i-1])
				}
			}
			visited[rope[len(rope)-1].toString()] = true
		}
	}

	return len(visited)
}
