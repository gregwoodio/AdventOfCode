package main

import (
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestPartOne(t *testing.T) {
	input := aocutil.ReadIntsFromFile("sample.txt")
	expected := 5

	actual := partOne(input)
	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := aocutil.ReadIntsFromFile("sample.txt")
	expected := 10

	actual := partTwo(input)

	fmt.Println(input)

	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}
