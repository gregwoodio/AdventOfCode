package day06

func solvePartOne(input string) int {
	return findDistinctCharacters(input, 4)
}

func solvePartTwo(input string) int {
	return findDistinctCharacters(input, 14)
}

func findDistinctCharacters(input string, n int) int {
	i := n
	for {
		buffer := input[i-n : i]

		values := map[rune]bool{}
		for _, ch := range buffer {
			values[ch] = true
		}

		if len(values) == n {
			break
		}
		i++
	}

	return i
}
