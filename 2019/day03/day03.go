package day03

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type direction int

type instruction struct {
	dir   direction
	value int
}

const (
	up direction = iota
	right
	down
	left
)

func solvePartOne(input []string) int {
	wire1 := parseInstructions(input[0])
	wire2 := parseInstructions(input[1])

	wirePath1 := travel(wire1)
	wirePath2 := travel(wire2)

	minTaxiDistance := math.MaxInt32

	for key := range wirePath1 {
		if _, ok := wirePath2[key]; ok {
			distance := taxicab(splitKeyToCoordinate(key))
			if distance < minTaxiDistance {
				minTaxiDistance = distance
			}
		}
	}

	return minTaxiDistance
}

func solvePartTwo(input []string) int {
	wire1 := parseInstructions(input[0])
	wire2 := parseInstructions(input[1])

	wirePath1 := travel(wire1)
	wirePath2 := travel(wire2)

	minSteps := math.MaxInt32

	for key := range wirePath1 {
		if _, ok := wirePath2[key]; ok {
			steps := wirePath1[key] + wirePath2[key]
			if steps < minSteps {
				minSteps = steps
			}
		}
	}

	return minSteps
}

func parseInstructions(input string) []instruction {
	instructions := []instruction{}

	values := strings.Split(input, ",")
	for _, value := range values {
		var dir direction
		switch value[0] {
		case 'U':
			dir = up
			break
		case 'R':
			dir = right
			break
		case 'D':
			dir = down
			break
		case 'L':
			dir = left
			break
		}

		num, err := strconv.Atoi(value[1:])
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, instruction{
			dir:   dir,
			value: num,
		})
	}

	return instructions
}

func travel(instructions []instruction) map[string]int {
	var x, y int
	var change func(int, int) (int, int)
	visited := map[string]int{}
	step := 1

	for _, inst := range instructions {
		switch inst.dir {
		case up:
			change = func(x, y int) (int, int) {
				y++
				return x, y
			}
			break
		case right:
			change = func(x, y int) (int, int) {
				x++
				return x, y
			}
			break
		case down:
			change = func(x, y int) (int, int) {
				y--
				return x, y
			}
			break
		case left:
			change = func(x, y int) (int, int) {
				x--
				return x, y
			}
			break
		}

		for i := 0; i < inst.value; i++ {
			x, y = change(x, y)
			key := fmt.Sprintf("%d,%d", x, y)
			if _, ok := visited[key]; !ok {
				// only record steps for unvisited coordinates
				visited[key] = step
			}
			step++
		}
	}

	return visited
}

func taxicab(x, y int) int {
	return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
}

func splitKeyToCoordinate(key string) (int, int) {
	nums := strings.Split(key, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return x, y
}
