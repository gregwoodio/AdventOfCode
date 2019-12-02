package day02

import (
	"log"
	"strconv"
	"strings"
)

func solvePartOne(input string) int {
	inst := parseInput(input)

	// 1202 program state
	inst[1] = 12
	inst[2] = 2

	return processIntcode(inst)
}

func processIntcode(inst []int) int {
	op := 0

	for {
		if inst[op] == 1 {
			// add
			inst[inst[op+3]] = inst[inst[op+1]] + inst[inst[op+2]]
		} else if inst[op] == 2 {
			// multiply
			inst[inst[op+3]] = inst[inst[op+1]] * inst[inst[op+2]]
		} else {
			return inst[0]
		}

		op += 4
	}
}

func parseInput(input string) []int {
	output := []int{}

	split := strings.Split(input, ",")

	for _, strVal := range split {
		intVal, err := strconv.Atoi(strVal)

		if err != nil {
			log.Fatal(err)
		}

		output = append(output, intVal)
	}

	return output
}
