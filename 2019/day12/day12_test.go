package day12

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input    []string
	steps    int
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testDatas := []testData{
		testData{
			input: []string{
				"<x=-1, y=0, z=2>",
				"<x=2, y=-10, z=-7>",
				"<x=4, y=-8, z=8>",
				"<x=3, y=5, z=-1>",
			},
			steps:    10,
			expected: 179,
		},
		testData{
			input: []string{
				"<x=-8, y=-10, z=0>",
				"<x=5, y=5, z=10>",
				"<x=2, y=-7, z=3>",
				"<x=9, y=-8, z=-3>",
			},
			steps:    100,
			expected: 1940,
		},
	}

	for _, td := range testDatas {
		actual := solvePartOne(td.input, td.steps)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day12_input.txt")
	result := solvePartOne(input, 1000)
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	testDatas := []testData{
		testData{
			input: []string{
				"<x=-1, y=0, z=2>",
				"<x=2, y=-10, z=-7>",
				"<x=4, y=-8, z=8>",
				"<x=3, y=5, z=-1>",
			},
			expected: 2772,
		},
		testData{
			input: []string{
				"<x=-8, y=-10, z=0>",
				"<x=5, y=5, z=10>",
				"<x=2, y=-7, z=3>",
				"<x=9, y=-8, z=-3>",
			},
			expected: 4686774924,
		},
	}

	for _, td := range testDatas {
		actual := solvePartTwo(td.input)
		if actual != uint64(td.expected) {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day12_input.txt")
	result := solvePartTwo(input)
	fmt.Println(result)
}
