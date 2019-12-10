package day05

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input     string
	expected  int64
	userInput string
}

func TestSolvePartOne(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day05_input.txt")

	result := solve(input[0], 1)

	fmt.Printf("Done:\n\tOutput: %d\n", result)

}

func TestSolvePartTwo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day05_input.txt")

	result := solve(input[0], 5)

	fmt.Printf("Done:\n\tOutput: %d\n", result)
}
