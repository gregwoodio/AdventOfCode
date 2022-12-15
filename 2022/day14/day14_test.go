package day14

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day14_SolvePartOne(t *testing.T) {
	inputs := []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}

	expected := 24
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day14_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day14_input.txt"))
	fmt.Println(solution)
}

func Test2022Day14_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}

	expected := 93
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day14_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day14_input.txt"))
	fmt.Println(solution)
}
