package day08

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

type testData struct {
	input  []int
	width  int
	height int
}

func TestFindLayers(t *testing.T) {
	td := testData{
		input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		width:  3,
		height: 2,
	}

	layers := findLayers(td.width, td.height, td.input)
	if len(layers) != 2 {
		t.Errorf("Expected %d layers but there were %d", 2, len(layers))
	}
}

func TestSolvePartOne(t *testing.T) {
	result := solvePartOne("123456789012", 3, 2)
	if result != 1 {
		t.Errorf("Expected %d but was %d", 1, result)
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day08_input.txt")
	result := solvePartOne(input[0], 25, 6)
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	result := solvePartTwo("0222112222120000", 2, 2)
	if result[0][0] != 0 || result[0][1] != 1 ||
		result[1][0] != 1 || result[1][1] != 0 {
		t.Errorf("Output did not match expected result")
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day08_input.txt")
	result := solvePartTwo(input[0], 25, 6)
	fmt.Println(result)
}
