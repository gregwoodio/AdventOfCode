package day07

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input    string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			expected: 43210,
		},
		testData{
			input:    "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			expected: 54321,
		},
		testData{
			input:    "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			expected: 65210,
		},
	}

	for _, td := range testDatas {
		result := solvePartOne(td.input)
		fmt.Println(result)
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day07_input.txt")

	result := solvePartOne(input[0])
	fmt.Println(result)
}
