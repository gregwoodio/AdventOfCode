package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gregwoodio/aocutil"
)

var regex = `(?P<TargetReg>[a-z]+) (?P<Inst>((inc)|(dec))) (?P<Increment>\-?[0-9]+) if (?P<CondReg>[a-z]+) (?P<Condition>((==)|(!=)|(\>)|(\>=)|(\<)|(\<=))) (?P<CondValue>\-?[0-9]+)`

func main() {
	input := aocutil.ReadStringsFromFile("input.txt")

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input []string) int {
	return solve(input, false)
}

func partTwo(input []string) int {
	return solve(input, true)
}

func solve(input []string, isPartTwo bool) int {
	registers := make(map[string]int)
	var maxValue int
	comp := regexp.MustCompile(regex)

	for _, line := range input {
		match := comp.FindStringSubmatch(line)

		var targetReg, condReg, inst, cond string
		var inc, condValue int
		var doCondition bool

		for i, subexpName := range comp.SubexpNames() {
			if i == 0 || subexpName == "" {
				continue
			}

			if i > len(match) {
				break
			}

			switch subexpName {
			case "TargetReg":
				targetReg = match[i]
			case "CondReg":
				condReg = match[i]
			case "Inst":
				inst = match[i]
			case "Condition":
				cond = match[i]
			case "Increment":
				inc64, _ := strconv.ParseInt(match[i], 10, 64)
				inc = int(inc64)
			case "CondValue":
				condValue64, _ := strconv.ParseInt(match[i], 10, 64)
				condValue = int(condValue64)
			}
		}

		switch cond {
		case ">":
			doCondition = registers[condReg] > condValue
		case ">=":
			doCondition = registers[condReg] >= condValue
		case "<":
			doCondition = registers[condReg] < condValue
		case "<=":
			doCondition = registers[condReg] <= condValue
		case "==":
			doCondition = registers[condReg] == condValue
		case "!=":
			doCondition = registers[condReg] != condValue
		}

		if doCondition {
			if inst == "inc" {
				registers[targetReg] += inc
			} else if inst == "dec" {
				registers[targetReg] -= inc
			}

			if isPartTwo && registers[targetReg] > maxValue {
				maxValue = registers[targetReg]
			}
		}
	}

	if isPartTwo {
		return maxValue
	}

	var max int
	for _, value := range registers {
		if value > max {
			max = value
		}
	}
	return max
}
