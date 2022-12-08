package day08

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day08_SolvePartOne(t *testing.T) {
	inputs := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expected := 21
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day08_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day08_input.txt"))
	fmt.Println(solution)
}

func Test2022Day08_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expected := 8
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day08_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day08_input.txt"))
	fmt.Println(solution)
}
