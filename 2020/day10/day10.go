package day10

import "sort"

type numbers []int

func (n numbers) Len() int {
	return len(n)
}

func (n numbers) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func solvePartOne(input numbers) int {
	input = append(input, 0)
	sort.Sort(input)

	// add three more than highest value
	input = append(input, input[len(input)-1]+3)

	ones, threes, prev := 0, 0, input[0]
	for i := 1; i < len(input); i++ {
		if input[i]-prev == 1 {
			ones++
		} else if input[i]-prev == 3 {
			threes++
		}

		prev = input[i]
	}

	return ones * threes
}

type node struct {
	value   int
	visited int
	next    []*node
}

func solvePartTwo(input numbers) int {
	input = append(input, 0)
	sort.Sort(input)

	// add three more than highest value
	input = append(input, input[len(input)-1]+3)

	// make tree
	nodes := []*node{}
	for _, val := range input {
		nodes = append(nodes, &node{
			value: val,
			next:  []*node{},
		})
	}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < i+5 && j < len(nodes); j++ {
			if nodes[j].value <= nodes[i].value+3 {
				nodes[i].next = append(nodes[i].next, nodes[j])
			} else {
				break
			}
		}
	}

	queue := []*node{
		nodes[0],
	}

	for {
		if len(queue) == 0 {
			break
		}

		n := queue[0]
		n.visited++
		queue = queue[1:]

		for _, next := range n.next {
			queue = append(queue, next)
		}
	}

	return nodes[len(nodes)-1].visited
}
