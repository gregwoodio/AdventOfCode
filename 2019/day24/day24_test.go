package day24

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func TestMakeGrid(t *testing.T) {
	testGrid := makeGrid([]string{
		".....",
		".....",
		".....",
		"#....",
		".#...",
	})

	if testGrid != 2129920 {
		t.Errorf("Expected biodiversity rating of 2129920 but was %d\n", testGrid)
	}
}

func TestIterate(t *testing.T) {
	grid := makeGrid([]string{
		"....#",
		"#..#.",
		"#..##",
		"..#..",
		"#....",
	})
	fmt.Println("Initial state:")
	printGrid(grid)

	for i := 0; i < 4; i++ {
		fmt.Printf("\nAfter %d minutes:\n", i+1)
		grid = iterate(grid)
		printGrid(grid)
	}
}

func TestSolvePartOne(t *testing.T) {
	grid := solvePartOne([]string{
		"....#",
		"#..#.",
		"#..##",
		"..#..",
		"#....",
	})

	if grid != 2129920 {
		t.Errorf("Expected biodiversity rating of 2129920 but was %d\n", grid)
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day24_input.txt")
	result := solvePartOne(input)
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	bugs := solvePartTwo([]string{
		"....#",
		"#..#.",
		"#.?##",
		"..#..",
		"#....",
	}, 10)

	if bugs != 99 {
		t.Errorf("Expected bug count to be 99 but was %d\n", bugs)
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day24_input.txt")
	result := solvePartTwo(input, 200)
	fmt.Println(result)
}
