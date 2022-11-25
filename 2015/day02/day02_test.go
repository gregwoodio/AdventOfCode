package day02

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2015Day02_SolvePartOne(t *testing.T) {
	inputs := []string{
		"2x3x4",
		"1x1x10",
	}

	expected := 58 + 43
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2015Day02_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day02_input.txt"))
	fmt.Println(solution)
}

func Test2015Day02_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"2x3x4",
		"1x1x10",
	}

	expected := 34 + 14
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2015Day02_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day02_input.txt"))
	fmt.Println(solution)
}
