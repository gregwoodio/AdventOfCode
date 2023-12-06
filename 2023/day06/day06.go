package day06

import (
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	times := parseLine(input[0])
	distances := parseLine(input[1])
	totalResults := []int{}

	for i := 0; i < len(times); i++ {
		results := 0

		for wait := 1; wait < times[i]; wait++ {
			result := (times[i] - wait) * wait
			if result > distances[i] {
				results++
			}
		}

		totalResults = append(totalResults, results)
	}

	total := totalResults[0]
	for i := 1; i < len(totalResults); i++ {
		total *= totalResults[i]
	}

	return total
}

func solvePartTwo(input []string) int {
	time := parseLinePart2(input[0])
	distance := parseLinePart2(input[1])
	totalResults := []int{}

	results := 0

	for wait := 1; wait < time; wait++ {
		result := (time - wait) * wait
		if result > distance {
			results++
		}
	}

	totalResults = append(totalResults, results)

	total := totalResults[0]
	for i := 1; i < len(totalResults); i++ {
		total *= totalResults[i]
	}

	return total
}

func parseLine(line string) []int {
	splitFn := func(c rune) bool {
		return c == ' '
	}
	str := strings.FieldsFunc(line, splitFn)

	nums := []int{}
	for i, s := range str {
		if i == 0 {
			continue
		}

		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}

	return nums
}

func parseLinePart2(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	numStr := line[strings.Index(line, ":")+1:]
	num, _ := strconv.Atoi(numStr)
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
