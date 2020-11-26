package main

import "fmt"

func main() {
	input := 304

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(stepsForward int) int {
	buffer := []int{0}
	curr := 0

	for count := 1; count < 2018; count++ {
		curr = ((curr + stepsForward) % len(buffer)) + 1

		// insert
		buffer = append(buffer[:curr], append([]int{count}, buffer[curr:]...)...)
	}
	return buffer[curr+1]
}

func partTwo(stepsForward int) int {
	curr := 1
	pos := 0
	out := 0
	for i := 0; i < 50000000; i++ {
		toIns := i + 1
		new := (pos + stepsForward) % curr
		new++
		if new == 1 {
			out = toIns
		}
		pos = new
		curr++
	}

	return out
}
