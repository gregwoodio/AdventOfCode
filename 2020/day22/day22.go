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

func (p *player) getState() string {
	cards := []string{}
	for _, c := range p.deck {
		cards = append(cards, strconv.Itoa(c))
	}

	return fmt.Sprintf("%s:%s;", p.id, strings.Join(cards, ","))
}

func solvePartOne(input []string, debug bool) int {
	players := setup(input)
	score, _ := play(players, 1, false, debug)
	return score
}

func solvePartTwo(input []string, debug bool) int {
	players := setup(input)
	score, _ := play(players, 1, true, debug)
	return score
}

// returns score, winner
func play(players []player, game int, isPartTwo, debug bool) (int, int) {
	round := 1

	gameState := make(map[string]bool)

	if debug {
		fmt.Printf("=== Game %d ===\n\n", game)
	}

	for {
		// check win
		if len(players[0].deck) == 0 {
			return players[1].getScore(), 1
		} else if len(players[1].deck) == 0 {
			return players[0].getScore(), 0
		}

		if debug {
			fmt.Printf("-- Round %d --\n", round)
		}

		if isPartTwo {
			state := players[0].getState() + players[1].getState()

			if _, ok := gameState[state]; ok {
				// player 1 wins when the state has been seen before
				return -1, 0
			}

			gameState[state] = true
		}

		p1 := players[0].playTopCard(debug)
		p2 := players[1].playTopCard(debug)

		if debug {
			fmt.Printf("%s plays: %d\n", players[0].id, p1)
			fmt.Printf("%s plays: %d\n", players[1].id, p2)
		}

		if isPartTwo && len(players[0].deck) >= p1 && len(players[1].deck) >= p2 {
			// play a recursive game
			newPlayers := []player{
				player{
					id:   players[0].id,
					deck: make([]int, p1),
				},
				player{
					id:   players[1].id,
					deck: make([]int, p2),
				},
			}

			copy(newPlayers[0].deck, players[0].deck)
			copy(newPlayers[1].deck, players[1].deck)

			if debug {
				fmt.Printf("Playing a sub-game to determine the winner...\n\n")
			}

			_, winner := play(newPlayers, game+1, isPartTwo, debug)

			if debug {
				fmt.Printf("The winner of game %d is %s!\n\n", game+1, newPlayers[winner].id)
				fmt.Printf("...anyway, back to game %d\n", game)
			}

			if winner == 0 {
				players[winner].deck = append(players[winner].deck, p1, p2)

				if debug {
					fmt.Printf("%s wins the round!\n\n", players[winner].id)
				}
			} else {
				players[winner].deck = append(players[winner].deck, p2, p1)

				if debug {
					fmt.Printf("%s wins the round!\n\n", players[winner].id)
				}
			}

			round++
			continue
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
