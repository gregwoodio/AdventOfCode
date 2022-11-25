package day18

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day18_SolvePartOne(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			"1 + 2 * 3 + 4 * 5 + 6",
			71,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			26,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			437,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			12240,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			13632,
		},
	}

	for _, td := range testData {
		actual := solvePartOne([]string{td.input})
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day18_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day18_input.txt"))
	fmt.Println(solution)
}

func Test2020Day18_SolvePartTwo(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			"2 * 3 + 4",
			14,
		},
		{
			"2 + 3 * 4",
			20,
		},
		{
			"1 + 2 * 3 + 4 * 5 + 6",
			231,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			46,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			1445,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			669060,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			23340,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo([]string{td.input})
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day18_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day18_input.txt"))
	fmt.Println(solution)
}
