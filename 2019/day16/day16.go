package day16

import (
	"fmt"
	"strconv"
)

func solvePartOne(input string, phases int) []int {
	nums := toIntSlice(input)
	pattern := []int{0, 1, 0, -1}

	for phase := 0; phase < phases; phase++ {
		newNums := []int{}

		for j := 1; j <= len(nums); j++ {
			patternPos := 0
			pos := 0

			innerNums := []int{}

			for _, num := range nums {

				patternPos++
				if patternPos%j == 0 {
					pos++
					if pos%len(pattern) == 0 {
						pos = 0
					}
				}

				innerNums = append(innerNums, ((num * pattern[pos]) % 10))
			}

			sum := 0
			for _, val := range innerNums {
				sum += val
			}

			newNums = append(newNums, abs(sum%10))
		}

		nums = newNums
	}

	return nums[:8]
}

func toIntSlice(input string) []int {
	nums := []int{}
	for _, val := range input {
		num, _ := strconv.Atoi(fmt.Sprintf("%c", val))
		nums = append(nums, num)
	}
	return nums
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
