package day01

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day01_SolvePartOne(t *testing.T) {
	inputs := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	actual := solvePartOne(inputs)
	expected := 514579

	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day01_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadIntsFromFile("./day01_input.txt"))
	fmt.Println(solution)
}

func Test2020Day01_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadIntsFromFile("./day01_input.txt"))
	fmt.Println(solution)
}
