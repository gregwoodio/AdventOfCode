package day06

import (
	"math"
	"strings"
)

func solvePartOne(input []string) string {
	return solve(input, false)
}

func solvePartTwo(input []string) string {
	return solve(input, true)
}

func solve(input []string, isPartTwo bool) string {
	if len(input) == 0 {
		return ""
	}

	pos := make([]map[rune]int, len(input[0]))
	for i := range pos {
		pos[i] = make(map[rune]int)
	}

	for _, str := range input {
		for i, ch := range str {
			pos[i][ch]++
		}
	}

	sb := strings.Builder{}

	for _, p := range pos {
		value := 0
		if isPartTwo {
			value = math.MaxInt32
		}

		var selectedRune rune

		for r, val := range p {
			if isPartTwo {
				if val < value {
					selectedRune = r
					value = val
				}
			} else {
				if val > value {
					selectedRune = r
					value = val
				}
			}
		}

		sb.Write([]byte{byte(selectedRune)})
	}

	return sb.String()
}
