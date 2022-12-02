package day02

import "strings"

func solvePartOne(input []string) int {
	score := 0
	for _, game := range input {
		choices := strings.Split(game, " ")
		if choices[1] == "X" {
			score += 1

		} else if choices[1] == "Y" {
			score += 2
		} else if choices[1] == "Z" {
			score += 3
		}

		// win
		if choices[1] == "X" && choices[0] == "C" ||
			choices[1] == "Y" && choices[0] == "A" ||
			choices[1] == "Z" && choices[0] == "B" {
			score += 6
		}
		// draw
		if choices[1] == "X" && choices[0] == "A" ||
			choices[1] == "Y" && choices[0] == "B" ||
			choices[1] == "Z" && choices[0] == "C" {
			score += 3
		}

	}
	return score
}

func solvePartTwo(input []string) int {
	score := 0
	for _, game := range input {
		choices := strings.Split(game, " ")

		if choices[1] == "X" {
			// lose
			if choices[0] == "A" {
				//scissors
				score += 3
			} else if choices[0] == "B" {
				//rock
				score += 1

			} else if choices[0] == "C" {
				score += 2
			}

		} else if choices[1] == "Y" {
			// draw

			score += 3
			if choices[0] == "A" {
				score += 1
			} else if choices[0] == "B" {
				score += 2
			} else if choices[0] == "C" {
				score += 3
			}
		} else if choices[1] == "Z" {
			// win
			score += 6
			if choices[0] == "A" {
				score += 2
			} else if choices[0] == "B" {
				score += 3
			} else if choices[0] == "C" {
				score += 1
			}
		}

	}
	return score
}
