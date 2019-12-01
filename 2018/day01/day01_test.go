package day01

import (
	"aocutil"
	"fmt"
	"testing"
)

type testData struct {
	input    []string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testCases := []testData{
		testData{
			[]string{"+1", "+1", "+1"},
			3,
		},
		testData{
			[]string{"+1", "+1", "-2"},
			0,
		},
		testData{
			[]string{"-1", "-2", "-3"},
			-6,
		},
	}

	for _, testCase := range testCases {
		actual := solvePartOne(testCase.input)
		if actual != testCase.expected {
			t.Errorf("Expected %d but was %d", testCase.expected, actual)
		}
		fmt.Println(actual)
	}
}

func TestSolvePartOneActual(t *testing.T) {
	input := aocutil.ReadStringsFromFile("./day01_input.txt")
	actual := solvePartOne(input)
	fmt.Println(actual)
}

func TestSolvePartTwo(t *testing.T) {
	testCases := []testData{
		testData{
			[]string{"+1", "-1"},
			0,
		},
		testData{
			[]string{"+3", "+3", "+4", "-2", "-4"},
			10,
		},
		testData{
			[]string{"-6", "+3", "+8", "+5", "-6"},
			5,
		},
		testData{
			[]string{"+7", "+7", "-2", "-7", "-4"},
			14,
		},
	}

	for _, testCase := range testCases {
		actual := solvePartTwo(testCase.input)
		if actual != testCase.expected {
			t.Errorf("Expected %d but was %d", testCase.expected, actual)
		}
		fmt.Println(actual)
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	input := aocutil.ReadStringsFromFile("./day01_input.txt")
	actual := solvePartTwo(input)
	fmt.Println(actual)
}
