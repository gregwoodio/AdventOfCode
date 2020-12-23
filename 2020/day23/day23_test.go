package day23

import (
	"fmt"
	"testing"
)

func Test2020Day23_SolvePartOne(t *testing.T) {
	input := "389125467"

	expected := "67384529"
	actual := solvePartOne(input)
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func Test2020Day23_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne("853192647")
	fmt.Println(solution)
}

func Test2020Day23_SolvePartTwo(t *testing.T) {
	input := "389125467"

	var expected int64 = 149245887792
	actual := solvePartTwo(input)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day23_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo("853192647")
	fmt.Println(solution)
}
