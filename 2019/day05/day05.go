package day05

import (
	"io"

	"github.com/gregwoodio/aoc2019shared"
)

func solve(input string, reader io.Reader, writer io.Writer) int {
	ici := aoc2019shared.NewIntCodeInterpreter(input)
	return ici.Process(reader, writer)
}
