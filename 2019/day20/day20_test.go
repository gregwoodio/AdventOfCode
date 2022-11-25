package day20

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

type testData struct {
	input    []string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	testCases := []testData{
		testData{
			input: []string{
				"        ",
				"        ",
				"  ###   ",
				"AA..##  ",
				"  #...ZZ",
				"  ####  ",
				"        ",
				"        ",
			},
			expected: 4,
		},
		testData{
			input: []string{
				"         A           ",
				"         A           ",
				"  #######.#########  ", //  2
				"  #######.........#  ", //  3
				"  #######.#######.#  ", //  4
				"  #######.#######.#  ", //  5
				"  #######.#######.#  ", //  6
				"  #####  B    ###.#  ", //  7
				"BC...##  C    ###.#  ", //  8
				"  ##.##       ###.#  ", //  9
				"  ##...DE  F  ###.#  ", // 10
				"  #####    G  ###.#  ", // 11
				"  #########.#####.#  ", // 12
				"DE..#######...###.#  ", // 13
				"  #.#########.###.#  ", // 14
				"FG..#########.....#  ", // 15
				"  ###########.#####  ", // 16
				"             Z       ",
				"             Z       ",
			},
			expected: 23,
		},
		testData{
			input: []string{
				"                   A               ",
				"                   A               ",
				"  #################.#############  ", //  3
				"  #.#...#...................#.#.#  ", //  4
				"  #.#.#.###.###.###.#########.#.#  ", //  5
				"  #.#.#.......#...#.....#.#.#...#  ", //  6
				"  #.#########.###.#####.#.#.###.#  ", //  7
				"  #.............#.#.....#.......#  ", //  8
				"  ###.###########.###.#####.#.#.#  ", //  9
				"  #.....#        A   C    #.#.#.#  ", // 10
				"  #######        S   P    #####.#  ", // 11
				"  #.#...#                 #......VT", // 12
				"  #.#.#.#                 #.#####  ", // 13
				"  #...#.#               YN....#.#  ", // 14
				"  #.###.#                 #####.#  ", // 15
				"DI....#.#                 #.....#  ", // 16
				"  #####.#                 #.###.#  ", // 17
				"ZZ......#               QG....#..AS", // 18
				"  ###.###                 #######  ", // 19
				"JO..#.#.#                 #.....#  ", // 20
				"  #.#.#.#                 ###.#.#  ", // 21
				"  #...#..DI             BU....#..LF", // 22
				"  #####.#                 #.#####  ", // 23
				"YN......#               VT..#....QG", // 24
				"  #.###.#                 #.###.#  ", // 25
				"  #.#...#                 #.....#  ", // 26
				"  ###.###    J L     J    #.#.###  ", // 27
				"  #.....#    O F     P    #.#...#  ", // 28
				"  #.###.#####.#.#####.#####.###.#  ", // 29
				"  #...#.#.#...#.....#.....#.#...#  ", // 30
				"  #.#####.###.###.#.#.#########.#  ", // 31
				"  #...#.#.....#...#.#.#.#.....#.#  ", // 32
				"  #.###.#####.###.###.#.#.#######  ", // 33
				"  #.#.........#...#.............#  ", // 34
				"  #########.###.###.#############  ", // 35
				"           B   J   C               ",
				"           U   P   P               ",
			},
			expected: 58,
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
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day20_input.txt")
	result := solvePartOne(input)
	fmt.Println(result)
}
