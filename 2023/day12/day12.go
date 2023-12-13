package day12

import (
	"strconv"
	"strings"
)

type springs struct {
	line   string
	values []int
}

func parse(input string) springs {
	i := strings.Index(input, " ")
	line := input[:i]
	parts := strings.Split(input[i+1:], ",")

	values := []int{}
	for _, part := range parts {
		value, _ := strconv.Atoi(part)
		values = append(values, value)
	}

	return springs{
		line:   line,
		values: values,
	}
}

func solvePartOne(input []string) int {
	sum := 0

	for _, line := range input {
		currentSum := 0
		s := parse(line)

		// find all possible combinations
		combinations := []string{s.line}
		for i := range s.line {
			if s.line[i] == '?' {
				newCombinations := []string{}
				for _, c := range combinations {
					newCombinations = append(newCombinations, c[:i]+"#"+c[i+1:])
					newCombinations = append(newCombinations, c[:i]+"."+c[i+1:])
				}
				combinations = newCombinations
			}
		}
		// fmt.Printf("Found %d combinations for %s\n", len(combinations), s.line)

		// count the valid combinations
		for _, combination := range combinations {
			valid := true
			index, countIndex := 0, 0

			// quick check: the number of # should match the sum of the values
			valueSum := 0
			for _, value := range s.values {
				valueSum += value
			}
			replaced := strings.ReplaceAll(combination, ".", "")
			if len(replaced) != valueSum {
				valid = false
				index = len(combination)
			}

		Outer:
			for index < len(combination) {
				// skip leading .
				for _, ch := range combination[index:] {
					if ch == '.' {
						index++
						continue
					} else {
						break
					}
				}

				// count the #
				currCount := 0
				for _, ch := range combination[index:] {
					if ch != '#' {
						valid = false
						break Outer
					} else {
						index++
						currCount++

						if countIndex < len(s.values) && currCount >= s.values[countIndex] {
							break
						}
					}
				}

				// each group should end with a . or the end of the string
				if index < len(combination) && combination[index] != '.' {
					// fmt.Printf("Invalid combination: %s\n", combination)
					valid = false
					break Outer
				} else {
					index++
					countIndex++
				}

			}

			if valid {
				// fmt.Printf("Valid combination: %s\n", combination)
				currentSum++
			}
		}

		// fmt.Printf("Found %d valid combinations for %s\n\n", currentSum, s.line)

		sum += currentSum
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}
