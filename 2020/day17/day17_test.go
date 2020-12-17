package day17

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2020Day17_SolvePartOne(t *testing.T) {
	inputs := []string{
		".#.",
		"..#",
		"###",
	}

	expected := 112
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day17_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day17_input.txt"))
	fmt.Println(solution)
}

func Test2020Day17_SolvePartTwo(t *testing.T) {
	inputs := []string{
		".#.",
		"..#",
		"###",
	}

	expected := 848
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day17_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day17_input.txt"))
	fmt.Println(solution)
}
