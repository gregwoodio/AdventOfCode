package day20

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day20_SolvePartOne(t *testing.T) {
	inputs := []int{
		1,
		2,
		-3,
		3,
		-2,
		0,
		4,
	}

	var expected int64 = 3
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day20_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadIntsFromFile("./day20_input.txt"))
	fmt.Println(solution)
}

// func Test2022Day20_SolvePartTwo(t *testing.T) {
// 	inputs := []int{
// 		1,
// 		2,
// 		-3,
// 		3,
// 		-2,
// 		0,
// 		4,
// 	}

// 	expected := -1
// 	actual := solvePartTwo(inputs)
// 	if actual != expected {
// 		t.Errorf("Expected %d but was %d", expected, actual)
// 	}
// }

// func Test2022Day20_SolvePartTwoActual(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping actual solution")
// 	}

// 	solution := solvePartTwo(aocutil.ReadIntsFromFile("./day21_input.txt"))
// 	fmt.Println(solution)
// }
