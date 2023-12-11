package day10

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2023Day10_SolvePartOne(t *testing.T) {
	inputs := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			expected: 4,
		},
		{
			input: []string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			expected: 8,
		},
	}

	for _, td := range inputs {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2023Day10_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day10_input.txt"))
	fmt.Println(solution)
}

func Test2023Day10_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}

	expected := 10
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day10_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day10_input.txt"))
	fmt.Println(solution)
}
