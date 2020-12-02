package day02

import (
	"fmt"
	"regexp"
	"strconv"
)

func solvePartOne(inputs []string) int {
	return parse(inputs, false)
}

func solvePartTwo(input []string) int {
	return parse(input, true)
}

func parse(inputs []string, isPartTwo bool) int {
	validPasswords := 0

	regex := regexp.MustCompile(`(?P<low>\d+)-(?P<high>\d+) (?P<ch>[a-z]): (?P<pass>[a-z]+)`)
	for _, input := range inputs {
		parts := regex.FindStringSubmatch(input)

		if len(parts) != 5 {
			fmt.Errorf("did not parse correctly")
		}

		low, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Errorf("could not parse %s to int", parts[1])
		}

		high, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Errorf("could not parse %s to int", parts[2])
		}

		ch := rune(parts[3][0])

		count := 0
		for i, letter := range parts[4] {
			if !isPartTwo {
				if letter == ch {
					count++
				}
			} else {
				if (i+1 == low && letter == ch) || (i+1 == high && letter == ch) {
					count++
				}
			}
		}

		if !isPartTwo {
			if count >= low && count <= high {
				validPasswords++
			}
		} else {
			if count == 1 {
				validPasswords++
			}
		}
	}

	return validPasswords
}
