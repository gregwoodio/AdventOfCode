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
			} else if ch == 'o' {
				if i+2 < len(line) && line[i+1] == 'n' && line[i+2] == 'e' {
					if !hasFirst {
						first = 1
						hasFirst = true
					}
					last = 1
				}
			} else if ch == 't' {
				if i+2 < len(line) && line[i+1] == 'w' && line[i+2] == 'o' {
					if !hasFirst {
						first = 2
						hasFirst = true
					}
					last = 2
				}

				if i+4 < len(line) && line[i+1] == 'h' && line[i+2] == 'r' && line[i+3] == 'e' && line[i+4] == 'e' {
					if !hasFirst {
						first = 3
						hasFirst = true
					}
					last = 3
				}

			} else if ch == 'f' {
				if i+3 < len(line) && line[i+1] == 'o' && line[i+2] == 'u' && line[i+3] == 'r' {
					if !hasFirst {
						first = 4
						hasFirst = true
					}
					last = 4
				}

				if i+3 < len(line) && line[i+1] == 'i' && line[i+2] == 'v' && line[i+3] == 'e' {
					if !hasFirst {
						first = 5
						hasFirst = true
					}
					last = 5
				}

			} else if ch == 's' {
				if i+2 < len(line) && line[i+1] == 'i' && line[i+2] == 'x' {
					if !hasFirst {
						first = 6
						hasFirst = true
					}
					last = 6
				}

				if i+4 < len(line) && line[i+1] == 'e' && line[i+2] == 'v' && line[i+3] == 'e' && line[i+4] == 'n' {
					if !hasFirst {
						first = 7
						hasFirst = true
					}
					last = 7
				}
			} else if ch == 'e' {
				if i+4 < len(line) && line[i+1] == 'i' && line[i+2] == 'g' && line[i+3] == 'h' && line[i+4] == 't' {
					if !hasFirst {
						first = 8
						hasFirst = true
					}
					last = 8
				}

			} else if ch == 'n' {
				if i+3 < len(line) && line[i+1] == 'i' && line[i+2] == 'n' && line[i+3] == 'e' {
					if !hasFirst {
						first = 9
						hasFirst = true
					}
					last = 9
				}

			}
		}

		sum += 10*first + last
	}

	return sum
}
