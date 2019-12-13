package day12

import (
	"fmt"
	"math/big"
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

func solvePartTwo(input []string) uint64 {
	moons := parseMoons(input)
	periods := []uint64{}

	for _, val := range []int{1, 2, 3} { // x, y, z
		periods = append(periods, findPeriod(val, &moons))
	}

	fmt.Printf("Periods:\n\tX: %d\n\tY: %d\n\tZ:%d\n", periods[0], periods[1], periods[2])
	var x, y, z, lcm1, lcm2, lcm3 big.Int

	// LCM(a, b) = a * b / GCD(a, b)
	x.SetUint64(periods[0])
	y.SetUint64(periods[1])
	z.SetUint64(periods[2])

	lcm1.Mul(lcm1.Div(&x, lcm1.GCD(nil, nil, &x, &y)), &y)
	lcm2.Mul(lcm2.Div(&y, lcm2.GCD(nil, nil, &y, &z)), &z)

	lcm3.Mul(lcm3.Div(&lcm1, lcm3.GCD(nil, nil, &lcm1, &lcm2)), &lcm2)

	return lcm3.Uint64()
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

// variable is 1=x, 2=y, 3=z
func findPeriod(variable int, moons *[]moon) uint64 {
	positions := make(map[string]bool)
	var tick uint64

	for tick = 0; ; tick++ {
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

				switch variable {
				case 1:
					if moon.pos.x > compMoon.pos.x {
						gravities[i].x--
					} else if moon.pos.x < compMoon.pos.x {
						gravities[i].x++
					}
					break

				case 2:
					if moon.pos.y > compMoon.pos.y {
						gravities[i].y--
					} else if moon.pos.y < compMoon.pos.y {
						gravities[i].y++
					}
					break

				case 3:
					if moon.pos.z > compMoon.pos.z {
						gravities[i].z--
					} else if moon.pos.z < compMoon.pos.z {
						gravities[i].z++
					}
					break
				}
			}
		}

		for i := range *moons {
			switch variable {
			case 1:
				(*moons)[i].velocity.x += gravities[i].x
				(*moons)[i].pos.x += (*moons)[i].velocity.x
				break
			case 2:

				(*moons)[i].velocity.y += gravities[i].y
				(*moons)[i].pos.y += (*moons)[i].velocity.y
				break

			case 3:
				(*moons)[i].velocity.z += gravities[i].z
				(*moons)[i].pos.z += (*moons)[i].velocity.z
				break
			}
		}

		var key string
		for i := range *moons {
			switch variable {
			case 1:
				key += fmt.Sprintf("<%d,%d>", (*moons)[i].pos.x, (*moons)[i].velocity.x)
				break
			case 2:
				key += fmt.Sprintf("<%d,%d>", (*moons)[i].pos.y, (*moons)[i].velocity.y)
				break
			case 3:
				key += fmt.Sprintf("<%d,%d>", (*moons)[i].pos.z, (*moons)[i].velocity.z)
				break
			}
		}

		if positions[key] == true {
			return tick
		}

		positions[key] = true
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
