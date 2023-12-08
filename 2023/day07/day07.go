package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type card int

const (
	two card = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

type strength int

const (
	highCard strength = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards []card
	str   strength
	bid   int
	value string
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")
	counts := map[rune]int{}
	cards := []card{}

	for _, r := range parts[0] {
		if r == '2' {
			cards = append(cards, two)
		} else if r == '3' {
			cards = append(cards, three)
		} else if r == '4' {
			cards = append(cards, four)
		} else if r == '5' {
			cards = append(cards, five)
		} else if r == '6' {
			cards = append(cards, six)
		} else if r == '7' {
			cards = append(cards, seven)
		} else if r == '8' {
			cards = append(cards, eight)
		} else if r == '9' {
			cards = append(cards, nine)
		} else if r == 'T' {
			cards = append(cards, ten)
		} else if r == 'J' {
			cards = append(cards, jack)
		} else if r == 'Q' {
			cards = append(cards, queen)
		} else if r == 'K' {
			cards = append(cards, king)
		} else if r == 'A' {
			cards = append(cards, ace)
		}

		counts[r]++
	}

	countValues := []int{}
	for _, value := range counts {
		countValues = append(countValues, value)
	}
	sort.Ints(countValues)

	var str strength
	highest := countValues[len(countValues)-1]
	if highest == 5 {
		str = fiveOfAKind
	} else if highest == 4 {
		str = fourOfAKind
	} else if highest == 3 {
		if len(countValues)-2 >= 0 && countValues[len(countValues)-2] == 2 {
			str = fullHouse
		} else {
			str = threeOfAKind
		}
	} else if highest == 2 {
		if len(countValues)-2 >= 0 && countValues[len(countValues)-2] == 2 {
			str = twoPair
		} else {
			str = onePair
		}
	} else {
		str = highCard
	}

	bid, _ := strconv.Atoi(parts[1])

	return hand{
		bid:   bid,
		cards: cards,
		str:   str,
		value: line,
	}
}

func solvePartOne(input []string) int {
	hands := []hand{}
	for _, line := range input {
		hands = append(hands, parseHand(line))
	}
	fmt.Println(len(hands))

	sort.SliceStable(hands, func(i, j int) bool {

		if hands[i].str == hands[j].str {
			for k := range hands[i].cards {
				if hands[i].cards[k] != hands[j].cards[k] {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
		}

		return hands[i].str < hands[j].str
	})

	sum := 0

	for i, hand := range hands {
		fmt.Printf("%s \t%d\t%d\t%d\t%d\n", hand.value, hand.str, i+1, hand.bid, (i+1)*hand.bid)
		sum += hand.bid * (i + 1)
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}
