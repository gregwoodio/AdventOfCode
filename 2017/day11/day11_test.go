package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	actual := partOne("sample1.txt")
	expected := 0

	if actual != expected {
		t.Errorf("Expected %d, was %d", expected, actual)
	}

	actual = partOne("sample2.txt")
	expected = 3

	if actual != expected {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}