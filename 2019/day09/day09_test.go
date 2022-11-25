package day09

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

type testData struct {
	input string
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	inputs := aocutil.ReadStringsFromFile("./day09_input.txt")
	fmt.Println(solve(inputs[0], 1))
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	inputs := aocutil.ReadStringsFromFile("./day09_input.txt")
	fmt.Println(solve(inputs[0], 2))
}
