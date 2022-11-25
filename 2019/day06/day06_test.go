package day06

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
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
			},
			expected: 42,
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

	input := aocutil.ReadStringsFromFile("./day06_input.txt")
	result := solvePartOne(input)
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	testDatas := []testData{
		testData{
			input: []string{
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
				"K)YOU",
				"I)SAN",
			},
			expected: 4,
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

	input := aocutil.ReadStringsFromFile("./day06_input.txt")
	result := solvePartTwo(input)
	fmt.Println(result)
}
