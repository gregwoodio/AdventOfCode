package day15

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day15_SolvePartOne(t *testing.T) {
	inputs := []string{
		"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
		"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
		"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
		"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
		"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
		"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
		"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
		"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
		"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
		"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
		"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
		"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
		"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
		"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
	}

	expected := 26
	actual := solvePartOne(inputs, 10)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day15_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day15_input.txt"), 2000000)
	fmt.Println(solution)
}

func Test2022Day15_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
		"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
		"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
		"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
		"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
		"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
		"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
		"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
		"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
		"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
		"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
		"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
		"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
		"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
	}

	expected := -1
	actual := solvePartTwo(inputs, 10)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day15_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day15_input.txt"), 2000000)
	fmt.Println(solution)
}
