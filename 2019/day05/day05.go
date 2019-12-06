package day05

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func solvePartOne(input string, reader io.Reader) int {
	inst := parseInput(input)
	return processIntcode(inst, reader)
}

func solvePartTwo(input string, reader io.Reader) int {
	inst := parseInput(input)
	return processIntcode(inst, reader)
}

func processIntcode(inst []int, input io.Reader) int {
	ip := 0
	var isParam1Immediate, isParam2Immediate, isParam3Immediate bool
	reader := bufio.NewReader(input)

	for {
		operation := inst[ip] % 10

		if operation < 1 || operation > 8 {
			return inst[0]
		}

		isParam1Immediate = (inst[ip]/100)%10 == 1

		var p1, p2, p3 *int
		if isParam1Immediate {
			p1 = &inst[ip+1]
		} else {
			p1 = &inst[inst[ip+1]]
		}

		if operation == 3 || operation == 4 {
			isParam2Immediate = false
			isParam3Immediate = false
		} else {
			isParam2Immediate = (inst[ip]/1000)%10 == 1
			isParam3Immediate = false //(inst[ip]/10000)%10 == 1

			if isParam2Immediate {
				p2 = &inst[ip+2]
			} else {
				p2 = &inst[inst[ip+2]]
			}

			if operation != 5 && operation != 6 {
				if isParam3Immediate {
					p3 = &inst[ip+3]
				} else {
					p3 = &inst[inst[ip+3]]
				}
			}
		}

		if operation == 1 {
			// add
			*p3 = *p1 + *p2
			ip += 4
		} else if operation == 2 {
			// multiply
			*p3 = *p1 * *p2
			ip += 4
		} else if operation == 3 {
			// input
			fmt.Println("Input integer: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			text = text[0:strings.Index(text, "\n")]
			val, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}

			// Parameters that an instruction writes to will never be immediate
			p1 = &inst[inst[ip+1]]

			*p1 = val
			ip += 2
		} else if operation == 4 {
			// output
			fmt.Printf("Output: %d\n", *p1)
			ip += 2
		} else if operation == 5 {
			// jump if true
			if *p1 != 0 {
				ip = *p2
			} else {
				ip += 3
			}
		} else if operation == 6 {
			// jump if false
			if *p1 == 0 {
				ip = *p2
			} else {
				ip += 3
			}
		} else if operation == 7 {
			// less than
			if *p1 < *p2 {
				*p3 = 1
			} else {
				*p3 = 0
			}
			ip += 4
		} else if operation == 8 {
			// equals
			if *p1 == *p2 {
				*p3 = 1
			} else {
				*p3 = 0
			}
			ip += 4
		}

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
