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

// func solvePartTwo(input string) int {
// 	target := 19690720
// 	ch := make(chan int)
// 	stop := make(chan bool)

// 	defer close(ch)
// 	defer close(stop)

// 	var result int

// 	for noun := 0; noun <= 99; noun++ {

// 		go func(noun int) {
// 			for verb := 0; verb <= 99; verb++ {
// 				select {
// 				default:
// 					inst := parseInput(input)

// 					inst[1] = noun
// 					inst[2] = verb

// 					if processIntcode(inst) == target {
// 						ch <- 100*noun + verb
// 					}
// 					break
// 				case <-stop:
// 					return
// 				}
// 			}
// 		}(noun)
// 	}

// 	result = <-ch
// 	stop <- true

// 	return result
// }

func processIntcode(inst []int, input io.Reader) int {
	ip := 0
	var isParam1Immediate, isParam2Immediate, isParam3Immediate bool
	reader := bufio.NewReader(input)

	for {
		operation := inst[ip] % 10

		if operation != 1 && operation != 2 && operation != 3 && operation != 4 {
			return inst[0]
		}

		isParam1Immediate = (inst[ip]/100)%10 == 1

		var p1, p2, p3 *int
		if isParam1Immediate {
			p1 = &inst[ip+1]
		} else {
			p1 = &inst[inst[ip+1]]
		}

		if operation == 1 || operation == 2 {
			isParam2Immediate = (inst[ip]/1000)%10 == 1
			isParam3Immediate = false //(inst[ip]/10000)%10 == 1

			if isParam2Immediate {
				p2 = &inst[ip+2]
			} else {
				p2 = &inst[inst[ip+2]]
			}

			if isParam3Immediate {
				p3 = &inst[ip+3]
			} else {
				p3 = &inst[inst[ip+3]]
			}
		} else {
			isParam2Immediate = false
			isParam3Immediate = false
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
