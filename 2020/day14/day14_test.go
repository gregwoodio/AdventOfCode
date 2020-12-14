package day14

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2020Day14_SolvePartOne(t *testing.T) {
	inputs := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	expected := 165
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day14_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day14_input.txt"))
	fmt.Println(solution)
}

func Test2020Day14_ExpandMask(t *testing.T) {
	testData := []struct {
		input    string
		expected []string
	}{
		{
			input: "101",
			expected: []string{
				"101",
			},
		},
		{
			input: "10X",
			expected: []string{
				"100",
				"101",
			},
		},
		{
			input: "1XXX",
			expected: []string{
				"1000",
				"1001",
				"1010",
				"1011",
				"1100",
				"1101",
				"1110",
				"1111",
			},
		},
	}

	for _, td := range testData {
		actual := expandMask(td.input)
		if len(actual) != len(td.expected) {
			t.Errorf("Expected %d but was %d", len(td.expected), len(actual))
		}

		for _, expected := range td.expected {
			found := false
			for _, mask := range actual {
				if mask == expected {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected mask '%s' was not found", expected)
			}
		}
	}
}

func Test2020Day14_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	var expected int64 = 208
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day14_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day14_input.txt"))
	fmt.Println(solution)
}
