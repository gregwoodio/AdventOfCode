package day05

import (
	"sync"

	"github.com/gregwoodio/aoc2019shared"
)

func solve(instructions string, input int) int {
	ici := aoc2019shared.NewIntCodeInterpreter("aoc", instructions)
	var wg sync.WaitGroup

	go func(out chan int) {
		for {
			// just grab and ignore all output from the interpreter
			<-out
		}
	}(ici.Output)

	wg.Add(1)
	ici.Input <- input

	go ici.Process(&wg)

	wg.Wait()

	return *ici.LastOutput
}
