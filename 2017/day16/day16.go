package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	programs := []byte("abcdefghijklmnop")

	partOneOrder := partOne(programs, string(input))
	fmt.Println(partOneOrder)
	fmt.Println(partTwo([]byte(partOneOrder), string(input)))
}

func partOne(programs []byte, input string) string {
	return dance(programs, input)
}

func partTwo(programs []byte, input string) string {
	prog := string(programs)
	solutions := make(map[string]int)
	var cycles int

	for i := 0; i < 1000000000; i++ {
		prog = dance([]byte(prog), input)
		if solutions[prog] > 0 {
			cycles = i - solutions[prog]
			break
		} else {
			solutions[prog] = i
		}
	}

	prog = string(programs)
	for i := 0; i < (1000000000%cycles)-1; i++ {
		prog = dance([]byte(prog), input)
	}

	return prog
}

func dance(programs []byte, input string) string {
	steps := strings.Split(input, ",")

	for _, step := range steps {
		switch step[0] {
		case 's':
			num, _ := strconv.Atoi(step[1:])
			programs = append(programs[len(programs)-num:], programs[:len(programs)-num]...)
		case 'x':
			positions := strings.Split(step[1:], "/")
			pos1, _ := strconv.Atoi(positions[0])
			pos2, _ := strconv.Atoi(positions[1])

			temp := programs[pos1]
			programs[pos1] = programs[pos2]
			programs[pos2] = temp
		case 'p':
			progs := strings.Split(step[1:], "/")
			pos1 := bytes.IndexAny(programs, progs[0])
			pos2 := bytes.IndexAny(programs, progs[1])

			temp := programs[pos1]
			programs[pos1] = programs[pos2]
			programs[pos2] = temp
		}
	}

	return string(programs)
}
