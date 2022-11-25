package day02

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	sum := 0

	for _, directions := range input {
		sides := strings.Split(directions, "x")
		nums := []int{}
		for _, side := range sides {
			val, err := strconv.Atoi(side)

			if err != nil {
				fmt.Printf("Could not parse %s", sides[0])
			}
			nums = append(nums, val)
		}

		side1 := nums[0] * nums[1]
		side2 := nums[1] * nums[2]
		side3 := nums[2] * nums[0]
		smallest := int(math.Min(math.Min(float64(side1), float64(side2)), float64(side3)))

		sum += 2*side1 + 2*side2 + 2*side3 + smallest
	}

	return sum
}

func solvePartTwo(input []string) int {
	sum := 0

	for _, directions := range input {
		sides := strings.Split(directions, "x")
		nums := []int{}
		for _, side := range sides {
			val, err := strconv.Atoi(side)

			if err != nil {
				fmt.Printf("Could not parse %s", sides[0])
			}
			nums = append(nums, val)
		}

		sort.Ints(nums)
		sum += 2*nums[0] + 2*nums[1] + nums[0]*nums[1]*nums[2]
	}

	return sum
}
