package day16

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2020Day16_SolvePartOne(t *testing.T) {
	inputs := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}

	expected := 71
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day16_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day16_input.txt"))
	fmt.Println(solution)
}

func Test2020Day16_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13",
		"",
		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}

	expected := 1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day16_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day16_input.txt"))
	fmt.Println(solution)
}
