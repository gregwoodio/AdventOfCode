package day25

import (
	"fmt"
	"testing"
)

func Test2020Day25_SolvePartOne(t *testing.T) {
	cardPubKey := 5764801
	doorPubKey := 17807724

	expected := 14897079
	actual := solvePartOne(cardPubKey, doorPubKey)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func Test2020Day25_SolvePartOneActual(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping actual solution")
	}

	solution := solvePartOne(3248366, 4738476)
	fmt.Println(solution)
}
