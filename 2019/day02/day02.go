package day02

import (
	"sync"

	"github.com/gregwoodio/adventofcode/m/aoc2019shared"
)

func solvePartOne(input string) int64 {
	var wg sync.WaitGroup
	ici := aoc2019shared.NewIntCodeInterpreter("day02part01", input)
	ici.Inst[1] = 12
	ici.Inst[2] = 2

	wg.Add(1)
	zero := ici.Process(&wg)

	wg.Wait()
	return zero
}

func solvePartTwo(input string) int {
	var target int64 = 19690720

	var wg sync.WaitGroup
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			ici := aoc2019shared.NewIntCodeInterpreter("day02part02", input)
			ici.Inst[1] = int64(noun)
			ici.Inst[2] = int64(verb)

			wg.Add(1)
			result := ici.Process(&wg)

			wg.Wait()
			if result == target {
				return 100*noun + verb
			}
		}
	}

	return -1
}
