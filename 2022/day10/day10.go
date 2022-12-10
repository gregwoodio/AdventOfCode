package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	x := 1
	tick := 1
	delay := 0
	inst := 0
	targetTick := 20
	val := 0
	sum := 0
	tickMul := 0

	// draw the screen
	if x-1 == tick || x == tick || x+1 == tick {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}

	for {
		if tick == targetTick {
			sum += tick * x
			targetTick += 40
		}

		if tick%40 == 0 {
			fmt.Printf("\n")
			tickMul += 40
		}

		if delay == 0 {
			//read instruction
			if inst >= len(input) {
				break
			}

			parts := strings.Split(input[inst], " ")
			if parts[0] == "addx" {
				v, err := strconv.Atoi(parts[1])
				if err != nil {
					panic("Could not parse value")
				}
				val = v
				delay = 1
			} else {
				// noop
				val = 0
				delay = 0
			}
			inst++
		} else {
			delay--
			if delay == 0 {
				x += val
				val = 0
			}
		}

		// draw the screen
		if x-1 == tick-tickMul || x == tick-tickMul || x+1 == tick-tickMul {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if tick > 240 {
			break
		}

		tick++
	}

	fmt.Printf("\n")

	return sum
}
