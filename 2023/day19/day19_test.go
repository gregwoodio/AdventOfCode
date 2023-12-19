package day19

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func Test2023Day19_SolvePartOne(t *testing.T) {
	inputs := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"",
		"{x=787,m=2655,a=1222,s=2876}",
		"{x=1679,m=44,a=2067,s=496}",
		"{x=2036,m=264,a=79,s=2244}",
		"{x=2461,m=1339,a=466,s=291}",
		"{x=2127,m=1623,a=2188,s=1013}",
	}

	expected := 19114
	actual := solvePartOne(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day19_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(aocutil.ReadStringsFromFile("./day19_input.txt"))
	fmt.Println(solution)
}

func Test2023Day19_SolvePartTwo(t *testing.T) {
	inputs := []string{
		"foo",
	}

	expected := -1
	actual := solvePartTwo(inputs)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2023Day19_SolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartTwo(aocutil.ReadStringsFromFile("./day19_input.txt"))
	fmt.Println(solution)
}
