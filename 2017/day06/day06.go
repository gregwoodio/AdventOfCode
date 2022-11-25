package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func main() {
	input := aocutil.ReadAndSplitInts("input.txt", "\t")
	fmt.Println(partOne(input[0]))

	input = aocutil.ReadAndSplitInts("input.txt", "\t")
	fmt.Println(partTwo(input[0]))
}

func partOne(input []int) int {
	return solve(input, false)
}

func partTwo(input []int) int {
	return solve(input, true)
}

func solve(input []int, isPartTwo bool) int {
	cycles := 0

	knownPatterns := make(map[string]int)

	for {
		id := toStringIdentifier(input)
		if knownPatterns[id] > 0 {
			if isPartTwo {
				return cycles - knownPatterns[id]
			}
			return cycles
		}
		knownPatterns[id] = cycles

		i := findHighestIndex(input)
		v := input[i]
		input[i] = 0

		for ; v > 0; v-- {
			i++
			if i >= len(input) {
				i = 0
			}
			input[i]++
		}

		cycles++
	}
}

func findHighestIndex(input []int) int {
	highestIndex, highestValue := 0, 0
	for index, value := range input {
		if value > highestValue {
			highestValue = value
			highestIndex = index
		}
	}
	return highestIndex
}

func toStringIdentifier(input []int) string {
	var numberStrings []string
	for _, num := range input {
		numberStrings = append(numberStrings, strconv.Itoa(num))
	}
	return strings.Join(numberStrings, ",")
}
