package day23

import (
	"math"
	"strconv"
)

type cup struct {
	id               int
	clockwise        *cup
	counterClockwise *cup
}

func solvePartOne(input string) string {
	curr, low, high, dict := makeCircle(input, false)

	play(curr, 100, low, high, dict)

	result := ""

	curr = dict[1].clockwise
	for curr.id != 1 {
		result += strconv.Itoa(curr.id)
		curr = curr.clockwise
	}

	return result
}

func solvePartTwo(input string) int64 {
	curr, low, high, dict := makeCircle(input, true)

	play(curr, 10000000, low, high, dict)

	return int64(dict[1].clockwise.id) * int64(dict[1].clockwise.clockwise.id)
}

func play(curr *cup, rounds, low, high int, dict map[int]*cup) *cup {

	for i := 0; i < rounds; i++ {
		sublistStart := curr.clockwise
		sublistEnd := sublistStart.clockwise.clockwise
		connect := sublistEnd.clockwise

		// close circle
		curr.clockwise = connect
		connect.counterClockwise = curr

		// disconnect sublist
		sublistStart.counterClockwise = nil
		sublistEnd.clockwise = nil

		target := curr.id - 1
		if target < low {
			target = high
		}

		for {
			targetInSublist := false
			sl := sublistStart
			for sl != nil {
				if sl.id == target {
					targetInSublist = true
					break
				}
				sl = sl.clockwise
			}

			if targetInSublist {
				target--
				if target < low {
					target = high
				}
			} else {
				break
			}
		}

		destStart := dict[target]

		destEnd := destStart.clockwise
		destStart.clockwise = sublistStart
		sublistStart.counterClockwise = destStart

		sublistEnd.clockwise = destEnd
		destEnd.counterClockwise = sublistEnd

		curr = curr.clockwise
	}

	return curr
}

func makeCircle(input string, isPartTwo bool) (*cup, int, int, map[int]*cup) {
	var curr, c, last *cup
	low, high := math.MaxInt32, 0
	dict := make(map[int]*cup)

	for _, ch := range input {
		id, _ := strconv.Atoi(string(ch))

		if id > high {
			high = id
		}
		if id < low {
			low = id
		}

		c = &cup{
			id: id,
		}
		dict[id] = c

		if curr == nil {
			curr = c
		}

		if last != nil {
			last.clockwise = c
			c.counterClockwise = last
		}

		last = c
	}

	if isPartTwo {
		id := high + 1
		for count := 9; count < 1000000; count++ {
			c = &cup{
				id: id,
			}
			dict[id] = c

			last.clockwise = c
			c.counterClockwise = last
			id++
			high++

			last = c
		}
	}

	curr.counterClockwise = last
	last.clockwise = curr

	return curr, low, high, dict
}
