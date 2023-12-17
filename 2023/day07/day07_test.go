package day07

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2023Day07_SolvePartOne(t *testing.T) {
	testData := []struct {
		inputs   []string
		expected int
	}{
		{
			inputs: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			expected: 6440,
		},
		{
			inputs: []string{
				"33332 2",
				"2AAAA 3",
				"77888 4",
				"77788 5",
			},
			expected: 30,
		},
		{
			inputs: []string{
				"2345A 1",
				"Q2KJJ 13",
				"Q2Q2Q 19",
				"T3T3J 17",
				"T3Q33 11",
				"2345J 3",
				"J345A 2",
				"32T3K 5",
				"T55J5 29",
				"KK677 7",
				"KTJJT 34",
				"QQQJA 31",
				"JJJJJ 37",
				"JAAAA 43",
				"AAAAJ 59",
				"AAAAA 61",
				"2AAAA 23",
				"2JJJJ 53",
				"JJJJ2 41",
			},
			expected: 6592,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.inputs)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2023Day07_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day07_input.txt"))
	fmt.Println(solution)
}

func Test2023Day07_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	expected := 5905
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day07_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day07_input.txt"))
	fmt.Println(solution)
}
