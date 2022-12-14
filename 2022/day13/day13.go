package day13

import (
	"strconv"
	"strings"
)

type node struct {
	parent *node
	nodes  []*node
	value  *int
}

func solvePartOne(input []string) int {
	sum := 0

	index := 1
	for i := 0; i < len(input); i += 3 {
		res := compare(input[i], input[i+1])
		if res != -1 {
			sum += index
		}
		index++
	}

	return sum
}

func solvePartTwo(input []string) int {
	list := []string{"[[2]]", "[[6]]"}

	// filter out empty strings
	for _, s := range input {
		if len(s) != 0 {
			list = append(list, s)
		}
	}

	// bubble sort ftw
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(list)-1; i++ {
			if compare(list[i], list[i+1]) == -1 {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
	}

	div1 := -1
	div2 := -1
	for i, s := range list {
		if s == "[[2]]" {
			div1 = i + 1
		}
		if s == "[[6]]" {
			div2 = i + 1
		}
	}

	return div1 * div2
}

func compareBool(left, right string) bool {
	return compare(left, right) != -1
}

func compare(left, right string) int {
	if right == "" {
		if left == "" {
			return 0
		} else {
			return -1
		}
	}
	if left == "" {
		return 1
	}

	if !strings.HasPrefix(left, "[") && !strings.HasPrefix(right, "[") {
		n1, err := strconv.Atoi(left)
		if err != nil {
			panic(err.Error())
		}
		n2, err := strconv.Atoi(right)
		if err != nil {
			panic(err.Error())
		}
		if n1 < n2 {
			return 1
		} else if n1 == n2 {
			return 0
		} else {
			return -1
		}
	}

	if strings.HasPrefix(left, "[") && strings.HasPrefix(right, "[") {
		leftParts := unwrap(left)
		rightParts := unwrap(right)
		last := 0
		for i := 0; true; i++ {
			if i == len(rightParts) && i == len(leftParts) {
				return 0
			}
			if i >= len(rightParts) {
				return -1
			}
			if i >= len(leftParts) {
				return 1
			}
			last = compare(leftParts[i], rightParts[i])
			if last == 0 {
				continue
			}
			return last
		}
	}

	if !strings.HasPrefix(left, "[") {
		left = "[" + left + "]"
		return compare(left, right)
	}
	if !strings.HasPrefix(right, "]") {
		right = "[" + right + "]"
		return compare(left, right)
	}

	panic("Every case should be handled, but here we are")
}

func unwrap(s string) []string {
	parts := []string{}
	s = strings.TrimSuffix(strings.TrimPrefix(s, "["), "]")
	depth := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			depth++
			continue
		}
		if s[i] == ']' {
			depth--
			continue
		}

		if s[i] == ',' {
			if depth != 0 {
				continue
			}
			parts = append(parts, s[start:i])
			start = i + 1
		}
	}

	parts = append(parts, s[start:])
	return parts
}
