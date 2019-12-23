package day05

import (
	"sync"

	"github.com/gregwoodio/aoc2019shared"
)

func solve(instructions string, input int64) int64 {
	ici := aoc2019shared.NewIntCodeInterpreter("aoc", instructions)
	var wg sync.WaitGroup

	output := make(chan int64)
	defer close(output)
	go receiver(ici.Output, output)

	wg.Add(1)
	go ici.Process(&wg)
	ici.Input <- input
	wg.Wait()

	return <-output
}

func receiver(in, out chan int64) {
	for {
		val := <-in
		if val != 0 {
			out <- val
			return
		}
	}
}
