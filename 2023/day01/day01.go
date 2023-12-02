package day01

import (
	"strconv"
	"unicode"
)

func solvePartOne(input []string) int {
	sum := 0

	for _, line := range input {
		var first, last int
		var hasFirst bool

		for _, ch := range line {
			if unicode.IsDigit(ch) {
				if !hasFirst {
					first, _ = strconv.Atoi(string(ch))
					hasFirst = true
				}

				last, _ = strconv.Atoi(string(ch))
			}
		}

		sum += 10*first + last
	}

	return sum
}

func solvePartTwo(input []string) int {
	sum := 0

	wordValues := []struct {
		word  string
		value int
	}{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}

	for _, line := range input {
		var first, last int
		var hasFirst bool

		for i := 0; i < len(line); i++ {
			ch := rune(line[i])

			if unicode.IsDigit(ch) {
				if !hasFirst {
					first, _ = strconv.Atoi(string(ch))
					hasFirst = true
				}

				last, _ = strconv.Atoi(string(ch))
			} else {
				for _, wordValue := range wordValues {
					if checkContains(wordValue.word, line, i) {
						if !hasFirst {
							first = wordValue.value
							hasFirst = true
						}

						last = wordValue.value
						break
					}
				}

			}
		}

		sum += 10*first + last
	}

	return sum
}

func checkContains(word string, line string, i int) bool {
	if i+len(word) > len(line) {
		return false
	}

	for j, ch := range word {
		if rune(line[i+j]) != ch {
			return false
		}
	}

	return true
}
