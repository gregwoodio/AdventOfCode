package day03

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

type testData struct {
	input    []string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testDatas := []testData{
		testData{
			input: []string{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			expected: 6,
		},
		testData{
			input: []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			expected: 159,
		},
		testData{
			input: []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			expected: 135,
		},
	}

	for _, td := range testDatas {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day03_input.txt")
	result := solvePartOne(input)

	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	testDatas := []testData{
		testData{
			input: []string{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			expected: 30,
		},
		testData{
			input: []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			expected: 610,
		},
		testData{
			input: []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			expected: 410,
		},
	}

	for _, td := range testDatas {
		actual := solvePartTwo(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day03_input.txt")
	result := solvePartTwo(input)

	fmt.Println(result)
}
