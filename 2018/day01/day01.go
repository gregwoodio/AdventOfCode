package day01

import (
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	value := 0

	for _, cmd := range input {
		num, _ := strconv.Atoi(cmd[1:])
		if strings.HasPrefix(cmd, "+") {
			value += num
		} else {
			value -= num
		}
	}

	return value
}

func solvePartTwo(input []string) int {
	value := 0
	valueMap := make(map[int]bool)
	valueMap[0] = true

	for {
		for _, cmd := range input {
			num, _ := strconv.Atoi(cmd[1:])
			if strings.HasPrefix(cmd, "+") {
				value += num
			} else {
				value -= num
			}

			if valueMap[value] == true {
				return value
			}
			valueMap[value] = true
		}
	}
}
