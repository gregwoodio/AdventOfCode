package day13

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day13_SolvePartOne(t *testing.T) {
	inputs := []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}

	expected := 13
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day13_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day13_input.txt"))
	fmt.Println(solution)
}

func Test2022Day13_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}

	expected := 140
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day13_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day13_input.txt"))
	fmt.Println(solution)
}
