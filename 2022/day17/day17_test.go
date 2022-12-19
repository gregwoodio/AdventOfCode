package day17

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day17_SolvePartOne(t *testing.T) {
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	expected := 3068
	actual := solvePartOne(input)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day17_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day17_input.txt")[0])
	fmt.Println(solution)
}

func Test2022Day17_SolvePartTwo(t *testing.T) {
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	expected := 3068
	actual := solvePartTwo(input)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day17_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day17_input.txt")[0])
	fmt.Println(solution)
}
