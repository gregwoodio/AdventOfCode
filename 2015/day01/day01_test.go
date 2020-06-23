package day01

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestSolvePartOne(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, test := range tests {
		actual := solvePartOne(test.input)

		if actual != test.expected {
			t.Errorf("Expected %d but was %d", test.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day01_input.txt")
	result := solvePartOne(input[0])
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		actual := solvePartTwo(test.input)

		if actual != test.expected {
			t.Errorf("Expected %d but was %d", test.expected, actual)
		}
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day01_input.txt")
	result := solvePartTwo(input[0])
	fmt.Println(result)
}