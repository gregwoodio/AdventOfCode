package day14

import (
	"strconv"
	"strings"
)

// Term is the count of an element
type term struct {
	element string
	num     int
}

// reaction takes some inputs to produce an output
type reaction struct {
	inputs []term
	output term
}

// state tracks the reaction's leftover values and the
// total amount of ore used
type state struct {
	inventory map[string]int
	oreUsed   int
}

func newState() *state {
	return &state{
		inventory: make(map[string]int),
	}
}

func solvePartOne(input []string) int {
	reactions := parseInputs(input)
	state := newState()
	// fmt.Printf("%s\n", (*reactions)["FUEL"][0].output.element)
	react("FUEL", state, reactions)
	return state.oreUsed
}

func react(element string, state *state, reactions *map[string]reaction) {
	reaction := (*reactions)[element]

	// make sure we have required elements in our inventory
	for _, term := range reaction.inputs {
		if (*state).inventory[term.element] < term.num {
			// create the reaction or use ore
			if term.element == "ORE" {
				(*state).oreUsed += term.num
				(*state).inventory["ORE"] += term.num
			} else {
				for (*state).inventory[term.element] < term.num {
					react(term.element, state, reactions)
				}
			}
		}

		// spend
		(*state).inventory[term.element] -= term.num
	}

	// increase output
	(*state).inventory[element] += reaction.output.num
}

// parseInputs reads the input strings and produces a map
// where the key is the element and the value is a slice of
// reactions that produce that element.
func parseInputs(input []string) *map[string]reaction {
	reactions := make(map[string]reaction)

	for _, line := range input {
		inOut := strings.Split(line, " => ")
		r := reaction{
			inputs: []term{},
		}

		inputTerms := strings.Split(inOut[0], ", ")
		for _, t := range inputTerms {
			parts := strings.Split(t, " ")
			value, _ := strconv.Atoi(parts[0])

			r.inputs = append(r.inputs, term{
				num:     value,
				element: parts[1],
			})
		}

		parts := strings.Split(inOut[1], " ")
		value, _ := strconv.Atoi(parts[0])

		r.output = term{
			num:     value,
			element: parts[1],
		}

		reactions[r.output.element] = r
	}

	return &reactions
}
