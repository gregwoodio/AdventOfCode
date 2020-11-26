package main

import (
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadAndSplitInts("sample.txt", ",")[0]
	actual := partOne(input, 5)
	expected := 12

	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}
