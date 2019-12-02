package day02

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input    string
	expected int
}

func TestProcessIntCode(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    "1,9,10,3,2,3,11,0,99,30,40,50",
			expected: 3500,
		},
		testData{
			input:    "1,0,0,0,99",
			expected: 2,
		},
		testData{
			input:    "2,3,0,3,99",
			expected: 2,
		},
		testData{
			input:    "2,4,4,5,99,0",
			expected: 2,
		},
		testData{
			input:    "1,1,1,4,99,5,6,0,99",
			expected: 30,
		},
	}

	for _, td := range testDatas {
		input := parseInput(td.input)
		actual := processIntcode(input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day02_input.txt")
	solution := solvePartOne(input[0])

	fmt.Println(solution)
}
