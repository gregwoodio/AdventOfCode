package day10

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input    []string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testCases := []testData{
		// testData{
		// 	input: []string{
		// 		"#.........",
		// 		"...#......",
		// 		"...#..#...",
		// 		".####....#",
		// 		"..#.#.#...",
		// 		".....#....",
		// 		"..###.#.##",
		// 		".......#..",
		// 		"....#...#.",
		// 		"...#..#..#",
		// 	},
		// 	expected: 7,
		// },
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

func TestSolvePartOneActual(t *testing.T) {
	input := aocutil.ReadStringsFromFile("./day10_input.txt")
	result := solvePartOne(input)
	fmt.Println(result)
}
