package day12

import (
	"regexp"
	"strconv"
)

type point struct {
	x, y, z int
}

type moon struct {
	pos      point
	velocity point
}

func solvePartOne(input []string, timeLimit int) int {
	moons := parseMoons(input)
	simulate(&moons, timeLimit)
	return calcPotentialEnergy(&moons)
}

func simulate(moons *[]moon, timeLimit int) {
	for tick := 0; tick < timeLimit; tick++ {
		gravities := make([]point, len(*moons))

		for i, moon := range *moons {
			gravities[i] = point{
				x: 0,
				y: 0,
				z: 0,
			}

			for j, compMoon := range *moons {
				if i == j {
					continue
				}

				if moon.pos.x > compMoon.pos.x {
					gravities[i].x--
				} else if moon.pos.x < compMoon.pos.x {
					gravities[i].x++
				}

				if moon.pos.y > compMoon.pos.y {
					gravities[i].y--
				} else if moon.pos.y < compMoon.pos.y {
					gravities[i].y++
				}

				if moon.pos.z > compMoon.pos.z {
					gravities[i].z--
				} else if moon.pos.z < compMoon.pos.z {
					gravities[i].z++
				}
			}
		}

		for i := range *moons {
			(*moons)[i].velocity.x += gravities[i].x
			(*moons)[i].velocity.y += gravities[i].y
			(*moons)[i].velocity.z += gravities[i].z

			(*moons)[i].pos.x += (*moons)[i].velocity.x
			(*moons)[i].pos.y += (*moons)[i].velocity.y
			(*moons)[i].pos.z += (*moons)[i].velocity.z
		}
	}
}

func calcPotentialEnergy(moons *[]moon) int {
	sum := 0
	for i := range *moons {
		pot := abs((*moons)[i].pos.x) + abs((*moons)[i].pos.y) + abs((*moons)[i].pos.z)
		kin := abs((*moons)[i].velocity.x) + abs((*moons)[i].velocity.y) + abs((*moons)[i].velocity.z)

		sum += pot * kin
	}

	return sum
}

func parseMoons(input []string) []moon {
	regex := `<x=(?P<xValue>-?[0-9]+), y=(?P<yValue>-?[0-9]+), z=(?P<zValue>-?[0-9]+)>`
	moons := []moon{}

	for _, line := range input {
		var x, y, z int

		compiled := regexp.MustCompile(regex)
		match := compiled.FindStringSubmatch(line)

		for i, subexp := range compiled.SubexpNames() {
			if i == 0 || subexp == "" {
				continue
			}

			if i > len(match) {
				break
			}

			switch subexp {
			case "xValue":
				x, _ = strconv.Atoi(match[i])
				break
			case "yValue":
				y, _ = strconv.Atoi(match[i])
				break
			case "zValue":
				z, _ = strconv.Atoi(match[i])
				break
			}
		}

		moons = append(moons, moon{
			pos: point{
				x: x,
				y: y,
				z: z,
			},
		})
	}

	return moons
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
