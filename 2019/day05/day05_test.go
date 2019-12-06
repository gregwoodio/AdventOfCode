package day05

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input     string
	expected  int
	userInput string
}

func TestProcessIntCode(t *testing.T) {
	testDatas := []testData{
		// part one tests (day 2 as well)
		testData{
			input:     "1,9,10,3,2,3,11,0,99,30,40,50",
			expected:  3500,
			userInput: "",
		},
		testData{
			input:     "1,0,0,0,99",
			expected:  2,
			userInput: "",
		},
		testData{
			input:     "2,3,0,3,99",
			expected:  2,
			userInput: "",
		},
		testData{
			input:     "2,4,4,5,99,0",
			expected:  2,
			userInput: "",
		},
		testData{
			input:     "1,1,1,4,99,5,6,0,99",
			expected:  30,
			userInput: "",
		},
		testData{
			input:     "1002,4,3,4,33",
			expected:  1002,
			userInput: "",
		},
		testData{
			input:     "1101,100,-1,4,0",
			expected:  1101,
			userInput: "",
		},
		testData{
			input:     "3,1,4,1,99",
			expected:  3,
			userInput: "42\n",
		},
		// part two tests (less than and equal operators)
		testData{
			input:     "3,9,8,9,10,9,4,9,99,-1,8",
			expected:  3,
			userInput: "8\n",
		},
		testData{
			input:     "3,9,8,9,10,9,4,9,99,-1,8",
			expected:  3,
			userInput: "7\n",
		},
		testData{
			input:     "3,9,7,9,10,9,4,9,99,-1,8",
			expected:  3,
			userInput: "8\n",
		},
		testData{
			input:     "3,9,7,9,10,9,4,9,99,-1,8",
			expected:  3,
			userInput: "7\n",
		},
		testData{
			input:     "3,3,1108,-1,8,3,4,3,99",
			expected:  3,
			userInput: "8\n",
		},
		testData{
			input:     "3,3,1108,-1,8,3,4,3,99",
			expected:  3,
			userInput: "7\n",
		},
		testData{
			input:     "3,3,1107,-1,8,3,4,3,99",
			expected:  3,
			userInput: "8\n",
		},
		testData{
			input:     "3,3,1107,-1,8,3,4,3,99",
			expected:  3,
			userInput: "7\n",
		},
		// jump tests
		testData{
			input:     "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			expected:  3,
			userInput: "0\n",
		},
		testData{
			input:     "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			expected:  3,
			userInput: "999\n",
		},
		testData{
			input:     "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			expected:  3,
			userInput: "0\n",
		},
		testData{
			input:     "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			expected:  3,
			userInput: "999\n",
		},
	}

	var mockInput bytes.Buffer

	for _, td := range testDatas {
		input := parseInput(td.input)

		mockInput.Write([]byte(td.userInput))

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
	mockInput.Write([]byte("5\n"))

	result := solvePartOne(input[0], &mockInput)
	fmt.Printf("Done: %d\n", result)
}

func TestSolvePartTwo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day05_input.txt")

	var mockInput bytes.Buffer
	mockInput.Write([]byte("5\n"))

	result := solvePartOne(input[0], &mockInput)
	fmt.Printf("Done: %d\n", result)
}
