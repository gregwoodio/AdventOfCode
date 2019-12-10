package day07

import (
	"fmt"
	"sync"

	"github.com/gregwoodio/aoc2019shared"
	"github.com/gregwoodio/aocutil"
)

func solve(input string, isPartTwo bool) int64 {
	var initialSettings []int

	if !isPartTwo {
		initialSettings = []int{0, 1, 2, 3, 4}
	} else {
		initialSettings = []int{5, 6, 7, 8, 9}
	}

	permutations := aocutil.Permutations(initialSettings)
	var max int64
	maxOrder := make([]int, 5)

	for _, perm := range permutations {
		perm64 := toInt64Slice(perm)

		curr := processPermutation(input, perm64, isPartTwo)

		if curr > max {
			max = curr
			copy(maxOrder, perm)
		}
	}

	fmt.Printf("Max value: %d\n[", max)
	for _, val := range maxOrder {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("]\n")

	return max
}

func processPermutation(instructions string, input []int64, isPartTwo bool) int64 {
	amps := []*aoc2019shared.IntCodeInterpreter{
		aoc2019shared.NewIntCodeInterpreter("Amp A", instructions),
		aoc2019shared.NewIntCodeInterpreter("Amp B", instructions),
		aoc2019shared.NewIntCodeInterpreter("Amp C", instructions),
		aoc2019shared.NewIntCodeInterpreter("Amp D", instructions),
		aoc2019shared.NewIntCodeInterpreter("Amp E", instructions),
	}

	// Connect amplifiers
	amps[1].Input = amps[0].Output
	amps[2].Input = amps[1].Output
	amps[3].Input = amps[2].Output
	amps[4].Input = amps[3].Output

	if isPartTwo {
		amps[0].Input = amps[4].Output
	}

	var wg sync.WaitGroup

	for i, amp := range amps {
		amp.Input <- input[i]
		if i == 0 {
			amp.Input <- 0
		}

		if i == 4 {
			wg.Add(1)
			go amp.Process(&wg)
		} else {
			go amp.Process(nil)
		}
	}

	wg.Wait()
	return <-amps[4].Output
}

func toInt64Slice(slice []int) []int64 {
	newSlice := []int64{}

	for _, val := range slice {
		newSlice = append(newSlice, int64(val))
	}

	return newSlice
}
