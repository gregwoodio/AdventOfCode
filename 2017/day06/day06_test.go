package main

import (
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadAndSplitInts("sample.txt", "\t")

	actual := partOne(input[0])
	expected := 5
	if actual != expected {
		t.Errorf("Expected %d, received %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := aocutil.ReadAndSplitInts("sample.txt", "\t")

	actual := partTwo(input[0])
	expected := 4
	if actual != expected {
		t.Errorf("Expected %d, received %d", expected, actual)
	}
}
