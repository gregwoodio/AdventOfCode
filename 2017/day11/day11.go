package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	fmt.Println(partOne("input.txt"))
	fmt.Println(partTwo("input.txt"))
}

func partOne(filename string) int {
	input, _ := ioutil.ReadFile(filename)
	steps := strings.Split(string(input), ",")

	var stepsNorth, stepsEast float64

	for _, step := range steps {
		// fmt.Println(step)
		switch step {
		case "n":
			stepsNorth++
		case "ne":
			stepsNorth += 0.5
			stepsEast++
		case "se":
			stepsNorth -= 0.5
			stepsEast++
		case "s":
			stepsNorth--
		case "sw":
			stepsNorth -= 0.5
			stepsEast--
		case "nw":
			stepsNorth += 0.5
			stepsEast--
		}
	}

	// fmt.Printf("North: %.1f, East: %.1f\n", stepsNorth, stepsEast)

	stepsEast = math.Abs(stepsEast)
	stepsNorth = math.Abs(stepsNorth)

	count := 0
	for ; stepsEast > 0; stepsEast-- {
		stepsNorth -= 0.5
		count++
	}

	for ; stepsNorth > 0; stepsNorth-- {
		count++
	}

	return count
}

func partTwo(filename string) int {
	input, _ := ioutil.ReadFile(filename)
	steps := strings.Split(string(input), ",")

	var stepsNorth, stepsEast float64
	var max int

	for _, step := range steps {
		// fmt.Println(step)
		switch step {
		case "n":
			stepsNorth++
		case "ne":
			stepsNorth += 0.5
			stepsEast++
		case "se":
			stepsNorth -= 0.5
			stepsEast++
		case "s":
			stepsNorth--
		case "sw":
			stepsNorth -= 0.5
			stepsEast--
		case "nw":
			stepsNorth += 0.5
			stepsEast--
		}

		stepsBack := countStepsBack(stepsEast, stepsNorth)
		if stepsBack > max {
			max = stepsBack
		}
	}

	// fmt.Printf("North: %.1f, East: %.1f\n", stepsNorth, stepsEast)
	return max
}

func countStepsBack(stepsEast, stepsNorth float64) int {
	count := 0
	east := math.Abs(stepsEast)
	north := math.Abs(stepsNorth)
	for ; east > 0; east-- {
		north -= 0.5
		count++
	}

	for ; north > 0; north-- {
		count++
	}

	return count
}
