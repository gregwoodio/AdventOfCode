package day12

import (
	"container/heap"
)

type node struct {
	elevation rune
	paths     []*node
	weight    int
}

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(n interface{}) {
	item := n.(*node)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	index := len(old) - 1
	node := old[index]
	old[index] = nil
	*pq = old[:index]

	return node
}

func solvePartOne(input []string) int {
	start, end, nodes := makeGraph(input)

	dijkstra(start, nodes)
	return end.weight
}

func solvePartTwo(input []string) int {
	_, end, nodes := makeGraph(input)

	as := []*node{}
	for _, n := range nodes {
		if n.elevation == 'a' {
			as = append(as, n)
		}
	}

	lowest := 9999
	for _, a := range as {
		a.weight = 0
		dijkstra(a, nodes)

		if end.weight < lowest {
			lowest = end.weight
		}

		// reset

		for _, n := range nodes {
			n.weight = 9999
		}
	}

	return lowest
}

func makeGraph(input []string) (*node, *node, []*node) {
	var start, end, n *node

	graph := [][]*node{}
	allNodes := []*node{}
	for i, line := range input {
		graph = append(graph, []*node{})

		for j := 0; j < len(line); j++ {
			ch := line[j]
			if ch == 'S' {
				n = &node{
					elevation: 'a',
					paths:     []*node{},
					weight:    0,
				}
				start = n
			} else if ch == 'E' {
				n = &node{
					elevation: 'z',
					paths:     []*node{},
					weight:    99999,
				}
				end = n
			} else {
				n = &node{
					elevation: rune(ch),
					paths:     []*node{},
					weight:    99999,
				}
			}

			graph[i] = append(graph[i], n)
			allNodes = append(allNodes, n)

			// check up
			if i > 0 {
				up := graph[i-1][j]
				if up.elevation == n.elevation {
					n.paths = append(n.paths, up)
					up.paths = append(up.paths, n)
				} else if up.elevation == n.elevation+1 {
					n.paths = append(n.paths, up)
				} else if up.elevation+1 == n.elevation {
					up.paths = append(up.paths, n)
				}
				if up.elevation > n.elevation {
					up.paths = append(up.paths, n)
				}
				if up.elevation < n.elevation {
					n.paths = append(n.paths, up)
				}
			}

			// check left
			if j > 0 {
				left := graph[i][j-1]
				if left.elevation == n.elevation {
					n.paths = append(n.paths, left)
					left.paths = append(left.paths, n)
				} else if left.elevation == n.elevation+1 {
					n.paths = append(n.paths, left)
				} else if left.elevation+1 == n.elevation {
					left.paths = append(left.paths, n)
				}
				if left.elevation > n.elevation {
					left.paths = append(left.paths, n)
				}
				if left.elevation < n.elevation {
					n.paths = append(n.paths, left)
				}
			}
		}
	}

	return start, end, allNodes
}

func dijkstra(start *node, nodes []*node) {
	pq := priorityQueue{}

	for _, n := range nodes {
		pq.Push(n)
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*node)

		for i := range node.paths {
			if node.paths[i].weight > node.weight+1 {
				node.paths[i].weight = node.weight + 1
			}
		}

		heap.Init(&pq)
	}
}
