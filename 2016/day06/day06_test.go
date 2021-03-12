package day06

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2016Day06_SolvePartOne(t *testing.T) {
	inputs := []string{
		"eedadn",
		"drvtee",
		"eandsr",
		"raavrd",
		"atevrs",
		"tsrnev",
		"sdttsa",
		"rasrtv",
		"nssdts",
		"ntnada",
		"svetve",
		"tesnvt",
		"vntsnd",
		"vrdear",
		"dvrsen",
		"enarar",
	}

	expected := "easter"
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func Test2016Day06_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day06_input.txt"))
	fmt.Println(solution)
}

func Test2016Day06_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"eedadn",
		"drvtee",
		"eandsr",
		"raavrd",
		"atevrs",
		"tsrnev",
		"sdttsa",
		"rasrtv",
		"nssdts",
		"ntnada",
		"svetve",
		"tesnvt",
		"vntsnd",
		"vrdear",
		"dvrsen",
		"enarar",
	}

	expected := "advent"
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func Test2016Day06_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day06_input.txt"))
	fmt.Println(solution)
}
