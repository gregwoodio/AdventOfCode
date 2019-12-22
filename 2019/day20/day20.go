package day20

import (
	"container/heap"
	"fmt"
	"math"
	"unicode"
)

type node struct {
	id     string
	paths  []*path
	weight int
	index  int
	parent *node
}

type path struct {
	distance int
	node     *node
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
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(n interface{}) {
	index := len(*pq)
	item := n.(*node)
	item.index = index
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	index := len(old) - 1
	node := old[index]
	old[index] = nil
	node.index = -1
	*pq = old[:index]

	return node
}

func solvePartOne(input []string) int {

	start, end, allNodes := makeMaze(input)

	dijkstra(start, allNodes)

	return end.weight
}

func dijkstra(start *node, nodes []*node) {
	// add all nodes to my priority queue
	start.weight = 0
	pq := priorityQueue{}
	for i := range nodes {
		pq.Push(nodes[i])
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*node)
		// if node.weight == math.MaxInt32 {
		// 	// there's no path to other nodes from here
		// 	break
		// }

		for i := range node.paths {
			if node.paths[i].distance+node.weight < node.paths[i].node.weight {
				node.paths[i].node.weight = node.paths[i].distance + node.weight
				node.paths[i].node.parent = node
			}
		}

		heap.Init(&pq)
	}
}

