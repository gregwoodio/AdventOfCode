package main

import (
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")
	expected := 1
	actual := partOne(input)

	if expected != actual {
		t.Errorf("Expected %d, was actually %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")
	expected := 10
	actual := partTwo(input)

	if expected != actual {
		t.Errorf("Expected %d, was actually %d", expected, actual)
	}
}
