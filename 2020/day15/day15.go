package day15

type info struct {
	lastTurnSaid int
	timesSaid    int
	next         int
}

func solvePartOne(input []int) int {
	return solve(input, 2020)
}

func solvePartTwo(input []int) int {
	return solve(input, 30000000)
}

func solve(input []int, limit int) int {
	numbers := make(map[int]*info)

	// setup starting nums
	turn := 1
	for _, num := range input {
		numbers[num] = &info{
			lastTurnSaid: turn,
			timesSaid:    1,
			next:         turn,
		}
		turn++
	}

	last := input[len(input)-1]
	for turn <= limit {
		if numbers[last].timesSaid == 1 {
			last = 0
		} else {
			last = numbers[last].next
		}

		if _, ok := numbers[last]; !ok {
			numbers[last] = &info{}
		}

		numbers[last].next = turn - numbers[last].lastTurnSaid
		numbers[last].lastTurnSaid = turn
		numbers[last].timesSaid++
		turn++
	}

	return last
}
