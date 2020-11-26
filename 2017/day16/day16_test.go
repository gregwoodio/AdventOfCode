package main

import (
	"io/ioutil"
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := "baedc"
	sample, _ := ioutil.ReadFile("sample.txt")
	actual := partOne([]byte("abcde"), string(sample))

	if expected != actual {
		t.Errorf("Expected '%s', received '%s'", expected, actual)
	}
}
