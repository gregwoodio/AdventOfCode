package day09

func solvePartOne(input [][]int) int {
	sum := 0

	for _, line := range input {
		diffs := [][]int{line}

		for !isAllZero(diffs[len(diffs)-1]) {
			last := len(diffs) - 1
			diff := make([]int, len(diffs[last])-1)

			for i := 0; i < len(diffs[last])-1; i++ {
				diff[i] = diffs[last][i+1] - diffs[last][i]
			}

			diffs = append(diffs, diff)
		}

		// extrapolate, add a new 0 to the end of the last diff
		diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)

		for i := len(diffs) - 2; i >= 0; i-- {
			nextLast := diffs[i+1][len(diffs[i+1])-1]
			last := diffs[i][len(diffs[i])-1]
			diffs[i] = append(diffs[i], nextLast+last)
		}

		sum += diffs[0][len(diffs[0])-1]
	}

	return sum
}

func solvePartTwo(input [][]int) int {
	return -1
}

func isAllZero(input []int) bool {
	for _, i := range input {
		if i != 0 {
			return false
		}
	}

	return true
}
