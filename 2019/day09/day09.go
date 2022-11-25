package day09

import (
	"fmt"
	"sync"

	"github.com/gregwoodio/adventofcode/m/aoc2019shared"
)

func solve(input string, inst int) int64 {
	ici := aoc2019shared.NewIntCodeInterpreter("day09", input)

	var wg sync.WaitGroup

	wg.Add(1)
	go ici.Process(&wg)
	ici.Input <- int64(inst)

	out := <-ici.Output

	wg.Wait()
	fmt.Println(out)
	return out
}
