package day22

import (
	"log"
	"strconv"
	"strings"
)

func solvePartOne(size int, input []string) []int {
	deck := createDeck(size)
	for _, val := range input {
		if val == "deal into new stack" {
			deck = *(dealIntoNewStack(&deck))
		} else if strings.HasPrefix(val, "cut ") {
			n, _ := strconv.Atoi(string(val[4:]))
			deck = *(cut(&deck, n))
		} else if strings.HasPrefix(val, "deal with increment ") {
			n, _ := strconv.Atoi(string(val[20:]))
			deck = *(dealWithIncrement(&deck, n))
		} else {
			log.Fatalf("unknown command: %s", val)
		}
	}

	return deck
}

func dealIntoNewStack(deck *[]int) *[]int {
	for i := 0; i < len(*deck)/2; i++ {
		j := len(*deck) - 1 - i
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	}

	return deck
}

func cut(deck *[]int, n int) *[]int {
	if n >= 0 {
		d := append((*deck)[n:], (*deck)[:n]...)
		deck = &d
	} else {
		limit := len(*deck) + n
		d := append((*deck)[limit:], (*deck)[:limit]...)
		deck = &d
	}

	return deck
}

func dealWithIncrement(deck *[]int, n int) *[]int {
	newDeck := make([]int, len(*deck))
	index := 0

	for i := 0; i < len(*deck); i++ {
		newDeck[index] = (*deck)[i]
		index = (index + n) % len(*deck)
	}

	return &newDeck
}

func createDeck(size int) []int {
	deck := make([]int, size)
	for i := range deck {
		deck[i] = i
	}
	return deck
}
