package day07

import (
	"fmt"

	"github.com/gregwoodio/aoc2019shared"
	"github.com/gregwoodio/aocutil"
)

func solvePartOne(input string) int {

	permutations := aocutil.Permutations([]int{0, 1, 2, 3, 4})
	max := 0
	maxOrder := make([]int, 5)

	for _, perm := range permutations {
		curr := processPermutation(input, perm)

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

func processPermutation(instructions string, input []int) int {
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

	for i, amp := range amps {
		amp.Input <- input[i]
		if i == 0 {
			amp.Input <- 0
		}

		go amp.Process()
	}

	allDone := make([]bool, 5)
	for {
		select {
		case val := <-amps[4].Output:
			return val
		case allDone[0] = <-amps[0].Done:
			if allDone[0] && allDone[1] && allDone[2] && allDone[3] && allDone[4] {
				return *amps[4].LastOutput
			}
		case allDone[1] = <-amps[1].Done:
			if allDone[0] && allDone[1] && allDone[2] && allDone[3] && allDone[4] {
				return *amps[4].LastOutput
			}
		case allDone[2] = <-amps[2].Done:
			if allDone[0] && allDone[1] && allDone[2] && allDone[3] && allDone[4] {
				return *amps[4].LastOutput
			}
		case allDone[3] = <-amps[3].Done:
			if allDone[0] && allDone[1] && allDone[2] && allDone[3] && allDone[4] {
				return *amps[4].LastOutput
			}
		case allDone[4] = <-amps[4].Done:
			if allDone[0] && allDone[1] && allDone[2] && allDone[3] && allDone[4] {
				return *amps[4].LastOutput
			}
		}
	}
}
