package day13

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2020Day13_SolvePartOne(t *testing.T) {
	inputs := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	expected := 295
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day13_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day13_input.txt"))
	fmt.Println(solution)
}

func Test2020Day13_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	var expected int = 1068781
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day13_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day13_input.txt"))
	fmt.Println(solution)
}
