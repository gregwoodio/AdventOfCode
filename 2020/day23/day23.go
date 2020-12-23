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
	curr, low, high := makeCircle(input)

	for i := 0; i < 100; i++ {
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

		destStart := curr.clockwise
		for {
			if destStart.id == target {
				break
			}

			destStart = destStart.clockwise
		}

		destEnd := destStart.clockwise
		destStart.clockwise = sublistStart
		sublistStart.counterClockwise = destStart

		sublistEnd.clockwise = destEnd
		destEnd.counterClockwise = sublistEnd

		curr = curr.clockwise
	}

	result := ""

	for curr.id != 1 {
		curr = curr.clockwise
	}

	curr = curr.clockwise
	for curr.id != 1 {
		result += strconv.Itoa(curr.id)
		curr = curr.clockwise
	}

	return result
}

func solvePartTwo(input string) string {
	return ""
}

// makes the linked list of cups, returns current cup, low and high values
func makeCircle(input string) (*cup, int, int) {
	var curr, c, last *cup
	low, high := math.MaxInt32, 0

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

		if curr == nil {
			curr = c
		}

		if last != nil {
			last.clockwise = c
			c.counterClockwise = last
		}

		last = c
	}

	curr.counterClockwise = last
	last.clockwise = curr

	return curr, low, high
}
