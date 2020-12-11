package day11

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2020Day11_SolvePartOne(t *testing.T) {
	inputs := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	expected := 37
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day11_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day11_input.txt"))
	fmt.Println(solution)
}

func Test2020Day11_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day11_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day11_input.txt"))
	fmt.Println(solution)
}
