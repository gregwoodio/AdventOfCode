package day11

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day11_SolvePartOne(t *testing.T) {
	inputs := []string{
		"Monkey 0:",
		"Starting items: 79, 98",
		"Operation: new = old * 19",
		"Test: divisible by 23",
		"If true: throw to monkey 2",
		"If false: throw to monkey 3",
		"",
		"Monkey 1:",
		"Starting items: 54, 65, 75, 74",
		"Operation: new = old + 6",
		"Test: divisible by 19",
		"If true: throw to monkey 2",
		"If false: throw to monkey 0",
		"",
		"Monkey 2:",
		"Starting items: 79, 60, 97",
		"Operation: new = old * old",
		"Test: divisible by 13",
		"If true: throw to monkey 1",
		"If false: throw to monkey 3",
		"",
		"Monkey 3:",
		"Starting items: 74",
		"Operation: new = old + 3",
		"Test: divisible by 17",
		"If true: throw to monkey 0",
		"If false: throw to monkey 1",
	}

	expected := 10605
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day11_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day11_input.txt"))
	fmt.Println(solution)
}

func Test2022Day11_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"Monkey 0:",
		"Starting items: 79, 98",
		"Operation: new = old * 19",
		"Test: divisible by 23",
		"If true: throw to monkey 2",
		"If false: throw to monkey 3",
		"",
		"Monkey 1:",
		"Starting items: 54, 65, 75, 74",
		"Operation: new = old + 6",
		"Test: divisible by 19",
		"If true: throw to monkey 2",
		"If false: throw to monkey 0",
		"",
		"Monkey 2:",
		"Starting items: 79, 60, 97",
		"Operation: new = old * old",
		"Test: divisible by 13",
		"If true: throw to monkey 1",
		"If false: throw to monkey 3",
		"",
		"Monkey 3:",
		"Starting items: 74",
		"Operation: new = old + 3",
		"Test: divisible by 17",
		"If true: throw to monkey 0",
		"If false: throw to monkey 1",
	}

	expected := 2713310158
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day11_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day11_input.txt"))
	fmt.Println(solution)
}
