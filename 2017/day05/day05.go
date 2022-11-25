package main

import (
	"fmt"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func main() {
	input := aocutil.ReadIntsFromFile("input.txt")
	fmt.Println(partOne(input))

	input = aocutil.ReadIntsFromFile("input.txt")
	fmt.Println(partTwo(input))
}

func partOne(input []int) int {
	return solve(input, false)
}

func partTwo(input []int) int {
	return solve(input, true)
}

func solve(input []int, isPartTwo bool) int {
	pointer, steps := 0, 0
	for pointer < len(input) {
		step := input[pointer]
		increment := 1
		if isPartTwo && step >= 3 {
			increment = -1
		}
		input[pointer] += increment
		pointer += step
		steps++
	}
	return steps
}
