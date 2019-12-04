package day01

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input    []int
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{12},
			expected: 2,
		},
		testData{
			input:    []int{14},
			expected: 2,
		},
		testData{
			input:    []int{1969},
			expected: 654,
		},
		testData{
			input:    []int{100756},
			expected: 33583,
		},
	}

	for _, td := range testDatas {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadIntsFromFile("./day01_input.txt")
	solution := solvePartOne(input)

	fmt.Println(solution)
}

func TestSolvePartTwo(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{14},
			expected: 2,
		},
		testData{
			input:    []int{1969},
			expected: 966,
		},
		testData{
			input:    []int{100756},
			expected: 50346,
		},
	}

	for _, td := range testDatas {
		actual := solvePartTwo(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadIntsFromFile("./day01_input.txt")
	solution := solvePartTwo(input)

	fmt.Println(solution)
}
