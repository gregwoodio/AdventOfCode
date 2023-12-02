package day02

import (
	"strconv"
	"strings"
)

type results struct {
	id    int
	red   int
	green int
	blue  int
}

func solvePartOne(input []string) int {
	sum := 0
	for _, line := range input {
		result := parse(line)

		if result.red <= 12 && result.green <= 13 && result.blue <= 14 {
			sum += result.id
		}
	}
	return sum
}

func solvePartTwo(input []string) int {
	return -1
}

func parse(line string) results {
	colon := strings.Index(line, ":")
	id, _ := strconv.Atoi(line[len("Game "):colon])

	subsets := strings.Split(line[colon+2:], "; ")

	red, green, blue := 0, 0, 0
	for _, subset := range subsets {
		parts := strings.Split(subset, ", ")
		for _, part := range parts {
			space := strings.Index(part, " ")
			value, _ := strconv.Atoi(part[:space])
			if part[space+1] == byte('r') {
				red = max(red, value)
			}
			if part[space+1] == byte('g') {
				green = max(green, value)
			}
			if part[space+1] == byte('b') {
				blue = max(blue, value)
			}
		}
	}

	return results{
		id:    id,
		red:   red,
		green: green,
		blue:  blue,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
