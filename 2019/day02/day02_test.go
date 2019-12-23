package day02

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day02_input.txt")
	solution := solvePartOne(input[0])

	fmt.Println(solution)
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day02_input.txt")
	solution := solvePartTwo(input[0])

	fmt.Println(solution)
}
