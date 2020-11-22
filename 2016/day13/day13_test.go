package day13

import (
	"fmt"
	"testing"
)

func TestSolvePartOne(t *testing.T) {

	td := struct {
		rows        int
		cols        int
		designerNum int
		tX          int
		tY          int
	}{
		rows:        7,
		cols:        10,
		designerNum: 10,
		tX:          7,
		tY:          4,
	}

	actual := solvePartOne(td.rows, td.cols, td.tX, td.tY, td.designerNum)
	expected := 11

	if actual != expected {
		t.Errorf("Expected %d, but was %d\n", expected, actual)
	}
}

func TestSolvePartOneActual(t *testing.T) {
	if testing.Short() {
		return
	}

	tX, tY := 31, 39
	designerNum := 1352
	buffer := 25 // space to expand the maze

	result := solvePartOne(tX+buffer, tY+buffer, tX, tY, designerNum)
	fmt.Println(result)
}

func TestSolvePartTwo(t *testing.T) {

	td := struct {
		rows        int
		cols        int
		designerNum int
	}{
		rows:        7,
		cols:        10,
		designerNum: 10,
	}

	actual := solvePartTwo(td.rows, td.cols, td.designerNum)
	expected := 19

	if actual != expected {
		t.Errorf("Expected %d, but was %d\n", expected, actual)
	}
}

func TestSolvePartTwoActual(t *testing.T) {
	if testing.Short() {
		return
	}

	tX, tY := 31, 39
	designerNum := 1352
	buffer := 25 // space to expand the maze

	result := solvePartTwo(tX+buffer, tY+buffer, designerNum)
	fmt.Println(result)
}
