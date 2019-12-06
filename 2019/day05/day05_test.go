package day05

import (
	"bytes"
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
		testData{
			input:    "1002,4,3,4,33",
			expected: 1002,
		},
		testData{
			input:    "1101,100,-1,4,0",
			expected: 1101,
		},
		testData{
			input:    "3,1,4,1,99",
			expected: 3,
		},
	}

	var mockInput bytes.Buffer
	mockInput.Write([]byte("42\n"))

	for _, td := range testDatas {
		input := parseInput(td.input)
		actual := processIntcode(input, &mockInput)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartOne(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day05_input.txt")

	var mockInput bytes.Buffer
	mockInput.Write([]byte("1\n"))

	result := solvePartOne(input[0], &mockInput)
	fmt.Printf("Done: %d\n", result)
}

// func TestIsValidPartTwo(t *testing.T) {
// 	testDatas := []testData{
// 		testData{
// 			input:    112233,
// 			expected: true,
// 		},
// 		testData{
// 			input:    123444,
// 			expected: false,
// 		},
// 		testData{
// 			input:    111122,
// 			expected: true,
// 		},
// 	}

// 	for _, td := range testDatas {
// 		actual := isValid(td.input, true)
// 		if actual != td.expected {
// 			t.Errorf("Expected %t but was %t\n for value %d", td.expected, actual, td.input)
// 		}
// 	}
// }

// func TestSolvePartTwo(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping actual solution run.")
// 	}

// 	result := solvePartTwo(240298, 784956)
// 	fmt.Println(result)
// }
