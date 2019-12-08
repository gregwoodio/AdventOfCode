package day05

import (
	"fmt"

	"github.com/gregwoodio/aoc2019shared"
)

func solve(instructions string, input int) int {
	ici := aoc2019shared.NewIntCodeInterpreter(instructions)
	go ici.Process()

	ici.Input <- input

	for {
		select {
		case o := <-ici.Output:
			fmt.Println(o)

		case <-ici.Done:
			return *ici.LastOutput
		}
	}
}
