package day03

import (
	"unicode"
)

func solvePartOne(input []string) int {
	sum := 0
	for _, bag := range input {
		compart1 := bag[0 : len(bag)/2]
		compart2 := bag[len(bag)/2:]
		c1map := toIntMap(compart1)
		c2Map := toIntMap(compart2)

		match := c1map & c2Map
		val := 0
		for ; match > 1; val++ {
			match >>= 1
		}
		sum += val
	}

	return sum
}

func solvePartTwo(input []string) int {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		elf1 := toIntMap(input[i])
		elf2 := toIntMap(input[i+1])
		elf3 := toIntMap(input[i+2])

		badge := elf1 & elf2 & elf3
		val := 0
		for ; badge > 1; val++ {
			badge >>= 1
		}
		sum += val
	}

	return sum
}

func toIntMap(bag string) int64 {
	var intMap int64
	for _, ch := range bag {
		intMap |= 1 << getLetterValue(ch)
	}

	return intMap
}

func getLetterValue(ch rune) int {
	if unicode.IsLower(ch) {
		// a to z == 1 to 26
		return int(ch) - 96
	}
	// A to Z == 27 to 52
	return int(ch) - 38
}
