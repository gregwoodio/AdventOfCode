package day03

import "strconv"

type point struct {
	x, y int
}

func solvePartOne(input string) int {
	return visitHouses(input, 1)
}

func solvePartTwo(input string) int {
	return visitHouses(input, 2)
}

func visitHouses(input string, santas int) int {
	houses := make(map[string]int)
	santaLocations := []*point{}
	santaIndex := 0

	for i := 0; i < santas; i++ {
		santaLocations = append(santaLocations, &point{
			x: 0,
			y: 0,
		})
	}

	// visit starting location
	houses[key(santaLocations[santaIndex])]++

	for _, inst := range input {
		santa := santaLocations[santaIndex]
		switch inst {
		case '^':
			santa.y--
			break
		case '>':
			santa.x++
			break
		case 'v':
			santa.y++
			break
		case '<':
			santa.x--
			break
		}

		houses[key(santa)]++

		santaIndex = (santaIndex + 1) % santas
	}

	return len(houses)
}

func key(p *point) string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}
