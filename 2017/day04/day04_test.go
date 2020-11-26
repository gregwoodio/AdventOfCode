package main

import (
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample1.txt")
	actual := partOne(input)

	if actual != 2 {
		t.Error("Expected 2, got ", actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample2.txt")
	actual := partTwo(input)

	if actual != 3 {
		t.Error("Expected 3, got ", actual)
	}
}
