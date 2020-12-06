package day06

func solvePartOne(input []string) int {
	return solve(input, false)
}

func solvePartTwo(input []string) int {
	return solve(input, true)
}

func solve(input []string, isPartTwo bool) int {
	sum := 0
	group := make(map[rune]int)
	groupSize := 0

	for _, line := range input {
		if line == "" {
			if !isPartTwo {
				sum += len(group)
			} else {
				for _, val := range group {
					if val == groupSize {
						sum++
					}
				}
			}
			group = make(map[rune]int)
			groupSize = 0
			continue
		}

		for _, ch := range line {
			group[ch]++
		}
		groupSize++
	}

	if len(group) > 0 {
		if !isPartTwo {
			sum += len(group)
		} else {
			for _, val := range group {
				if val == groupSize {
					sum++
				}
			}
		}
	}

	return sum
}
