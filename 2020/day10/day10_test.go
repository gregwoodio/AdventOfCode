package day10

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day10_SolvePartOne(t *testing.T) {
	inputs := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	expected := 35
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day10_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadIntsFromFile("./day10_input.txt"))
	fmt.Println(solution)
}

func Test2020Day10_SolvePartTwo(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			expected: 8,
		},
		{
			[]int{28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			19208,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day10_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadIntsFromFile("./day10_input.txt"))
	fmt.Println(solution)
}
