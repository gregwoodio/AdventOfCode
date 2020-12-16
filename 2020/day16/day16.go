package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type inclusiveRange struct {
	low, high int
}

func solvePartOne(input []string) int {
	validRanges, _, tickets := parseInput(input)

	sum := 0
	for _, ticket := range tickets {
		for _, val := range ticket {
			invalid := true

		rangeCheck:
			for _, ranges := range validRanges {
				for _, r := range ranges {
					if r.low <= val && val <= r.high {
						invalid = false
						break rangeCheck
					}
				}
			}

			if invalid {
				sum += val
			}
		}
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}

// parse input finds the values valid for each field, the numbers from your ticket,
// then the numbers from all other tickets.
func parseInput(input []string) (map[string][]inclusiveRange, []int, [][]int) {
	i := 0
	regex := regexp.MustCompile(`(?P<field>[a-z ]+): (?P<r1low>[0-9]+)-(?P<r1high>[0-9]+) or (?P<r2low>[0-9]+)-(?P<r2high>[0-9]+)`)
	validRanges := make(map[string][]inclusiveRange)
	for _, line := range input {
		i++
		if line == "" {
			break
		}

		parts := regex.FindStringSubmatch(line)

		r1low, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Errorf("could not parse %s to int", parts[2])
		}

		r1high, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Errorf("could not parse %s to int", parts[3])
		}

		r2low, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Errorf("could not parse %s to int", parts[4])
		}

		r2high, err := strconv.Atoi(parts[5])
		if err != nil {
			fmt.Errorf("could not parse %s to int", parts[5])
		}

		validRanges[parts[1]] = []inclusiveRange{
			{
				low: r1low, high: r1high,
			},
			{
				low: r2low, high: r2high,
			},
		}
	}

	i++
	yourTicket := parseLineToIntSlice(input[i])

	i += 3
	tickets := [][]int{}
	for ; i < len(input); i++ {
		tickets = append(tickets, parseLineToIntSlice(input[i]))
	}

	return validRanges, yourTicket, tickets
}

func parseLineToIntSlice(line string) []int {
	slice := []int{}
	nums := strings.Split(line, ",")
	for _, num := range nums {
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Errorf("could not parse int '%s'", num)
		}

		slice = append(slice, val)
	}
	return slice
}
