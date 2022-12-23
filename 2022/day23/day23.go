package day23

type dir int

const (
	north dir = iota
	northeast
	east
	southeast
	south
	southwest
	west
	northwest
)

type coord struct {
	x int
	y int
}

func (c coord) move(other coord) coord {
	return coord{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

// Return true if the elf has another elf in any of the 8 surrounding directions
func (c coord) hasNeighbour(elves map[coord]bool) bool {
	for _, tran := range transforms {
		check := c.move(tran)
		if _, ok := elves[check]; ok {
			return true
		}
	}
	return false
}

// Returns true if there are elves to the north
func (c coord) checkNorth(elves map[coord]bool) bool {
	_, ne := elves[c.move(transforms[northeast])]
	_, n := elves[c.move(transforms[north])]
	_, nw := elves[c.move(transforms[northwest])]

	return ne || n || nw
}

func (c coord) checkEast(elves map[coord]bool) bool {
	_, ne := elves[c.move(transforms[northeast])]
	_, e := elves[c.move(transforms[east])]
	_, se := elves[c.move(transforms[southeast])]

	return ne || e || se
}

func (c coord) checkSouth(elves map[coord]bool) bool {
	_, se := elves[c.move(transforms[southeast])]
	_, s := elves[c.move(transforms[south])]
	_, sw := elves[c.move(transforms[southwest])]

	return se || s || sw
}

func (c coord) checkWest(elves map[coord]bool) bool {
	_, sw := elves[c.move(transforms[southwest])]
	_, w := elves[c.move(transforms[west])]
	_, nw := elves[c.move(transforms[northwest])]

	return sw || w || nw
}

var transforms map[dir]coord = map[dir]coord{
	north:     {x: 0, y: -1},
	northeast: {x: 1, y: -1},
	east:      {x: 1, y: 0},
	southeast: {x: 1, y: 1},
	south:     {x: 0, y: 1},
	southwest: {x: -1, y: 1},
	west:      {x: -1, y: 0},
	northwest: {x: -1, y: -1},
}

func solvePartOne(input []string) int {
	elves := parse(input)
	directions := []dir{north, south, west, east}
	canMove := true

	for round := 0; canMove && round < 10; round++ {
		canMove = false

		// map proposed coord to list of elves who want to move there
		proposed := map[coord][]coord{}
		next := map[coord]bool{}

		for elf := range elves {
			if !elf.hasNeighbour(elves) {
				next[elf] = true
			} else {
				canMove = true
				hasProposed := false
			outer:
				for _, dir := range directions {
					switch dir {
					case north:
						if !elf.checkNorth(elves) {
							p := elf.move(transforms[north])

							if _, ok := proposed[p]; !ok {
								proposed[p] = []coord{}
							}
							proposed[p] = append(proposed[p], elf)
							hasProposed = true
							break outer
						}
					case south:
						if !elf.checkSouth(elves) {
							p := elf.move(transforms[south])

							if _, ok := proposed[p]; !ok {
								proposed[p] = []coord{}
							}
							proposed[p] = append(proposed[p], elf)
							hasProposed = true
							break outer
						}
					case west:
						if !elf.checkWest(elves) {
							p := elf.move(transforms[west])

							if _, ok := proposed[p]; !ok {
								proposed[p] = []coord{}
							}
							proposed[p] = append(proposed[p], elf)
							hasProposed = true
							break outer
						}
					case east:
						if !elf.checkEast(elves) {
							p := elf.move(transforms[east])

							if _, ok := proposed[p]; !ok {
								proposed[p] = []coord{}
							}
							proposed[p] = append(proposed[p], elf)
							hasProposed = true
							break outer
						}
					}
				}
				if !hasProposed {
					next[elf] = true
				}
			}
		}

		// checksum
		count := len(next)
		for _, movers := range proposed {
			count += len(movers)
		}
		if count != len(elves) {
			panic("Lost some elves!")
		}

		// each elf who can move has proposed a new destination. If any
		// list of proposed coords is longer than 1 add the elf to next
		// instead.
		for p, movers := range proposed {
			if len(movers) > 1 {
				for _, e := range movers {
					next[e] = true
				}
			} else {
				next[p] = true
			}
		}

		// rotate list
		directions = append(directions[1:], directions[0])

		elves = next
	}

	minX, maxX, minY, maxY := 9999, 0, 9999, 0

	for elf := range elves {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}

	// count number of empty spaces in the rectangle of elves
	count := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			e := coord{x: x, y: y}
			if _, ok := elves[e]; !ok {
				count++
			}
		}
	}

	return count
}

func solvePartTwo(input []string) int {
	return -1
}

// parse input and return a slice of elf coordinates
func parse(input []string) map[coord]bool {
	elves := map[coord]bool{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				elves[coord{x: x, y: y}] = true
			}
		}
	}

	return elves
}
