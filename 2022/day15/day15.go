package day15

import (
	"regexp"
	"strconv"
)

type coord struct {
	x, y int
}

type coverage struct {
	start, end int
}

type beacon struct {
	pos      coord
	distance int
}

var regex = `Sensor at x=(?P<sensorX>-?[0-9]+), y=(?P<sensorY>-?[0-9]+): closest beacon is at x=(?P<beaconX>-?[0-9]+), y=(?P<beaconY>-?[0-9]+)`

func solvePartOne(input []string, row int) int {
	sensors := parse(input)

	board := map[coord]bool{}
	for s, b := range sensors {
		for r := s.y - b.distance; r <= s.y+b.distance; r++ {
			if r != row {
				continue
			}
			for c := s.x - b.distance; c <= s.x+b.distance; c++ {
				p := coord{x: c, y: r}
				if manhattan(p, s) > b.distance {
					continue
				}
				board[p] = true
			}
		}
		board[b.pos] = false
	}

	count := 0
	for c, isEmpty := range board {
		if !isEmpty {
			continue
		}
		if c.y == row {
			count++
		}
	}

	return count
}

func solvePartTwo(input []string, row int) int {
	sensors := parse(input)
	final := coord{}

outer:
	for y := 0; y <= row; y++ {
		coverages := []coverage{}
		for s, b := range sensors {
			offset := abs(s.y - y)
			if offset > b.distance {
				continue
			}
			c := coverage{}
			c.start = s.x - b.distance + offset
			c.end = s.x + b.distance - offset
			coverages = append(coverages, c)
		}

		// more bubble sort 😳
		swapped := true
		for swapped {
			swapped = false
			for i := 0; i < len(coverages)-1; i++ {
				if coverages[i].start > coverages[i+1].start {
					coverages[i], coverages[i+1] = coverages[i+1], coverages[i]
					swapped = true
				}
			}
		}

		x := 0
		for _, c := range coverages {
			if c.start > x {
				final = coord{x: x, y: y}
				break outer
			}
			if c.end+1 > x {
				x = c.end + 1
			}
		}

		if x <= row {
			final = coord{x: x, y: y}
			break outer
		}
	}

	return final.x*4000000 + final.y
}

func parse(input []string) map[coord]beacon {

	sensors := map[coord]beacon{}

	compiled := regexp.MustCompile(regex)
	sx, sy, bx, by := 0, 0, 0, 0
	for _, line := range input {
		match := compiled.FindStringSubmatch(line)

		for i, subexpName := range compiled.SubexpNames() {
			if i == 0 || subexpName == "" {
				continue
			}

			if i > len(match) {
				break
			}

			switch subexpName {
			case "sensorX":
				sx, _ = strconv.Atoi(match[i])
				break
			case "sensorY":
				sy, _ = strconv.Atoi(match[i])
				break
			case "beaconX":
				bx, _ = strconv.Atoi(match[i])
				break
			case "beaconY":
				by, _ = strconv.Atoi(match[i])
				break
			}
		}

		s := coord{x: sx, y: sy}
		b := coord{x: bx, y: by}

		sensors[s] = beacon{
			pos:      b,
			distance: manhattan(s, b),
		}

	}

	return sensors
}

func manhattan(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}
