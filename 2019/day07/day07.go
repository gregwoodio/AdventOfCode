package day07

import (
	"fmt"

	"github.com/gregwoodio/aoc2019shared"
	"github.com/gregwoodio/aocutil"
)

func solvePartOne(input []string) int {

	permutations := aocutil.Permutations([]int{0, 1, 2, 3, 4})
	max := 0
	maxOrder := make([]int, 5)

	for _, perm := range permutations {

		amps := []*aoc2019shared.IntCodeInterpreter{
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
		}

		// Connect amplifiers
		amps[1].Input = amps[0].Output
		amps[2].Input = amps[1].Output
		amps[3].Input = amps[2].Output
		amps[4].Input = amps[3].Output

		// curr := 0

		for i, amp := range amps {
			go amp.Process()

			amp.Input <- perm[i]
			if i == 0 {
				amp.Input <- 0
			}
		}

		<-amps[4].Done
		curr := amps[4].LastOutput
		if curr > max {
			max = curr
			copy(maxOrder, perm)
		}

		// for _, amp := range perm {

		// 	ici := aoc2019shared.NewIntCodeInterpreter(input[0])
		// 	var in, out bytes.Buffer
		// 	formattedInput := fmt.Sprintf("%d\n%d\n", amp, curr)
		// 	in.Write([]byte(formattedInput))

		// 	ici.Process(&in, &out)

		// 	strOut := string(out.Bytes())
		// 	curr, _ = strconv.Atoi(strOut[:strings.Index(strOut, "\n")])
		// }

		// if curr > max {
		// 	max = curr
		// 	copy(maxOrder, perm)
		// }
	}

	fmt.Printf("Max value: %d\n[", max)
	for _, val := range maxOrder {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("]\n")

	return max
}

func solvePartTwo(input []string) int {
	// permutations := aocutil.Permutations([]int{5, 6, 7, 8, 9})
	permutations := [][]int{
		[]int{5, 6, 7, 8, 9},
	}

	max := 0
	maxOrder := make([]int, 5)

	for _, perm := range permutations {
		amps := []*aoc2019shared.IntCodeInterpreter{
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
			aoc2019shared.NewIntCodeInterpreter(input[0]),
		}

		for i, amp := range amps {
			amp.Input = amps[(i+4)%5].Output
		}

		for i, amp := range amps {
			go amp.Process()

			amp.Input <- perm[i]
			if i == 0 {
				amp.Input <- 0
			}
		}

		<-amps[4].Done
		curr := amps[4].LastOutput
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
