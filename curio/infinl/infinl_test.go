package infinl

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2030Day01_SolvePartOne(t *testing.T) {
	inputs := []string{
		"draai 90",
		"loop 6",
		"spring 2",
		"draai -45",
		"loop 2",
	}

	expected := 12
	actual := solve(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2030Day01_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solve(aocutil.ReadStringsFromFile("./infinl.txt"))
	fmt.Println(solution)
}
