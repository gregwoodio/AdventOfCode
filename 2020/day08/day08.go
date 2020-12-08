package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type operation int

const (
	acc operation = iota
	jmp
	nop
)

type instruction struct {
	op       operation
	value    int
	executed bool
}

func solvePartOne(input []string) int {
	instructions := parseInstructions(input)
	accumulator := 0
	index := 0

	for {
		if instructions[index].executed {
			return accumulator
		}

		instructions[index].executed = true
		switch instructions[index].op {
		case acc:
			accumulator += instructions[index].value
			index++
			break
		case jmp:
			index += instructions[index].value
			break
		case nop:
			index++
			break
		}
	}
}

func solvePartTwo(input []string) int {
	instructions := parseInstructions(input)
	accumulator := 0
	index := 0
	switchIndex := findNextJmpOrNop(-1, instructions)

	for {

		for {
			if index == len(instructions) {
				return accumulator
			}

			if instructions[index].executed {
				break
			}

			instructions[index].executed = true
			if index == switchIndex {
				switch instructions[index].op {
				case jmp:
					index++
					break
				case nop:
					index += instructions[index].value
					break
				}
			} else {
				switch instructions[index].op {
				case acc:
					accumulator += instructions[index].value
					index++
					break
				case jmp:
					index += instructions[index].value
					break
				case nop:
					index++
					break
				}
			}
		}

		switchIndex = findNextJmpOrNop(switchIndex, instructions)
		accumulator = 0
		index = 0
		accumulator = 0
		index = 0

		for i := range instructions {
			instructions[i].executed = false
		}
	}
}

func parseInstructions(input []string) []instruction {
	instructions := []instruction{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Errorf("invalid instruction '%s'", line)
		}

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Errorf("could not parse value '%s' to int", parts[1])
		}

		var op operation
		switch parts[0] {
		case "acc":
			op = acc
			break
		case "jmp":
			op = jmp
			break
		case "nop":
			op = nop
			break
		default:
			fmt.Errorf("unknown instruction '%s'", parts[0])
		}

		instructions = append(instructions, instruction{
			op:       op,
			value:    value,
			executed: false,
		})
	}

	return instructions
}

func findNextJmpOrNop(start int, instructions []instruction) int {
	for i := start + 1; i < len(instructions); i++ {
		if instructions[i].op == jmp || instructions[i].op == nop {
			return i
		}
	}

	return -1
}
