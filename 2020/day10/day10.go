package day10

import "sort"

func solvePartOne(input []int) int {
	input = append(input, 0)
	sort.Ints(input)

	// add three more than highest value
	input = append(input, input[len(input)-1]+3)

	ones, threes, prev := 0, 0, input[0]
	for i := 1; i < len(input); i++ {
		if input[i]-prev == 1 {
			ones++
		} else if input[i]-prev == 3 {
			threes++
		}

		prev = input[i]
	}

	return ones * threes
}

func solvePartTwo(input []int) int {
	input = append(input, 0)
	sort.Ints(input)

	// add three more than highest value
	input = append(input, input[len(input)-1]+3)

	paths := make(map[int]int)
	paths[0] = 1

	for _, num := range input[1:] {
		paths[num] = paths[num-1] + paths[num-2] + paths[num-3]
	}

	max := input[len(input)-1]
	return paths[max]
}
