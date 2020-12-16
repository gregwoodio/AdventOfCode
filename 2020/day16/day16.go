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
	_, sum := removeInvalidTickets(tickets, validRanges)
	return sum
}

func solvePartTwo(input []string) int {
	validRanges, myTicket, tickets := parseInput(input)
	tickets, _ = removeInvalidTickets(tickets, validRanges)

	// anything's possible!
	possibleIndex := make(map[string][]int)
	for field := range validRanges {
		indexes := make([]int, len(myTicket))
		for i := range indexes {
			indexes[i] = i
		}

		possibleIndex[field] = indexes
	}

	for field, ranges := range validRanges {
		for i := 0; i < len(myTicket); i++ {
			canBeField := true
			for _, ticket := range tickets {
				if !(ranges[0].low <= ticket[i] && ticket[i] <= ranges[0].high) && !(ranges[1].low <= ticket[i] && ticket[i] <= ranges[1].high) {
					canBeField = false
					break
				}
			}

			if !canBeField {
				possibleIndex[field] = removeValue(possibleIndex[field], i)
			}
		}
	}

	// look for fields that have only one possible index, then remove that index from
	// the other fields until all fields only have one possible index.
	for {
		// check to break
		multiplePossibilities := false
		for _, possible := range possibleIndex {
			if len(possible) > 1 {
				multiplePossibilities = true
			}
		}

		if !multiplePossibilities {
			break
		}

		// check for singles
		for i, possible := range possibleIndex {
			if len(possible) == 1 {
				for j := range possibleIndex {
					if j == i {
						continue
					}

					possibleIndex[j] = removeValue(possibleIndex[j], possible[0])
				}
			}
		}
	}

	prod := 1
	for key, val := range possibleIndex {
		if strings.HasPrefix(key, "departure") {
			prod *= myTicket[val[0]]
		}
	}

	return prod
}

func removeInvalidTickets(tickets [][]int, validRanges map[string][]inclusiveRange) ([][]int, int) {
	sum := 0
	tc := make([][]int, len(tickets))
	copy(tc, tickets)

	remove := []int{}
	for i, ticket := range tickets {
		validTicket := true
		for _, val := range ticket {
			invalidField := true

		rangeCheck:
			for _, ranges := range validRanges {
				for _, r := range ranges {
					if r.low <= val && val <= r.high {
						invalidField = false
						break rangeCheck
					}
				}
			}

			if invalidField {
				validTicket = false
				sum += val
			}
		}

		if !validTicket {
			remove = append([]int{i}, remove...)
		}
	}

	for _, r := range remove {
		tc = append(tc[:r], tc[r+1:]...)
	}

	return tc, sum
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

func removeValue(slice []int, value int) []int {
	for idx, val := range slice {
		if val == value {
			slice = append(slice[:idx], slice[idx+1:]...)
			break
		}
	}

	return slice
}
