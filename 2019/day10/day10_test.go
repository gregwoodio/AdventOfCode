package day10

import "testing"

type testData struct {
	input    []string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testCases := []testData{
		testData{
			input: []string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			expected: 8,
		},
		testData{
			input: []string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			expected: 33,
		},
	}

	for _, td := range testCases {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}
