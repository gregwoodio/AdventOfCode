package day05

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day05_SolvePartOne(t *testing.T) {
	inputs := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expected := "CMZ"
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func Test2022Day05_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day05_input.txt"))
	fmt.Println(solution)
}

func Test2022Day05_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expected := "MCD"
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func Test2022Day05_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day05_input.txt"))
	fmt.Println(solution)
}
