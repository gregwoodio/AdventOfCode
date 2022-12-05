package day05

import (
	"regexp"
	"strconv"
	"strings"
)

var regex = `move (?P<Count>[0-9]+) from (?P<Source>[0-9]+) to (?P<Target>[0-9]+)`

func solvePartOne(input []string) string {
	return solve(input, false)
}

func solvePartTwo(input []string) string {
	return solve(input, true)
}

func solve(input []string, isPartTwo bool) string {
	stacks, instructions := parseInput(input)
	compiled := regexp.MustCompile(regex)

	for _, inst := range instructions {
		match := compiled.FindStringSubmatch(inst)
		var count, source, target int

		for i, subexpName := range compiled.SubexpNames() {
			if i == 0 || subexpName == "" {
				continue
			}

			switch subexpName {
			case "Count":
				count, _ = strconv.Atoi(match[i])
			case "Source":
				source, _ = strconv.Atoi(match[i])
			case "Target":
				target, _ = strconv.Atoi(match[i])
			}
		}

		if !isPartTwo {
			for j := 0; j < count; j++ {
				// pop from source
				val := stacks[source-1][len(stacks[source-1])-1]
				stacks[source-1] = stacks[source-1][:len(stacks[source-1])-1]

				// push to target
				stacks[target-1] = append(stacks[target-1], val)
			}
		} else {
			// pop from source
			val := stacks[source-1][len(stacks[source-1])-count:]
			stacks[source-1] = stacks[source-1][:len(stacks[source-1])-count]

			// push to target
			stacks[target-1] = append(stacks[target-1], val...)

		}
	}

	// top value of each stack
	output := []byte{}
	for _, stack := range stacks {
		output = append(output, stack[len(stack)-1])
	}

	return string(output)
}

func parseInput(input []string) ([][]byte, []string) {

	stackLines := []string{}
	i := 0
	for {
		if input[i] == "" {
			break
		}

		if strings.Contains(input[i], "[") {
			stackLines = append([]string{input[i]}, stackLines...)
		}
		i++
	}

	stacks := [][]byte{}
	for _, line := range stackLines {
		stackIndex := 0
		for j := 1; j < len(line); j += 4 {
			if len(stacks) <= stackIndex {
				stacks = append(stacks, []byte{})
			}
			if line[j] != ' ' {
				stacks[stackIndex] = append(stacks[stackIndex], line[j])
			}
			stackIndex++
		}
	}

	return stacks, input[i+1:]
}
