package day04

import (
	"fmt"
	"testing"
)

type testData struct {
	input    int
	expected bool
}

func TestIsValid(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    122345,
			expected: true,
		},
		testData{
			input:    111123,
			expected: true,
		},
		testData{
			input:    135679,
			expected: false,
		},
		testData{
			input:    111111,
			expected: true,
		},
		testData{
			input:    223450,
			expected: false,
		},
		testData{
			input:    123789,
			expected: false,
		},
	}

	for _, td := range testDatas {
		actual := isValid(td.input, false)
		if actual != td.expected {
			t.Errorf("Expected %t but was %t\n for value %d", td.expected, actual, td.input)
		}
	}
}

func TestSolvePartOne(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	result := solvePartOne(240298, 784956)
	fmt.Println(result)
}

func TestIsValidPartTwo(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    112233,
			expected: true,
		},
		testData{
			input:    123444,
			expected: false,
		},
		testData{
			input:    111122,
			expected: true,
		},
	}

	for _, td := range testDatas {
		actual := isValid(td.input, true)
		if actual != td.expected {
			t.Errorf("Expected %t but was %t\n for value %d", td.expected, actual, td.input)
		}
	}
}

func TestSolvePartTwo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	result := solvePartTwo(240298, 784956)
	fmt.Println(result)
}
