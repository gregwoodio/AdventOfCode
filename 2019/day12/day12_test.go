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
