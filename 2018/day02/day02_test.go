package day02

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData1 struct {
	input    []string
	expected int
}

type testData2 struct {
	input    []string
	expected string
}

func TestSolvePartOne(t *testing.T) {
	testCase := testData1{
		[]string{
			"abcdef",
			"bababc",
			"abbcde",
			"abcccd",
			"aabcdd",
			"abcdee",
			"ababab",
		},
		12,
	}

	actual := solvePartOne(testCase.input)

	if actual != testCase.expected {
		t.Errorf("Expected %d but was %d\n", testCase.expected, actual)
	}

	fmt.Println(actual)
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day02_input.txt")
	result := solvePartOne(input)
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {
	testCase := testData2{
		[]string{
			"abcde",
			"fghij",
			"klmno",
			"pqrst",
			"fguij",
			"axcye",
			"wvxyz",
		},
		"fgij",
	}

	actual := solvePartTwo(testCase.input)

	if actual != testCase.expected {
		t.Errorf("Expected %s but was %s\n", testCase.expected, actual)
	}

	fmt.Println(actual)
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day02_input.txt")
	result := solvePartTwo(input)
	fmt.Println(result)
}
