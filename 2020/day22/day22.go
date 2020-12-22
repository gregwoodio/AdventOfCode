package day22

import (
	"fmt"
	"strconv"
	"strings"
)

type player struct {
	id   string
	deck []int
}

func (p *player) getScore() int {
	score := 0
	points := 1

	for i := len(p.deck) - 1; i >= 0; i-- {
		score += points * p.deck[i]
		points++
	}

	return score
}

func (p *player) playTopCard(debug bool) int {
	if debug {
		deck := []string{}
		for _, card := range p.deck {
			cardStr := strconv.Itoa(card)
			deck = append(deck, cardStr)
		}
		fmt.Printf("%s's deck: %s\n", p.id, strings.Join(deck, ","))
	}

	card := p.deck[0]
	p.deck = p.deck[1:]
	return card
}

func solvePartOne(input []string, debug bool) int {
	players := setup(input)
	round := 1

	for {
		// check win
		if len(players[0].deck) == 0 {
			return players[1].getScore()
		} else if len(players[1].deck) == 0 {
			return players[0].getScore()
		}

		if debug {
			fmt.Printf("-- Round %d --\n", round)
		}

		p1 := players[0].playTopCard(debug)
		p2 := players[1].playTopCard(debug)

		if debug {
			fmt.Printf("%s plays: %d\n", players[0].id, p1)
			fmt.Printf("%s plays: %d\n", players[1].id, p2)
		}

		if p1 > p2 {
			players[0].deck = append(players[0].deck, p1, p2)

			if debug {
				fmt.Printf("%s wins the round!\n\n", players[0].id)
			}
		} else {
			players[1].deck = append(players[1].deck, p2, p1)

			if debug {
				fmt.Printf("%s wins the round!\n\n", players[1].id)
			}
		}

		round++
	}
}

func solvePartTwo(input []string) int {
	return -1
}

func setup(input []string) []player {
	players := []player{}
	var p player

	for _, line := range input {
		if line == "" {
			players = append(players, p)
		} else if strings.HasPrefix(line, "Player") {
			p = player{
				id:   line[:len(line)-1],
				deck: []int{},
			}
		} else {
			card, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("could not parse '%s' to int", line)
			}

			p.deck = append(p.deck, card)
		}
	}

	return players
}
