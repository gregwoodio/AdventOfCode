package day08

import (
	"fmt"
	"strconv"
)

func solvePartOne(input []string) int {
	trees := buildTrees(input)
	count := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			visible := true

			// left
			for k := j - 1; k >= 0; k-- {
				if trees[i][k] >= trees[i][j] {
					visible = false
					break
				}
			}

			if visible {
				count++
			}

			// right
			if !visible {
				visible = true
				for k := j + 1; k < len(trees[i]); k++ {
					if trees[i][k] >= trees[i][j] {
						visible = false
						break
					}
				}
				if visible {
					count++
				}
			}

			// top
			if !visible {
				visible = true
				for k := i - 1; k >= 0; k-- {
					if trees[k][j] >= trees[i][j] {
						visible = false
						break
					}
				}
				if visible {
					count++
				}
			}

			// bottom
			if !visible {
				visible = true
				for k := i + 1; k < len(trees); k++ {
					if trees[k][j] >= trees[i][j] {
						visible = false
						break
					}
				}
				if visible {
					count++
				}
			}
		}
	}

	return count
}

func solvePartTwo(input []string) int {
	trees := buildTrees(input)
	max := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {

			left := 0
			for k := j - 1; k >= 0; k-- {
				left++
				if trees[i][k] >= trees[i][j] {
					break
				}
			}

			right := 0
			for k := j + 1; k < len(trees[i]); k++ {
				right++
				if trees[i][k] >= trees[i][j] {
					break
				}
			}

			top := 0
			for k := i - 1; k >= 0; k-- {
				top++
				if trees[k][j] >= trees[i][j] {
					break
				}
			}

			bottom := 0
			for k := i + 1; k < len(trees); k++ {
				bottom++
				if trees[k][j] >= trees[i][j] {
					break
				}
			}

			score := left * right * top * bottom
			if score > max {
				max = score
			}
		}
	}

	return max
}

func buildTrees(input []string) [][]int {
	trees := [][]int{}

	for _, line := range input {
		treeLine := []int{}
		for i := range line {
			height, err := strconv.Atoi(line[i : i+1])
			if err != nil {
				fmt.Printf("Could not parse '%s' as int", line[i:i+1])
			}
			treeLine = append(treeLine, height)
		}
		trees = append(trees, treeLine)
	}

	return trees
}
