package day04

import (
	"fmt"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	actual := solvePartOne("abcdef")
	expected := 609043

	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	fmt.Println(solvePartOne("bgvyzdsv"))
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	fmt.Println(solvePartTwo("bgvyzdsv"))
}
