package day22

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	size     int
	input    []string
	expected []int
}

func TestSolvePartOne(t *testing.T) {
	testCases := []testData{
		testData{
			size: 10,
			input: []string{
				"deal into new stack",
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		testData{
			size: 10,
			input: []string{
				"cut 3",
			},
			expected: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		testData{
			size: 10,
			input: []string{
				"cut -4",
			},
			expected: []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
		testData{
			size: 10,
			input: []string{
				"deal with increment 3",
			},
			expected: []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
		testData{
			size: 10,
			input: []string{
				"deal with increment 7",
				"deal into new stack",
				"deal into new stack",
			},
			expected: []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		testData{
			size: 10,
			input: []string{
				"cut 6",
				"deal with increment 7",
				"deal into new stack",
			},
			expected: []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		testData{
			size: 10,
			input: []string{
				"deal with increment 7",
				"deal with increment 9",
				"cut -2",
			},
			expected: []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		testData{
			size: 10,
			input: []string{
				"deal into new stack",
				"cut -2",
				"deal with increment 7",
				"cut 8",
				"cut -4",
				"deal with increment 7",
				"cut 3",
				"deal with increment 9",
				"deal with increment 3",
				"cut -1",
			},
			expected: []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}

	for _, td := range testCases {
		actual := solvePartOne(td.size, td.input)
		if len(td.expected) != len(actual) {
			t.Errorf("Expected length of %d but was %d\n", len(td.expected), len(actual))
		}

		for i := range td.expected {
			if td.expected[i] != actual[i] {
				t.Errorf("Expected %d but was %d at index %d\n", td.expected[i], actual[i], i)
			}
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day22_input.txt")
	result := solvePartOne(10007, input)
	for i, val := range result {
		if val == 2019 {
			fmt.Println(i)
			break
		}
	}
}
