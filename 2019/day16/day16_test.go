package day16

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

type testData struct {
	input    string
	expected []int
	phases   int
}

func TestSolvePartOne(t *testing.T) {
	testCases := []testData{
		testData{
			input:    "12345678",
			expected: []int{0, 1, 0, 2, 9, 4, 9, 8},
			phases:   4,
		},
		testData{
			input:    "80871224585914546619083218645595",
			expected: []int{2, 4, 1, 7, 6, 1, 7, 6},
			phases:   100,
		},
		testData{
			input:    "19617804207202209144916044189917",
			expected: []int{7, 3, 7, 4, 5, 4, 1, 8},
			phases:   100,
		},
		testData{
			input:    "69317163492948606335995924319873",
			expected: []int{5, 2, 4, 3, 2, 1, 3, 3},
			phases:   100,
		},
	}

	for _, td := range testCases {
		actual := solvePartOne(td.input, td.phases)

		for i := range actual {
			if actual[i] != td.expected[i] {
				t.Errorf("Expected %d but was %d\n", td.expected[i], actual[i])
			}
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day16_input.txt")[0]
	result := solvePartOne(input, 100)
	fmt.Println(result)
}
