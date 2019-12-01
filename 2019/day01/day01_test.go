package day01

import (
	"aocutil"
	"fmt"
	"testing"
)

type testData struct {
	input    int
	expected int
}

type testData2 struct {
	input    []int
	expected int
}

func TestCalculateFuel(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    12,
			expected: 2,
		},
		testData{
			input:    14,
			expected: 2,
		},
		testData{
			input:    1969,
			expected: 654,
		},
		testData{
			input:    100756,
			expected: 33583,
		},
	}

	for _, td := range testDatas {
		actual := calculateFuel(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartOneActual(t *testing.T) {
	input := aocutil.ReadIntsFromFile("./day01_input.txt")
	solution := SolvePartOne(input)

	fmt.Println(solution)
}

func TestSolvePartTwo(t *testing.T) {
	testDatas := []testData2{
		testData2{
			input:    []int{14},
			expected: 2,
		},
		testData2{
			input:    []int{1969},
			expected: 966,
		},
		testData2{
			input:    []int{100756},
			expected: 50346,
		},
	}

	for _, td := range testDatas {
		actual := SolvePartTwo(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	input := aocutil.ReadIntsFromFile("./day01_input.txt")
	solution := SolvePartTwo(input)

	fmt.Println(solution)
}
