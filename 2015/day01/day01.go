package day01

func solvePartOne(input string) int {
	return solve(input, false)
}

func solvePartTwo(input string) int {
	return solve(input, true)
}

func solve(input string, isPartTwo bool) int {
	floor := 0

	for i, ch := range input {
		switch ch {
		case '(':
			floor++
		case ')':
			floor--
		default:
			continue
		}

		if isPartTwo && floor == -1 {
			return i + 1
		}
	}

	if isPartTwo {
		// never went below 0
		return 0
	}

	return floor
}