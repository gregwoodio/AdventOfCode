package day19

import (
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	x, m, a, s int
}

func (p part) rating() int {
	return p.x + p.m + p.a + p.s
}

var partRegex string = `{x=(?P<X>\d+),m=(?P<M>\d+),a=(?P<A>\d+),s=(?P<S>\d+)}`

func parsePart(input string) part {
	r := regexp.MustCompile(partRegex)
	matches := r.FindStringSubmatch(input)
	x, _ := strconv.Atoi(matches[1])
	m, _ := strconv.Atoi(matches[2])
	a, _ := strconv.Atoi(matches[3])
	s, _ := strconv.Atoi(matches[4])
	return part{x, m, a, s}
}

type workflowStep struct {
	// function to process the part
	fn func(p part) bool

	// destination workflow if successful
	dest string
}

type workflow struct {
	name  string
	steps []workflowStep
}

func parseWorkflow(input string) workflow {
	i := strings.Index(input, "{")
	name := input[:i]
	stepValues := strings.Split(input[i+1:len(input)-1], ",")
	steps := []workflowStep{}

	for _, stepValue := range stepValues {
		i = strings.Index(stepValue, ":")
		if i == -1 {
			// no function, just a destination
			steps = append(steps, workflowStep{
				fn:   func(p part) bool { return true },
				dest: stepValue,
			})
			continue
		}

		var compare func(int, int) bool
		if stepValue[1] == '<' {
			compare = func(a, b int) bool { return a < b }
		} else {
			compare = func(a, b int) bool { return a > b }
		}

		value, _ := strconv.Atoi(stepValue[2:i])
		var fn func(p part) bool
		if stepValue[0] == 'x' {
			fn = func(p part) bool { return compare(p.x, value) }
		} else if stepValue[0] == 'm' {
			fn = func(p part) bool { return compare(p.m, value) }
		} else if stepValue[0] == 'a' {
			fn = func(p part) bool { return compare(p.a, value) }
		} else if stepValue[0] == 's' {
			fn = func(p part) bool { return compare(p.s, value) }
		}

		steps = append(steps, workflowStep{
			fn:   fn,
			dest: stepValue[i+1:],
		})
	}

	return workflow{
		name:  name,
		steps: steps,
	}
}

func solvePartOne(input []string) int {
	workflows := map[string]workflow{}
	i := 0
	for ; i < len(input); i++ {
		if input[i] == "" {
			break
		}
		workflow := parseWorkflow(input[i])
		workflows[workflow.name] = workflow
	}

	// skip blank line
	i++

	parts := []part{}
	for ; i < len(input); i++ {
		parts = append(parts, parsePart(input[i]))
	}

	accept := []part{}

	for _, p := range parts {
		curr := "in"

		for {
			workflow := workflows[curr]
			for _, step := range workflow.steps {
				if step.fn(p) {
					curr = step.dest
					break
				}
			}

			if curr == "A" {
				accept = append(accept, p)
				break
			} else if curr == "R" {
				break
			}
		}
	}

	sum := 0

	for _, part := range accept {
		sum += part.rating()
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}
