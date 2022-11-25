package main

import (
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")
	actual := partOne(input)
	expected := 24

	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")
	actual := partTwo(input)
	expected := 10

	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}
