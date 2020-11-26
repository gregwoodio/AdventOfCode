package main

import "testing"

func TestPartOne(t *testing.T) {
	expected := 8108
	actual := partOne("flqrgnkx")

	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 1242
	actual := partTwo("flqrgnkx")

	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}
