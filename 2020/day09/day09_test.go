package day09

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day09_SolvePartOne(t *testing.T) {
	inputs := []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}

	var expected int64
	expected = 127
	actual := solvePartOne(5, inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day09_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(25, aocutil.ReadStringsFromFile("./day09_input.txt"))
	fmt.Println(solution)
}

func Test2020Day09_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}

	var expected int64
	expected = 62

	actual := solvePartTwo(127, inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day09_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	lines := aocutil.ReadStringsFromFile("./day09_input.txt")
	solution := solvePartTwo(solvePartOne(25, lines), lines)
	fmt.Println(solution)
}
