package day21

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day21_SolvePartOne(t *testing.T) {
	inputs := []string{
		"root: pppw + sjmn",
		"dbpl: 5",
		"cczh: sllz + lgvd",
		"zczc: 2",
		"ptdq: humn - dvpt",
		"dvpt: 3",
		"lfqf: 4",
		"humn: 5",
		"ljgn: 2",
		"sjmn: drzm * dbpl",
		"sllz: 4",
		"pppw: cczh / lfqf",
		"lgvd: ljgn * ptdq",
		"drzm: hmdt - zczc",
		"hmdt: 32",
	}

	expected := 152
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day21_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day21_input.txt"))
	fmt.Println(solution)
}

func Test2022Day21_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day21_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day21_input.txt"))
	fmt.Println(solution)
}
