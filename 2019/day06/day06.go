package day06

import "strings"

type planet struct {
	name       string
	orbiting   *planet
	satellites []*planet
}

func (planet planet) countOrbits() int {
	if planet.orbiting == nil {
		return 0
	}
	return planet.orbiting.countOrbits() + 1
}

func (planet planet) hasChildOrbit(key string) bool {
	// recurse through satellites until we find the key or not
	if len(planet.satellites) == 0 {
		return false
	}

	for _, sat := range planet.satellites {
		if sat.name == key {
			return true
		}
	}

	childHasKey := false
	for _, sat := range planet.satellites {
		if childHasKey == true {
			break
		} else {
			childHasKey = sat.hasChildOrbit(key)
		}
	}

	return childHasKey
}

func solvePartOne(input []string) int {
	planets := findOrbits(input)

	var sum int
	for p := range planets {
		sum += planets[p].countOrbits()
	}

	return sum
}

func solvePartTwo(input []string) int {
	planets := findOrbits(input)

	steps := 0
	you := planets["YOU"]
	key := you.orbiting.name

	// step up from your orbit until we find a planet who has Santa in a child orbit
	for ; !planets[key].hasChildOrbit("SAN"); key = planets[key].orbiting.name {
		steps++
	}

	target := key

	// step up from Santa orbit until we find the target planet
	santa := planets["SAN"]
	key = santa.orbiting.name
	for ; key != target; key = planets[key].orbiting.name {
		steps++
	}

	return steps
}

func findOrbits(input []string) map[string]*planet {
	planets := make(map[string]*planet)

	for _, val := range input {
		values := strings.Split(val, ")")

		plan := planets[values[0]]
		if plan == nil {
			plan = &planet{
				name:       values[0],
				orbiting:   nil,
				satellites: []*planet{},
			}

			planets[plan.name] = plan
		}

		sat := planets[values[1]]

		if sat == nil {
			sat = &planet{
				name:       values[1],
				orbiting:   nil,
				satellites: []*planet{},
			}

			planets[sat.name] = sat
		}

		sat.orbiting = plan
		plan.satellites = append(plan.satellites, sat)
	}

	return planets
}
