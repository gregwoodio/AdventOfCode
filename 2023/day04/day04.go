package day04

import "strings"

func solvePartOne(input []string) int {
	sum := 0

	for _, line := range input {
		colon := strings.Index(line, ":")
		nums := strings.Split(line[colon+1:], "|")
		winningNums := map[string]bool{}
		score := 0
		for _, num := range strings.Split(nums[0], " ") {
			if len(num) != 0 {
				winningNums[num] = true
			}
		}
		for _, num := range strings.Split(nums[1], " ") {
			if _, ok := winningNums[num]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		sum += score

	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}
