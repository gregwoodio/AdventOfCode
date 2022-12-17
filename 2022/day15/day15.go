package day15

import (
	"regexp"
	"strconv"
)

type coord struct {
	x, y int
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
	return -1
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