// Makes a graph of the maze, returns start and end nodes
func makeMaze(input []string) (*node, *node, []*node) {

	// make temp array of all nodes known.
	// we're assuming here that all mazes will be square
	coords := make([][]*node, len(input[0]))
	for i := range coords {
		coords[i] = make([]*node, len(input[0]))
	}
	var nodeCount int
	warps := make(map[string]*node)

	allNodes := []*node{}

	// go through horizontally and create nodes. Each map
	// in the examples has a buffer of 2 spaces on any side
	// for warp names, so start at 2.
	for row := 2; row < len(input)-2; row++ {
		wasWall := false
		var prev, curr *node
		var weight int

		for col := 2; col < len(input[row])-2; col++ {
			if input[row][col] == '#' {
				wasWall = true
				prev = nil
				weight = 0
				continue
			}

			if input[row][col] == '.' {
				// if row == 2 {
				// 	// node on top edge must be a warp
				// 	curr = newNode(string([]byte{input[0][col], input[1][col]}))
				// 	coords[row][col] = curr
				// 	allNodes = append(allNodes, curr)
				// 	nodeCount++

				// } else if row == len(input)-3 {
				// 	// node on bottom edge must be a warp
				// 	curr = newNode(string([]byte{input[len(input)-2][col], input[len(input)-2][col]}))

				// 	coords[row][col] = curr
				// 	allNodes = append(allNodes, curr)
				// 	nodeCount++

				// } else if col == 2 {
				// 	// it's on the edge, so must be a warp
				// 	curr = newNode(string(input[row][:2]))
				// 	coords[row][col] = curr
				// 	allNodes = append(allNodes, curr)
				// 	nodeCount++

				// } else if col == len(input[row])-3 {
				// 	// it's on the opposite edge, so must be a warp
				// 	curr = newNode(string(input[row][len(input[row])-2:]))
				// 	coords[row][col] = curr
				// 	allNodes = append(allNodes, curr)
				// 	nodeCount++

				if unicode.IsLetter(rune(input[row+1][col])) ||
					unicode.IsLetter(rune(input[row-1][col])) ||
					unicode.IsLetter(rune(input[row][col-1])) ||
					unicode.IsLetter(rune(input[row][col+1])) ||
					input[row+1][col] == '.' ||
					input[row-1][col] == '.' {

					// Is there a warp on the inside of the donut?
					key := fmt.Sprintf("%d,%d", row, col)
					if unicode.IsLetter(rune(input[row+1][col])) {
						key = string([]byte{input[row+1][col], input[row+2][col]})
					} else if unicode.IsLetter(rune(input[row-1][col])) {
						key = string([]byte{input[row-2][col], input[row-1][col]})
					} else if unicode.IsLetter(rune(input[row][col-1])) {
						key = string([]byte{input[row][col-2], input[row][col-1]})
					} else if unicode.IsLetter(rune(input[row][col+1])) {
						key = string([]byte{input[row][col+1], input[row][col+2]})
					}

					curr = newNode(key)
					coords[row][col] = curr
					allNodes = append(allNodes, curr)
					nodeCount++
					wasWall = false

				} else if wasWall {
					if input[row-1][col] == '.' && input[row+1][col] == '.' &&
						input[row][col-1] == '#' && input[row][col+1] == '#' {
						// This point is on a vertical path, it should not be a node
						continue
					}

					// Is there a warp on the inside of the donut?
					key := fmt.Sprintf("%d,%d", row, col)
					// if unicode.IsLetter(rune(input[row+1][col])) {
					// 	key = string([]byte{input[row+1][col], input[row+2][col]})
					// } else if unicode.IsLetter(rune(input[row-1][col])) {
					// 	key = string([]byte{input[row-2][col], input[row-1][col]})
					// } else if unicode.IsLetter(rune(input[row][col-1])) {
					// 	key = string([]byte{input[row][col-2], input[row][col-1]})
					// } else if unicode.IsLetter(rune(input[row][col+1])) {
					// 	key = string([]byte{input[row][col+1], input[row][col+2]})
					// }

					curr = newNode(key)
					coords[row][col] = curr
					allNodes = append(allNodes, curr)
					nodeCount++
					wasWall = false

				} else if input[row][col+1] == '#' {
					// is the next block a wall?
					curr = newNode(fmt.Sprintf("%d,%d", row, col))

					coords[row][col] = curr
					allNodes = append(allNodes, curr)
					nodeCount++
					// } else if unicode.IsLetter(rune(input[row][col+1])) {
					// 	// is the next block a?
					// 	key := string([]byte{input[row][col+1], input[row][col+2]})
					// 	curr = newNode(key)

					// 	coords[row][col] = curr
					// 	allNodes = append(allNodes, curr)
					// 	nodeCount++
				}

				if curr != nil {

					if prev != nil {
						prev.paths = append(prev.paths, &path{
							distance: weight,
							node:     curr,
						})

						curr.paths = append(curr.paths, &path{
							distance: weight,
							node:     prev,
						})

						weight = 0
					}

					// connect or add warp point
					if unicode.IsLetter(rune(curr.id[0])) {
						if warp, ok := warps[curr.id]; ok {
							warp.paths = append(warp.paths, &path{
								distance: 1,
								node:     curr,
							})

							curr.paths = append(curr.paths, &path{
								distance: 1,
								node:     warp,
							})
						} else {
							warps[curr.id] = curr
						}
					}

					prev = curr
					curr = nil
				}
				weight++
			}
		}
	}

	// At this point we should have all the nodes.
	// Now go through vertically and connect them.
	for col := 2; col < len(input[0])-2; col++ {
		var curr, prev *node
		var weight int

		for row := 2; row < len(input)-2; row++ {
			if input[row][col] == '#' || unicode.IsLetter(rune(input[row][col])) {
				prev = nil
				weight = 0
				continue
			}

			if coords[row][col] != nil {
				curr = coords[row][col]

				if prev != nil {
					curr.paths = append(curr.paths, &path{
						distance: weight,
						node:     prev,
					})
					prev.paths = append(prev.paths, &path{
						distance: weight,
						node:     curr,
					})

					weight = 0
				}

				prev = curr
			}

			weight++
		}
	}

	fmt.Printf("Found %d nodes\n", nodeCount)
	for _, val := range allNodes {
		fmt.Println(val.id)
	}
	return warps["AA"], warps["ZZ"], allNodes
}

func newNode(id string) *node {
	return &node{
		id:     id,
		paths:  []*path{},
		weight: math.MaxInt32,
	}
}
