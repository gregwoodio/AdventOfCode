package day05

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2023Day05_SolvePartOne(t *testing.T) {
	inputs := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	expected := 35
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day05_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day05_input.txt"))
	fmt.Println(solution)
}

func Test2023Day05_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day05_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day05_input.txt"))
	fmt.Println(solution)
}
