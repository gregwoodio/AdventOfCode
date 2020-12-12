package day12

import (
	"fmt"
	"strconv"
)

type direction int

const (
	east  direction = 0
	north direction = 90
	west  direction = 180
	south direction = 270
)

type ship struct {
	facing                direction
	northUnits, eastUnits int
}

func solvePartOne(input []string) int {
	s := ship{
		facing: east,
	}

	for _, line := range input {
		inst := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Errorf("could not parse '%s' to int", line[1:])
		}

		switch inst {
		case 'F':
			if s.facing == east {
				s.eastUnits += val
			} else if s.facing == north {
				s.northUnits += val
			} else if s.facing == west {
				s.eastUnits -= val
			} else if s.facing == south {
				s.northUnits -= val
			}
			break
		case 'E':
			s.eastUnits += val
			break
		case 'N':
			s.northUnits += val
			break
		case 'W':
			s.eastUnits -= val
			break
		case 'S':
			s.northUnits -= val
			break
		case 'R':
			s.facing = direction(((int(s.facing) - val) + 360) % 360)
			break
		case 'L':
			s.facing = direction((int(s.facing) + val) % 360)
			break
		default:
			fmt.Printf("unknown instruction '%s'\n", line)
			break
		}
	}

	return abs(s.northUnits) + abs(s.eastUnits)
}

func solvePartTwo(input []string) int {
	s := ship{}
	wp := ship{
		facing:     east,
		eastUnits:  10,
		northUnits: 1,
	}

	for _, line := range input {
		inst := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Errorf("could not parse '%s' to int", line[1:])
		}

		switch inst {
		case 'F':
			s.eastUnits += wp.eastUnits * val
			s.northUnits += wp.northUnits * val
			break
		case 'E':
			wp.eastUnits += val
			break
		case 'N':
			wp.northUnits += val
			break
		case 'W':
			wp.eastUnits -= val
			break
		case 'S':
			wp.northUnits -= val
			break
		case 'R':
			dir := direction(((int(wp.facing) - val) + 360) % 360)
			for wp.facing != dir {
				wp.facing = (wp.facing + 270) % 360
				if wp.facing == south || wp.facing == north {
					wp.eastUnits *= -1
					wp.eastUnits, wp.northUnits = wp.northUnits, wp.eastUnits
				} else if wp.facing == east || wp.facing == west {
					wp.eastUnits, wp.northUnits = wp.northUnits, wp.eastUnits
					wp.northUnits *= -1
				}
			}
			break
		case 'L':
			dir := direction((int(wp.facing) + val) % 360)
			for wp.facing != dir {
				if wp.facing == south || wp.facing == north {
					wp.eastUnits, wp.northUnits = wp.northUnits, wp.eastUnits
					wp.eastUnits *= -1
				} else if wp.facing == east || wp.facing == west {
					wp.northUnits *= -1
					wp.eastUnits, wp.northUnits = wp.northUnits, wp.eastUnits
				}
				wp.facing = (wp.facing + 90) % 360
			}
			break
		default:
			fmt.Printf("unknown instruction '%s'\n", line)
			break
		}
	}

	return abs(s.northUnits) + abs(s.eastUnits)
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}
