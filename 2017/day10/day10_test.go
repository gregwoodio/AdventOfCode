package main

import (
	"testing"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadAndSplitInts("sample.txt", ",")[0]
	actual := partOne(input, 5)
	expected := 12

	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}
