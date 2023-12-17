package day07

import (
	// "fmt"
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

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards    []card
	handType handType
	bid      int
	value    string
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

	var _handType handType
	highest := countValues[len(countValues)-1]
	if highest == 5 {
		_handType = fiveOfAKind
	} else if highest == 4 {
		_handType = fourOfAKind
	} else if highest == 3 {
		if len(countValues)-2 >= 0 && countValues[len(countValues)-2] == 2 {
			_handType = fullHouse
		} else {
			_handType = threeOfAKind
		}
	} else if highest == 2 {
		if len(countValues)-2 >= 0 && countValues[len(countValues)-2] == 2 {
			_handType = twoPair
		} else {
			_handType = onePair
		}
	} else {
		_handType = highCard
	}

	// fmt.Printf("Hand: %s \tCount values: %v \t handType: %d\n", parts[0], countValues, _handType)

	bid, _ := strconv.Atoi(parts[1])

	return hand{
		bid:      bid,
		cards:    cards,
		handType: _handType,
		value:    line,
	}
}

func solvePartOne(input []string) int {
	hands := []hand{}
	for _, line := range input {
		hands = append(hands, parseHand(line))
	}
	// fmt.Println(len(hands))

	sort.Slice(hands, func(i, j int) bool {

		if hands[i].handType == hands[j].handType {
			for k := range hands[i].cards {
				if hands[i].cards[k] != hands[j].cards[k] {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
		}

		return hands[i].handType < hands[j].handType
	})

	sum := 0

	for i, hand := range hands {
		// fmt.Printf("%s \t%d\t%d\t%d\t%d\tcurrentSum: %d\n", hand.value, hand.handType, i+1, hand.bid, (i+1)*hand.bid, sum)
		sum += hand.bid * (i + 1)
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}
