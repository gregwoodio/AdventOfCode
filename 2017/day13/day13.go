package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregwoodio/aocutil"
)

// Layer represents a layer in the firewall
type Layer struct {
	Range     int
	Current   int
	Direction int
}

func main() {
	input := aocutil.ReadStringsFromFile("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func getLayer() Layer {
	return Layer{
		Direction: 0,
		Current:   -1,
		Range:     -1,
	}
}

func partOne(input []string) int {
	return solve(input, 0)
}

func partTwo(input []string) int {
	layers := make(map[int]int)

	for _, line := range input {
		parts := strings.Split(line, ": ")
		depth, _ := strconv.Atoi(parts[0])
		rng, _ := strconv.Atoi(parts[1])

		layers[depth] = rng
	}

	for delay := 0; ; delay++ {
		pass := true
		for depth, rng := range layers {
			div := rng*2 - 2
			if (depth+delay)%div == 0 {
				pass = false
				break
			}
		}
		if pass {
			return delay
		}
	}
}

func solve(input []string, delay int) int {
	layers := make(map[int]Layer)
	maxDepth := 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		depth, _ := strconv.Atoi(parts[0])
		rng, _ := strconv.Atoi(parts[1])

		layer := Layer{
			Current:   0,
			Range:     rng,
			Direction: 1,
		}
		layers[depth] = layer

		if depth > maxDepth {
			maxDepth = depth
		}
	}

	layerSlice := make([]Layer, maxDepth+1)
	for key, value := range layers {
		layerSlice[key] = value
	}

	severity := 0

	for d := 0; d < delay; d++ {
		for i := range layerSlice {
			if layerSlice[i].Direction == 1 {
				layerSlice[i].Current++
			} else if layerSlice[i].Direction == -1 {
				layerSlice[i].Current--
			}

			if (layerSlice[i].Current == 0 && layerSlice[i].Direction == -1) || (layerSlice[i].Current == layerSlice[i].Range-1 && layerSlice[i].Direction == 1) {
				layerSlice[i].Direction *= -1
			}
		}
	}

	for step := 0; step < len(layerSlice); step++ {

		// check collision
		if layerSlice[step].Current == 0 {
			if delay > 0 {
				return -1
			}
			severity += layerSlice[step].Range * step
		}

		// increment scanners
		for i := range layerSlice {
			if layerSlice[i].Direction == 1 {
				layerSlice[i].Current++
			} else if layerSlice[i].Direction == -1 {
				layerSlice[i].Current--
			}

			if (layerSlice[i].Current == 0 && layerSlice[i].Direction == -1) || (layerSlice[i].Current == layerSlice[i].Range-1 && layerSlice[i].Direction == 1) {
				layerSlice[i].Direction *= -1
			}
		}
	}

	return severity
}
