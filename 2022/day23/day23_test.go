package day23

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day23_SolvePartOneShort(t *testing.T) {
	inputs := []string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}

	expected := 25
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day23_SolvePartOne(t *testing.T) {
	inputs := []string{
		"....#..",
		"..###.#",
		"#...#.#",
		".#...##",
		"#.###..",
		"##.#.##",
		".#..#..",
	}

	expected := 110
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day23_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day23_input.txt"))
	fmt.Println(solution)
}

func Test2022Day23_SolvePartTwoShort(t *testing.T) {
	inputs := []string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}

	expected := 4
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day23_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"....#..",
		"..###.#",
		"#...#.#",
		".#...##",
		"#.###..",
		"##.#.##",
		".#..#..",
	}

	expected := 20
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day23_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day23_input.txt"))
	fmt.Println(solution)
}
