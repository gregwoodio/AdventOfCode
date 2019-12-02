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

func solvePartTwo(input string) int {
	target := 19690720

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			inst := parseInput(input)

			inst[1] = noun
			inst[2] = verb

			if processIntcode(inst) == target {
				return 100*noun + verb
			}
		}
	}

	log.Fatal("Did not find a noun and verb pair that returned the target value")
	return -1
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
