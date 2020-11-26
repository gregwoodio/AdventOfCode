package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	actual := partOne("sample1.txt")
	expected := 1

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample2.txt")
	expected = 6

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample3.txt")
	expected = 5

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample4.txt")
	expected = 16

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample5.txt")
	expected = 1

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample6.txt")
	expected = 9

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample7.txt")
	expected = 9

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partOne("sample8.txt")
	expected = 3

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := partTwo("sample9.txt")
	expected := 10

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partTwo("sample10.txt")
	expected = 0

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partTwo("sample11.txt")
	expected = 17

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partTwo("sample12.txt")
	expected = 3

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partTwo("sample13.txt")
	expected = 2

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partTwo("sample14.txt")
	expected = 0

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}

	actual = partTwo("sample15.txt")
	expected = 0

	if actual != expected {
		t.Errorf("Expected %d, actually %d", expected, actual)
	}
}
