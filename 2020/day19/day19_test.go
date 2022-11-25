package day19

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2020Day19_SolvePartOne(t *testing.T) {
	testData := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				`0: 1 2`,
				`1: "a"`,
				`2: 1 3 | 3 1`,
				`3: "b"`,
				``,
				// valid
				`aab`,
				`aba`,
				// invalid
				`a`,
				`b`,
				`ab`,
				`ba`,
				`baa`,
				`bba`,
				`bab`,
				`aaa`,
				`bbb`,
			},
			2,
		},
		{
			[]string{
				`0: 4 1 5`,
				`1: 2 3 | 3 2`,
				`2: 4 4 | 5 5`,
				`3: 4 5 | 5 4`,
				`4: "a"`,
				`5: "b"`,
				``,
				`ababbb`,
				`bababa`,
				`abbbab`,
				`aaabbb`,
				`aaaabbb`,
			},
			2,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day19_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day19_input.txt"))
	fmt.Println(solution)
}

func Test2020Day19_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day19_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day19_input.txt"))
	fmt.Println(solution)
}
