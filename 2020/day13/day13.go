package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	if len(input) != 2 {
		fmt.Errorf("invalid input")
	}

	timestamp, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Errorf("can't parse string to int: '%s'", input[0])
	}

	busIDString := strings.Split(input[1], ",")
	buses := []int{}
	for _, busID := range busIDString {
		bus, err := strconv.Atoi(busID)
		if err == nil {
			buses = append(buses, bus)
		}
		// x
	}

	for depart := timestamp; ; depart++ {
		for _, bus := range buses {
			if depart%bus == 0 {
				return (depart - timestamp) * bus
			}
		}
	}
}

func solvePartTwo(input []string) int {
	busIDString := strings.Split(input[1], ",")

	// buses will now use index as the offset
	buses := []int{}
	for _, busID := range busIDString {
		bus, err := strconv.Atoi(busID)
		if err == nil {
			buses = append(buses, bus)
		} else {
			// x
			buses = append(buses, -1)
		}
	}

	min := 0
	prod := 1
	for i, bus := range buses {
		if bus == -1 {
			continue
		}

		for (min+i)%bus != 0 {
			min += prod
		}
		prod *= bus
	}

	return min
}
