package day15

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

type coord struct {
	x, y int
}

func Test2018Day15_MakeBoard(t *testing.T) {
	testData := struct {
		input         []string
		expectedUnits []*coord
	}{
		input: []string{
			`#######`,
			`#.G.E.#`,
			`#E.G.E#`,
			`#.G.E.#`,
			`#######`,
		},
		expectedUnits: []*coord{
			{
				x: 4, y: 1,
			},
			{
				x: 1, y: 2,
			},
			{
				x: 5, y: 2,
			},
			{
				x: 4, y: 3,
			},
			{
				x: 2, y: 1,
			},
			{
				x: 3, y: 2,
			},
			{
				x: 2, y: 3,
			},
		},
	}

	board := makeBoard(testData.input, 3)
	if len(board.units) != len(testData.expectedUnits) {
		t.Errorf("Expected %d, but was %d\n", len(testData.expectedUnits), len(board.units))
	}

	found := false
	for _, expected := range testData.expectedUnits {
		for _, unit := range board.units {
			if expected.x == unit.location.x && expected.y == unit.location.y {
				found = true
			}
		}

		if !found {
			t.Errorf("Expected elf at location x: %d, y: %d but it was missing\n", expected.x, expected.y)
		}
		found = false
	}

	board.print()
}

func Test2018Day15_Move(t *testing.T) {
	testData := struct {
		input    []string
		expected coord
	}{
		input: []string{
			`#######`,
			`#E..G.#`,
			`#...#.#`,
			`#.G.#G#`,
			`#######`,
		},
		expected: coord{
			x: 3,
			y: 1,
		},
	}

	board := makeBoard(testData.input, 3)
	elf := board.units[0]

	elf.move(board)
	elf.move(board)

	if elf.location.x != testData.expected.x || elf.location.y != testData.expected.y {
		t.Errorf("Expected elf to be at x: %d y: %d, but was at x: %d, y: %d\n", testData.expected.x, testData.expected.y, elf.location.x, elf.location.y)
	}
}

func Test2018Day15_LargerMoveExample(t *testing.T) {
	input := []string{
		`#########`,
		`#G..G..G#`,
		`#.......#`,
		`#.......#`,
		`#G..E..G#`,
		`#.......#`,
		`#.......#`,
		`#G..G..G#`,
		`#########`,
	}

	testData := []struct {
		coords []*coord
	}{
		{
			[]*coord{
				{x: 2, y: 1},
				{x: 6, y: 1},
				{x: 4, y: 2},
				{x: 4, y: 3},
				{x: 7, y: 3},
				{x: 2, y: 4},
				{x: 1, y: 6},
				{x: 4, y: 6},
				{x: 7, y: 6},
			},
		},
		{
			[]*coord{
				{x: 3, y: 1},
				{x: 5, y: 1},
				{x: 4, y: 2},
				{x: 2, y: 3},
				{x: 4, y: 3},
				{x: 6, y: 3},
				{x: 1, y: 5},
				{x: 4, y: 5},
				{x: 7, y: 5},
			},
		},
		{
			[]*coord{
				{x: 3, y: 2},
				{x: 4, y: 2},
				{x: 5, y: 2},
				{x: 3, y: 3},
				{x: 4, y: 3},
				{x: 5, y: 3},
				{x: 1, y: 4},
				{x: 4, y: 4},
				{x: 7, y: 5},
			},
		},
	}

	b := makeBoard(input, 3)
	b.print()
	for _, td := range testData {
		sort.Sort(b.units)

		for _, u := range b.units {
			target := u.findTarget(b)
			if target == nil {
				u.move(b)
			}
		}

		b.turn++
		b.print()

		found := false
		for _, c := range td.coords {
			for _, u := range b.units {
				if u.location.x == c.x && u.location.y == c.y {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Units was expected in x: %d y: %d after round %d but was not", c.x, c.y, b.turn)
			}
			found = false
		}
	}
}

func Test2018Day15_Target(t *testing.T) {
	input := []string{
		`#####`,
		`#.G.#`,
		`#.EG#`,
		`#.G.#`,
		`#####`,
	}

	b := makeBoard(input, 3)

	// setup goblins
	b.units[0].hp = 9
	b.units[2].hp = 2
	b.units[3].hp = 2

	// elf attack
	target := b.units[1].findTarget(b)

	if target != b.units[2] {
		t.Errorf("Selected wrong target")
	}
}

func Test2018Day15_SampleCombat(t *testing.T) {
	testData := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				`#######`,
				`#.G...#`,
				`#...EG#`,
				`#.#.#G#`,
				`#..G#E#`,
				`#.....#`,
				`#######`,
			},
			expected: 27730,
		},
		{
			input: []string{
				`#######`,
				`#G..#E#`,
				`#E#E.E#`,
				`#G.##.#`,
				`#...#E#`,
				`#...E.#`,
				`#######`,
			},
			expected: 36334,
		},
		{
			input: []string{
				`#######`,
				`#E..EG#`,
				`#.#G.E#`,
				`#E.##E#`,
				`#G..#.#`,
				`#..E#.#`,
				`#######`,
			},
			expected: 39514,
		},
		{
			input: []string{
				`#######`,
				`#E.G#.#`,
				`#.#G..#`,
				`#G.#.G#`,
				`#G..#.#`,
				`#...E.#`,
				`#######`,
			},
			expected: 27755,
		},
		{
			input: []string{
				`#######`,
				`#.E...#`,
				`#.#..G#`,
				`#.###.#`,
				`#E#G#G#`,
				`#...#G#`,
				`#######`,
			},
			expected: 28944,
		},
		{
			input: []string{
				`#########`,
				`#G......#`,
				`#.E.#...#`,
				`#..##..G#`,
				`#...##..#`,
				`#...#...#`,
				`#.G...G.#`,
				`#.....G.#`,
				`#########`,
			},
			expected: 18740,
		},
	}

	for _, td := range testData {
		actual := solvePartOne(td.input, false)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2018Day15_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual test solution")
	}

	result := solvePartOne(aocutil.ReadStringsFromFile("./day15_input.txt"), false)

	fmt.Println(result)
}

func Test2018Day15_SampleCombatPartTwo(t *testing.T) {
	testData := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				`#######`,
				`#.G...#`,
				`#...EG#`,
				`#.#.#G#`,
				`#..G#E#`,
				`#.....#`,
				`#######`,
			},
			expected: 4988,
		},
		{
			input: []string{
				`#######`,
				`#E..EG#`,
				`#.#G.E#`,
				`#E.##E#`,
				`#G..#.#`,
				`#..E#.#`,
				`#######`,
			},
			expected: 31284,
		},
		{
			input: []string{
				`#######`,
				`#E.G#.#`,
				`#.#G..#`,
				`#G.#.G#`,
				`#G..#.#`,
				`#...E.#`,
				`#######`,
			},
			expected: 3478,
		},
		{
			input: []string{
				`#######`,
				`#.E...#`,
				`#.#..G#`,
				`#.###.#`,
				`#E#G#G#`,
				`#...#G#`,
				`#######`,
			},
			expected: 6474,
		},
		{
			input: []string{
				`#########`,
				`#G......#`,
				`#.E.#...#`,
				`#..##..G#`,
				`#...##..#`,
				`#...#...#`,
				`#.G...G.#`,
				`#.....G.#`,
				`#########`,
			},
			expected: 1140,
		},
	}

	for _, td := range testData {
		actual := solvePartTwo(td.input, false)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func Test2018Day15_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual test solution")
	}

	result := solvePartTwo(aocutil.ReadStringsFromFile("./day15_input.txt"), false)

	fmt.Println(result)
}
