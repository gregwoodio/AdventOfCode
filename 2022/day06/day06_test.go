package day06

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2022Day06_SolvePartOne(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 7,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2022Day06_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day06_input.txt")[0])
	fmt.Println(solution)
}

func Test2022Day06_SolvePartTwo(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2022Day06_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day06_input.txt")[0])
	fmt.Println(solution)
}
