package day08

import (
	"regexp"
)

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

	return traverse(nodes, inst, "AAA", "ZZZ")
}

func solvePartTwo(input []string) int {
	inst := input[0]

	nodes := map[string]node{}
	starts := []string{}

	for i := 2; i < len(input); i++ {
		n := parseLine(input[i])
		nodes[n.label] = n

		if n.label[len(n.label)-1] == 'A' {
			starts = append(starts, n.label)
		}
	}

	counts := []int{}

	for _, start := range starts {
		runCounts := traversePart2(nodes, inst, start)
		counts = append(counts, runCounts...)
	}

	// find least common multiple
	leastCommon := lcm(counts[0], counts[1])
	for i := 2; i < len(counts); i++ {
		leastCommon = lcm(leastCommon, counts[i])
	}

	return leastCommon
}

func traverse(nodes map[string]node, inst, start, end string) int {
	i := 0
	count := 0
	curr := nodes[start]
	var next node

	for {
		step := inst[i%len(inst)]

		if step == 'L' {
			next = nodes[curr.left]
		} else if step == 'R' {
			next = nodes[curr.right]
		}
		count++
		i++

		if next.label == curr.label {
			return -1
		}

		curr = next

		if curr.label == end {
			return count
		}
	}
}

func traversePart2(nodes map[string]node, inst, start string) []int {
	i := 0
	count := 0
	curr := nodes[start]
	var next node
	endsWithZ := []int{}

	for {
		step := inst[i%len(inst)]

		if step == 'L' {
			next = nodes[curr.left]
		} else if step == 'R' {
			next = nodes[curr.right]
		}
		count++
		i++

		if next.label == curr.label || count == len(inst)*1000 {
			return endsWithZ
		}

		curr = next

		if curr.label[len(curr.label)-1] == 'Z' {
			if len(endsWithZ) > 0 && endsWithZ[len(endsWithZ)-1] == count/2 {
				// found a loop
				return endsWithZ
			}

			endsWithZ = append(endsWithZ, count)
		}
	}
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

// / Find the least common multiple of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// / Find the greatest common divisor of two numbers
func gcd(a, b int) int {
	var remainder int
	for a%b > 0 {
		remainder = a % b
		a = b
		b = remainder
	}

	return b
}
