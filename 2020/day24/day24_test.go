package day24

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func Test2020Day24_SolvePartOne(t *testing.T) {
	testData := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"esew",
			},
			1,
		},
		{
			[]string{
				"nwwswee",
			},
			1,
		},
		{
			[]string{
				"sesenwnenenewseeswwswswwnenewsewsw",
				"neeenesenwnwwswnenewnwwsewnenwseswesw",
				"seswneswswsenwwnwse",
				"nwnwneseeswswnenewneswwnewseswneseene",
				"swweswneswnenwsewnwneneseenw",
				"eesenwseswswnenwswnwnwsewwnwsene",
				"sewnenenenesenwsewnenwwwse",
				"wenwwweseeeweswwwnwwe",
				"wsweesenenewnwwnwsenewsenwwsesesenwne",
				"neeswseenwwswnwswswnw",
				"nenwswwsewswnenenewsenwsenwnesesenew",
				"enewnwewneswsewnwswenweswnenwsenwsw",
				"sweneswneswneneenwnewenewwneswswnese",
				"swwesenesewenwneswnwwneseswwne",
				"enesenwswwswneneswsenwnewswseenwsese",
				"wnwnesenesenenwwnenwsewesewsesesew",
				"nenewswnwewswnenesenwnesewesw",
				"eneswnwswnwsenenwnwnwwseeswneewsenese",
				"neswnwewnwnwseenwseesewsenwsweewe",
				"wseweeenwnesenwwwswnew",
			},
			10,
		},
		{
			[]string{
				"nese",
				"sene",
				"w",
			},
			1,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day24_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day24_input.txt"))
	fmt.Println(solution)
}

func Test2020Day24_SolvePartTwo(t *testing.T) {
	testData := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"sesenwnenenewseeswwswswwnenewsewsw",
				"neeenesenwnwwswnenewnwwsewnenwseswesw",
				"seswneswswsenwwnwse",
				"nwnwneseeswswnenewneswwnewseswneseene",
				"swweswneswnenwsewnwneneseenw",
				"eesenwseswswnenwswnwnwsewwnwsene",
				"sewnenenenesenwsewnenwwwse",
				"wenwwweseeeweswwwnwwe",
				"wsweesenenewnwwnwsenewsenwwsesesenwne",
				"neeswseenwwswnwswswnw",
				"nenwswwsewswnenenewsenwsenwnesesenew",
				"enewnwewneswsewnwswenweswnenwsenwsw",
				"sweneswneswneneenwnewenewwneswswnese",
				"swwesenesewenwneswnwwneseswwne",
				"enesenwswwswneneswsenwnewswseenwsese",
				"wnwnesenesenenwwnenwsewesewsesesew",
				"nenewswnwewswnenesenwnesewesw",
				"eneswnwswnwsenenwnwnwwseeswneewsenese",
				"neswnwewnwnwseenwseesewsenwsweewe",
				"wseweeenwnesenwwwswnew",
			},
			2208,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2020Day24_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day24_input.txt"))
	fmt.Println(solution)
}
