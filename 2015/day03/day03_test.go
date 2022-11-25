package day03

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func TestSolvePartOne(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			input:    ">",
			expected: 2,
		},
		{
			input:    "^>v<",
			expected: 4,
		},
		{
			input:    "^v^v^v^v^v",
			expected: 2,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day03_input.txt")
	result := solvePartOne(input[0])
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			input:    "^v",
			expected: 3,
		},
		{
			input:    "^>v<",
			expected: 3,
		},
		{
			input:    "^v^v^v^v^v",
			expected: 11,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day03_input.txt")
	result := solvePartTwo(input[0])
	fmt.Println(result)
}
