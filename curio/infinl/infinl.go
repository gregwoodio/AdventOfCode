package infinl

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Navigation woes

This year, just like previous years, Santa Claus wants to visit all the houses again to bring the presents to all the lovely children. Unfortunately, during the preparations, it turns out that his Santa Positioning System (KPS) is no longer working and with the worldwide chip shortage, a new order is not possible.

One of the elves remembers the old-fashioned method and asks Santa to look in the glove compartment of the sleigh. Here he finds a large magic scroll with navigation instructions. This list is full of rules for walking, turning and jumping.

There is a small card on the roll with a separate instruction to always look towards his house when he starts the navigation instructions. Of course, as everyone knows, Santa lives in the North Pole because he is extremely allergic to penguins.

Suppose the following navigation instructions are on the list:

spin 90 (draai)
loop 6 (loop)
jump 2 (spring)
spin -45 (draai)
loop 2 (loop)

There are five instructions in this example .

    According to the small map, Santa should start facing north.
    With the first instruction, Santa turns 90 degrees (to the right) and faces east.
    If he follows the second instruction he takes 6 steps (towards the east).
    Instruction number three makes Santa take a leap forward. The distance he jumps is equal to that of two steps.
    The fourth instruction makes it turn -45 degrees. This instruction causes Santa to face north-east.
    With the fifth and final instruction, Santa takes two steps to the northeast. If Santa takes one diagonal step, this corresponds in terms of distance to one step horizontally and one step vertically.

Santa breaks out in sweat. One of the most important functions of his KPS is missing for his time schedule and that is the Manhattan distance between the starting point and the ending point. In the example above, that's a Manhattan distance of 12.

Download the navigation instructions ➜

Given the navigation instructions, find the Manhattan distance between the start and end points. This distance is then the answer to part 1 .

Part 2

While Santa was busy with all the navigation instructions, it just started snowing. Because of this, Santa has left a trail of steps in the snow. As Santa leaves on his horse-drawn sleigh, he looks back again and sees that his steps have made a pattern of letters.

Look closely at the tracks in the snow and look for the word that Santa left with his steps in the snow. This word is the answer to part 2.
*/

type direction int

const (
	north     direction = 0
	northeast direction = 45
	east      direction = 90
	southeast direction = 135
	south     direction = 180
	southwest direction = 225
	west      direction = 270
	northwest direction = 315
)

func solve(input []string) int {
	transforms := map[direction][]int{
		north:     []int{1, 0},
		northeast: []int{1, 1},
		east:      []int{0, 1},
		southeast: []int{-1, 1},
		south:     []int{-1, 0},
		southwest: []int{-1, -1},
		west:      []int{0, -1},
		northwest: []int{1, -1},
	}

	// steps north, east
	steps := []int{0, 0}
	visited := map[int]map[int]bool{}
	facing := north

	for _, line := range input {
		inst := strings.Split(line, " ")
		val, err := strconv.Atoi(inst[1])
		if err != nil {
			fmt.Printf("Could not parse '%s'\n", inst[1])
		}

		switch inst[0] {
		case "draai": // spin
			facing = direction(int(facing)+val) % 360
			if facing < 0 {
				facing = direction(facing + 360)
			}
			break
		case "loop": // step forward one
			transform := transforms[facing]
			for i := 0; i < val; i++ {
				steps[0] += transform[0]
				steps[1] += transform[1]
				_, ok := visited[steps[1]]
				if !ok {
					visited[steps[1]] = map[int]bool{}
				}
				visited[steps[1]][steps[0]] = true
			}
			break
		case "spring": // jump forward
			transform := transforms[facing]
			steps[0] += transform[0] * val
			steps[1] += transform[1] * val

			_, ok := visited[steps[1]]
			if !ok {
				visited[steps[1]] = map[int]bool{}
			}
			visited[steps[1]][steps[0]] = true
			break
		}
	}

	// print out Santa's steps
	maxX := 0
	maxY := 0
	for key, _ := range visited {
		if key > maxX {
			maxX = key
		}
		for key2, _ := range visited[key] {
			if key2 > maxY {
				maxY = key2
			}
		}
	}
	for i := maxY; i >= 0; i-- {
		for j := 0; j <= maxX; j++ {
			if _, ok := visited[j][i]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	// calculate manhattan distance
	return int(math.Abs(float64(steps[0])) + math.Abs(float64(steps[1])))
}
