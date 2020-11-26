package main

import (
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")
	tree := buildTree(input)
	expected := "tknk"
	actual := partOne(tree)

	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := aocutil.ReadStringsFromFile("sample.txt")
	tree := buildTree(input)
	expected := 60
	actual := partTwo(tree)

	if actual != expected {
		t.Errorf("Expected %d, received %d", expected, actual)
	}
}
