package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func main() {
	input := aocutil.ReadStringsFromFile("input.txt")
	fmt.Println(partOne(input))
}

func partOne(input []string) int {
	registers := make(map[string]int)
	sound := 0
	step := 0

	for {
		inst := input[step]

		parts := strings.Split(inst, " ")

		var val int
		if len(parts) > 2 {
			if isLetter(parts[2]) {
				val = registers[parts[2]]
			} else {
				val, _ = strconv.Atoi(parts[2])
			}
		}

		switch parts[0] {
		case "snd":
			if isLetter(parts[1]) {
				sound = registers[parts[1]]
			} else {
				sound, _ = strconv.Atoi(parts[1])
			}
			step++
		case "set":
			registers[parts[1]] = val
			step++
		case "add":
			registers[parts[1]] += val
			step++
		case "mul":
			registers[parts[1]] *= val
			step++
		case "mod":
			registers[parts[1]] %= val
			step++
		case "rcv":
			if registers[parts[1]] != 0 {
				return sound
			}
			step++
		case "jgz":
			if registers[parts[1]] > 0 {
				step += val
			} else {
				step++
			}
		}

		if step >= len(input) || step < 0 {
			break
		}
	}
	return -1
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
