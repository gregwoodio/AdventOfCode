package day21

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day21_SolvePartOne(t *testing.T) {
	inputs := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	expected := 5
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day21_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day21_input.txt"))
	fmt.Println(solution)
}

func Test2020Day21_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	expected := "mxmxvkd,sqjhc,fvjkl"
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func Test2020Day21_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day21_input.txt"))
	fmt.Println(solution)
}
