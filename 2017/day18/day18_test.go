package main

import (
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")

	expected := 4
	actual := partOne(input)

	if expected != actual {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}
