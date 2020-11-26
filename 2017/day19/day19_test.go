package main

import (
	"testing"

	"github.com/gregwoodio/aocutil"
)

func TestSolve(t *testing.T) {
	maze := aocutil.ReadStringsFromFile("sample.txt")
	actualLetters, actualSteps := solve(maze)
	expectedLetters, expectedSteps := "ABCDEF", 38

	if expectedLetters != actualLetters {
		t.Errorf("Expected %s, was %s", expectedLetters, actualLetters)
	}

	if expectedSteps != actualSteps {
		t.Errorf("Expected %d steps, was %d", expectedSteps, actualSteps)
	}
}
