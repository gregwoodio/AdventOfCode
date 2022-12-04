package day04

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	sum := 0
	for _, line := range input {

		// compare pairs of assignments. Compare the shorter list of
		// assignments against the longer to see that all are included
		pairs := createPairsOfAssignments(line)
		match := true
		comparer := pairs[0]
		comparee := pairs[1]
		if len(comparer) > len(pairs[1]) {
			comparer = pairs[1]
			comparee = pairs[0]
		}

		for key, _ := range comparer {
			if _, ok := comparee[key]; !ok {
				match = false
				break
			}
		}

		if match {
			sum++
		}
	}

	return sum
}

func solvePartTwo(input []string) int {
	sum := 0

	for _, line := range input {
		pairs := createPairsOfAssignments(line)

		// Compare pairs of assignments to find any overlap
		match := false
		for key, _ := range pairs[0] {
			if _, ok := pairs[1][key]; ok {
				match = true
				break
			}
		}

		if match {
			sum++
		}

	}

	return sum
}

func createPairsOfAssignments(line string) []map[int]bool {
	pairs := []map[int]bool{}

	// create map of ids assigned to each elf
	sections := strings.Split(line, ",")
	for _, section := range sections {
		ids := strings.Split(section, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Printf("Could not parse '%s'\n", ids[0])
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Printf("Could not parse '%s'\n", ids[1])
		}

		assigned := map[int]bool{}
		for i := start; i <= end; i++ {
			assigned[i] = true
		}

		pairs = append(pairs, assigned)
	}

	return pairs
}
