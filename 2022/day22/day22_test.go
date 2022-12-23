package day22

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day22_SolvePartOne(t *testing.T) {
	inputs := []string{
		"        ...#",
		"        .#..",
		"        #...",
		"        ....",
		"...#.......#",
		"........#...",
		"..#....#....",
		"..........#.",
		"        ...#....",
		"        .....#..",
		"        .#......",
		"        ......#.",
		"",
		"10R5L5R10L4R5L5",
		// "R5R10R4L3",
	}
	// col=10
	// row = 16
	// facing = left
	// 16083

	expected := 6032
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day22_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day22_input.txt"))
	fmt.Println(solution)
}

func Test2022Day22_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day22_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day22_input.txt"))
	fmt.Println(solution)
}
