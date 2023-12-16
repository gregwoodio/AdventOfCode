package day16

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2023Day16_SolvePartOne(t *testing.T) {
	inputs := []string{
		`.|...\....`,
		`|.-.\.....`,
		`.....|-...`,
		`........|.`,
		`..........`,
		`.........\`,
		`..../.\\..`,
		`.-.-/..|..`,
		`.|....-|.\`,
		`..//.|....`,
	}

	expected := 46
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day16_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day16_input.txt"))
	fmt.Println(solution)
}

func Test2023Day16_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day16_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day16_input.txt"))
	fmt.Println(solution)
}
