package day12

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day12_SolvePartOne(t *testing.T) {
	inputs := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	expected := 25
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day12_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day12_input.txt"))
	fmt.Println(solution)
}

func Test2020Day12_SolvePartTwo(t *testing.T) {
	testData := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			expected: 286,
		},
		{
			inputs: []string{
				"F1",
				"L90",
				"F1",
				"L90",
				"F1",
				"L90",
				"F1",
				"F1",
				"R90",
				"F1",
				"R90",
				"F1",
				"R90",
				"F1",
			},
			expected: 0,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo(td.inputs)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day12_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day12_input.txt"))
	fmt.Println(solution)
}
