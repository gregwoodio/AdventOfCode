package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 638
	actual := partOne(3)

	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}