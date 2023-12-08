package day08

import "regexp"

type node struct {
	label string
	left  string
	right string
}

func solvePartOne(input []string) int {
	inst := input[0]

	nodes := map[string]node{}

	for i := 2; i < len(input); i++ {
		n := parseLine(input[i])
		nodes[n.label] = n
	}

	i := 0
	count := 0
	curr := nodes["AAA"]

	for {
		step := inst[i%len(inst)]

		if step == 'L' {
			curr = nodes[curr.left]
		} else if step == 'R' {
			curr = nodes[curr.right]
		}
		count++
		i++

		if curr.label == "ZZZ" {
			return count
		}
	}
}

func solvePartTwo(input []string) int {
	return -1
}

func parseLine(line string) node {
	regex := regexp.MustCompile(`(?P<Label>[0-9A-Z]{3}) = \((?P<Left>[0-9A-Z]{3}), (?P<Right>[0-9A-Z]{3})\)`)
	parts := regex.FindStringSubmatch(line)
	return node{
		label: parts[1],
		left:  parts[2],
		right: parts[3],
	}

}
