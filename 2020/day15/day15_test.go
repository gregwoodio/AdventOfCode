package day15

import (
	"fmt"
	"testing"
)

func Test2020Day15_SolvePartOne(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{
			[]int{0, 3, 6},
			436,
		},
		{
			[]int{1, 3, 2},
			1,
		},
		{
			[]int{2, 1, 3},
			10,
		},
		{
			[]int{1, 2, 3},
			27,
		},
		{
			[]int{2, 3, 1},
			78,
		},
		{
			[]int{3, 2, 1},
			438,
		},
		{
			[]int{3, 1, 2},
			1836,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day15_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	input := []int{7, 14, 0, 17, 11, 1, 2}
	solution := solvePartOne(input)
	fmt.Println(solution)
}

func Test2020Day15_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	input := []int{7, 14, 0, 17, 11, 1, 2}
	solution := solvePartTwo(input)
	fmt.Println(solution)
}
