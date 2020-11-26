package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gregwoodio/aocutil"
)

var regex = `(?P<Name>[a-z]+).*\((?P<Weight>[0-9]+)\) \-\> (?P<Children>.*)`
var regexNoChildren = `(?P<Name>[a-z]+).*\((?P<Weight>[0-9]+)`

type treeNode struct {
	Name             string
	Weight           int
	TotalWeight      int
	Children         []*treeNode
	ChildrenUnparsed string
	Parent           *treeNode
}

func main() {
	input := aocutil.ReadStringsFromFile("input.txt")
	head := buildTree(input)
	fmt.Printf("%s\n", partOne(head))
	fmt.Printf("%d\n", partTwo(head))
}

func partOne(tree *treeNode) string {
	return tree.Name
}

func partTwo(tree *treeNode) int {
	findWeights(tree)
	// printBalances(tree)
	parentNode := checkBalances(tree)

	// find most repeated value
	counts := make(map[int]int)
	for _, child := range parentNode.Children {
		counts[child.TotalWeight]++
	}
	max, count := 0, 0
	for key := range counts {
		if counts[key] > count {
			count = counts[key]
			max = key
		}
	}

	// find child where totalweight does not match max value
	for _, child := range parentNode.Children {
		if child.TotalWeight != max {
			return child.Weight - (child.TotalWeight - max)
		}
	}
	return -1
}

func buildTree(input []string) *treeNode {
	nodes := make(map[string]*treeNode)

	// parse lines into map
	for _, line := range input {

		var compiled *regexp.Regexp
		var name, children string
		var weight int

		if strings.Contains(line, "->") {
			compiled = regexp.MustCompile(regex)
		} else {
			compiled = regexp.MustCompile(regexNoChildren)
		}

		match := compiled.FindStringSubmatch(line)

		for i, subexpName := range compiled.SubexpNames() {
			if i == 0 || subexpName == "" {
				continue
			}

			if i > len(match) {
				break
			}

			switch subexpName {
			case "Weight":
				weight64, _ := strconv.ParseInt(match[i], 10, 64)
				weight = int(weight64)
			case "Name":
				name = match[i]
			case "Children":
				children = match[i]
			}
		}

		node := treeNode{
			Name:             name,
			Weight:           weight,
			ChildrenUnparsed: children,
		}

		nodes[node.Name] = &node
	}

	// build relationships between nodes
	for _, node := range nodes {
		if node.ChildrenUnparsed == "" {
			continue
		}

		children := strings.Split(node.ChildrenUnparsed, ", ")

		for _, child := range children {
			childNode := nodes[child]
			node.Children = append(node.Children, childNode)
			childNode.Parent = node
		}
	}

	for key := range nodes {
		head := nodes[key]
		for {
			if head.Parent == nil {
				return head
			}
			head = head.Parent
		}
	}

	return nil
}

func isBalanced(node *treeNode) bool {
	if len(node.Children) == 0 {
		return true
	}

	var weight int
	for i, child := range node.Children {
		if i == 0 {
			weight = child.TotalWeight
		}

		if child.TotalWeight != weight {
			return false
		}
	}
	return true
}

func findWeights(node *treeNode) int {
	node.TotalWeight = node.Weight

	for _, child := range node.Children {
		node.TotalWeight += findWeights(child)
	}

	return node.TotalWeight
}

func checkBalances(node *treeNode) *treeNode {

	for !isBalanced(node) {
		unbalanceChild := -1
		for i, child := range node.Children {
			if !isBalanced(child) {
				unbalanceChild = i
			}
		}

		if unbalanceChild == -1 {
			return node
		}
		node = node.Children[unbalanceChild]
	}
	return nil
}

func printBalances(node *treeNode) {
	if !isBalanced(node) {
		fmt.Printf("%s is not balanced. Weight is %d, total weight is %d\n\tChildren:", node.Name, node.Weight, node.TotalWeight)
		for _, child := range node.Children {
			fmt.Printf("%s (%d), ", child.Name, child.TotalWeight)
		}

		fmt.Println()
		for _, child := range node.Children {
			printBalances(child)
		}
	}
}
