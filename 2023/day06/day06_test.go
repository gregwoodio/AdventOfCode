package day06

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2023Day06_SolvePartOne(t *testing.T) {
	inputs := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	expected := 288
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day06_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day06_input.txt"))
	fmt.Println(solution)
}

func Test2023Day06_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"Time:      7 15 30",
		"Distance:  9 40 200",
	}

	expected := 71503
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day06_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day06_input.txt"))
	fmt.Println(solution)
}
