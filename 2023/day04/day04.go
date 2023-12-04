package day04

import (
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	sum := 0

	for _, line := range input {
		colon := strings.Index(line, ":")
		nums := strings.Split(line[colon+1:], "|")
		winningNums := map[string]bool{}
		score := 0
		for _, num := range strings.Split(nums[0], " ") {
			if len(num) != 0 {
				winningNums[num] = true
			}
		}
		for _, num := range strings.Split(nums[1], " ") {
			if _, ok := winningNums[num]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		sum += score

	}

	return sum
}

type scratchcard struct {
	id        int
	matches   int
	instances int
}

func solvePartTwo(input []string) int {

	cards := []scratchcard{}

	for _, line := range input {
		colon := strings.Index(line, ":")
		id, _ := strconv.Atoi(line[strings.Index(line, " ")+1 : colon])

		nums := strings.Split(line[colon+1:], "|")
		winningNums := map[string]bool{}
		for _, num := range strings.Split(nums[0], " ") {
			if len(num) != 0 {
				winningNums[num] = true
			}
		}
		matches := 0
		for _, num := range strings.Split(nums[1], " ") {
			if _, ok := winningNums[num]; ok {
				matches++
			}
		}

		cards = append(cards, scratchcard{
			id:        id,
			matches:   matches,
			instances: 1,
		})
	}

	sum := 0
	for i, card := range cards {
		for j := 1; j <= card.matches; j++ {
			cards[i+j].instances = cards[i+j].instances + card.instances
		}
		sum += card.instances
	}

	return sum
}
