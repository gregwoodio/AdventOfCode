package day07

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day07_SolvePartOne(t *testing.T) {
	inputs := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	expected := 95437
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day07_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day07_input.txt"))
	fmt.Println(solution)
}

func Test2022Day07_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	expected := 24933642
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2022Day07_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day07_input.txt"))
	fmt.Println(solution)
}
