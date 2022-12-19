package day18

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day18_SolvePartOneShort(t *testing.T) {
	inputs := []string{
		"1,1,1",
		"2,1,1",
	}

	expected := 10
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day18_SolvePartOne(t *testing.T) {
	inputs := []string{
		"2,2,2",
		"1,2,2",
		"3,2,2",
		"2,1,2",
		"2,3,2",
		"2,2,1",
		"2,2,3",
		"2,2,4",
		"2,2,6",
		"1,2,5",
		"3,2,5",
		"2,1,5",
		"2,3,5",
	}

	expected := 64
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day18_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day18_input.txt"))
	fmt.Println(solution)
}

func Test2022Day18_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"2,2,2",
		"1,2,2",
		"3,2,2",
		"2,1,2",
		"2,3,2",
		"2,2,1",
		"2,2,3",
		"2,2,4",
		"2,2,6",
		"1,2,5",
		"3,2,5",
		"2,1,5",
		"2,3,5",
	}

	expected := 58
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day18_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day18_input.txt"))
	fmt.Println(solution)
}
