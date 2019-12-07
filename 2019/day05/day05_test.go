package day05

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregwoodio/aocutil"
)

type testData struct {
	input     string
	expected  int
	userInput string
}

func TestSolvePartOne(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day05_input.txt")

	var mockInput bytes.Buffer
	mockInput.Write([]byte("1\n"))

	var mockOutput bytes.Buffer

	result := solve(input[0], &mockInput, &mockOutput)
	strOut := string(mockOutput.Bytes())

	fmt.Printf("Done: \n\t0th index: %d\n\tOutput: %s\n", result, strOut)

}

func TestSolvePartTwo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution run.")
	}

	input := aocutil.ReadStringsFromFile("./day05_input.txt")

	var mockInput bytes.Buffer
	mockInput.Write([]byte("5\n"))

	var mockOutput bytes.Buffer

	result := solve(input[0], &mockInput, &mockOutput)
	strOut := string(mockOutput.Bytes())

	fmt.Printf("Done: \n\t0th index: %d\n\tOutput: %s\n", result, strOut)
}
