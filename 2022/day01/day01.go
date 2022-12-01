package day01

import (
	"fmt"
	"sort"
	"strconv"
)

func solvePartOne(input []string) int {
	elves := sum(input)
	max := 0
	for _, elf := range elves {
		if max < elf {
			max = elf
		}
	}

	return max
}

func solvePartTwo(input []string) int {
	elves := sum(input)
	sort.Ints(elves)

	topThree := elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	return topThree
}

func sum(input []string) []int {
	elves := []int{}
	calories := 0

	for _, value := range input {
		if value == "" {
			elves = append(elves, calories)
			calories = 0
			continue
		}

		calorie, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		calories += calorie
	}

	elves = append(elves, calories)
	return elves
}
