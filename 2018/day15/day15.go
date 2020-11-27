package day15

import (
	"fmt"
	"sort"
)

type team int

const (
	elfTeam team = iota
	goblinTeam
)

type unit struct {
	team        team
	hp          int
	attackPower int
	location    *square
	id          int
}

func (u *unit) isAlive() bool {
	return u.hp > 0
}

func (u *unit) move(b *board) {
	reachableSquares := make(map[square]bool)

	for k, v := range b.allUnits {
		if v.team == u.team {
			continue
		}

		for _, v := range k.passableNeighbours(b) {
			reachableSquares[v] = true
		}
	}

	// bfs until we reach a destination
	visited := make(map[square]bool)
	destinations := candidateSquares{}

	workList := []candidateSquare{
		{
			destination: *u.location,
			firstStep:   nil,
			steps:       0,
		},
	}

	for len(workList) != 0 {
		w := workList[0]
		workList = workList[1:]

		if len(destinations) > 0 && w.steps > destinations[0].steps {
			break
		}

		if reachableSquares[w.destination] {
			destinations = append(destinations, w)
			continue
		}

		for _, v := range w.destination.passableNeighbours(b) {
			if visited[v] {
				continue
			}

			var firstStep square

			if w.firstStep == nil {
				firstStep = v
			} else {
				firstStep = *w.firstStep
			}

			entry := candidateSquare{
				destination: v,
				firstStep:   &firstStep,
				steps:       w.steps + 1,
			}

			visited[v] = true
			workList = append(workList, entry)
		}
	}

	if len(destinations) != 0 {
		sort.Sort(destinations)
		dest := destinations[0]

		delete(b.allUnits, *u.location)
		u.location = dest.firstStep
		b.allUnits[*dest.firstStep] = u
	}
}

// find a target adjacent to the unit's current location
func (u *unit) findTarget(b *board) *unit {
	var targets unitsByAttackOrder

	up := square{x: u.location.x, y: u.location.y - 1}
	opp := b.allUnits[up]
	if b.hasSquare(u.location.x, u.location.y-1) && opp != nil && opp.team != u.team {
		targets = append(targets, opp)
	}

	left := square{x: u.location.x - 1, y: u.location.y}
	opp = b.allUnits[left]
	if b.hasSquare(u.location.x-1, u.location.y) && opp != nil && opp.team != u.team {
		targets = append(targets, opp)
	}

	down := square{x: u.location.x, y: u.location.y + 1}
	opp = b.allUnits[down]
	if b.hasSquare(u.location.x, u.location.y+1) && opp != nil && opp.team != u.team {
		targets = append(targets, opp)
	}

	right := square{x: u.location.x + 1, y: u.location.y}
	opp = b.allUnits[right]
	if b.hasSquare(u.location.x+1, u.location.y) && opp != nil && opp.team != u.team {
		targets = append(targets, opp)
	}

	if len(targets) == 0 {
		return nil
	}

	sort.Sort(targets)
	return targets[0]
}

type unitsByReadingOrder []*unit

func (units unitsByReadingOrder) Len() int {
	return len(units)
}

func (units unitsByReadingOrder) Less(i, j int) bool {
	return hasLowerReadingOrder(units[i].location, units[j].location)
}

func (units unitsByReadingOrder) Swap(i, j int) {
	units[i], units[j] = units[j], units[i]
}

type unitsByAttackOrder []*unit

func (units unitsByAttackOrder) Len() int {
	return len(units)
}

func (units unitsByAttackOrder) Less(i, j int) bool {
	// pull nils to the right
	if units[i] == nil {
		return false
	}

	if units[j] == nil {
		return true
	}

	if units[i].hp != units[j].hp {
		return units[i].hp < units[j].hp
	}

	return hasLowerReadingOrder(units[i].location, units[j].location)
}

func (units unitsByAttackOrder) Swap(i, j int) {
	units[i], units[j] = units[j], units[i]
}

type square struct {
	x, y int
}

func (s square) passableNeighbours(b *board) []square {
	up := s
	up.y--

	right := s
	right.x++

	down := s
	down.y++

	left := s
	left.x--

	neighbours := []square{}

	for _, sq := range []square{up, left, right, down} {
		if !b.hasSquare(sq.x, sq.y) {
			// wall
			continue
		}

		if b.allUnits[sq] != nil {
			// already a unit in that square
			continue
		}

		neighbours = append(neighbours, sq)
	}

	return neighbours
}

type candidateSquare struct {
	destination square
	firstStep   *square
	steps       int
}

type candidateSquares []candidateSquare

func (cs candidateSquares) Len() int {
	return len(cs)
}

func (cs candidateSquares) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs candidateSquares) Less(i, j int) bool {
	return hasLowerReadingOrder(&cs[i].destination, &cs[j].destination)
}

func hasLowerReadingOrder(a, b *square) bool {
	if a.y != b.y {
		return a.y < b.y
	}

	return a.x < b.x
}

type board struct {
	turn          int
	height, width int
	battlemap     map[square]bool
	allUnits      map[square]*unit
	units         unitsByReadingOrder
}

func (b *board) hasSquare(x, y int) bool {
	square := square{x: x, y: y}
	_, ok := b.battlemap[square]

	return ok
}

func (b *board) print() {
	if b.turn == 0 {
		fmt.Print("Initially:\n")
	} else if b.turn == 1 {
		fmt.Printf("After %d round:\n", b.turn)
	} else {
		fmt.Printf("After %d rounds:\n", b.turn)
	}

	for y := 0; y < b.height; y++ {
		units := []*unit{}

		for x := 0; x < b.width; x++ {
			if !b.hasSquare(x, y) {
				fmt.Print("#")
				continue
			}

			sq := square{x: x, y: y}
			occupant := b.allUnits[sq]
			if occupant == nil {
				fmt.Print(".")
			} else if occupant.team == elfTeam {
				units = append(units, occupant)
				fmt.Print("E")
			} else {
				units = append(units, occupant)
				fmt.Print("G")
			}
		}

		if len(units) > 0 {
			unitStr := "\t"

			for _, u := range units {
				if u.team == elfTeam {
					unitStr += fmt.Sprintf("E(%d) ", u.hp)
				} else {
					unitStr += fmt.Sprintf("G(%d) ", u.hp)
				}
			}

			fmt.Print(unitStr)
		}

		fmt.Print("\n")
	}

	fmt.Println()
}

func (b *board) checkWin() (bool, int) {
	elves, goblins := 0, 0

	for _, u := range b.units {
		if u.team == elfTeam {
			elves += u.hp
		}

		if u.team == goblinTeam {
			goblins += u.hp
		}

		if elves > 0 && goblins > 0 {
			return false, -1
		}
	}

	if elves > 0 {
		return true, elves * b.turn
	}

	return true, goblins * b.turn
}

func solvePartOne(input []string, debug bool) int {
	b := makeBoard(input, 3)

	for {
		result := simulate(b, false, debug)

		if result != -1 {
			return result
		}
	}
}

func solvePartTwo(input []string, debug bool) int {
	elfPower := 4
	b := makeBoard(input, elfPower)

	for {
		// for part two, result is either -1 for loss or the turn number of win
		result := simulate(b, true, debug)

		if result == -2 {
			// elf died in this simulation. Increase power!
			elfPower++
			b = makeBoard(input, elfPower)
		} else if result != -1 {
			return result
		}
	}
}

func simulate(b *board, isPartTwo, debug bool) int {

	if debug {
		b.print()
	}

	sort.Sort(b.units)

	for _, u := range b.units {
		if !u.isAlive() {
			continue
		}

		win, result := b.checkWin()
		if win {
			b.print()
			return result
		}

		target := u.findTarget(b)

		if target == nil {
			// not next to an enemy, need to move
			u.move(b)

			target = u.findTarget(b)
		}

		if target != nil {
			target.hp -= u.attackPower

			if target.hp <= 0 {
				delete(b.allUnits, *target.location)
				if isPartTwo && target.team == elfTeam {
					return -2
				}
			}
		}

	}

	alive := unitsByReadingOrder{}
	for _, u := range b.units {
		if u.isAlive() {
			alive = append(alive, u)
		}
	}
	b.units = alive

	b.turn++

	return -1
}

func makeBoard(input []string, elfAttackPower int) *board {
	b := &board{
		battlemap: make(map[square]bool),
		units:     []*unit{},
		allUnits:  make(map[square]*unit),
		height:    len(input),
		width:     len(input[0]),
	}

	id := 1

	for y, line := range input {
		for x, ch := range line {

			if ch == '#' {
				// wall
				continue
			}

			// add square
			s := square{
				x: x,
				y: y,
			}

			b.battlemap[s] = true

			// add unit
			if ch == 'E' {

				u := &unit{
					location:    &s,
					hp:          200,
					attackPower: elfAttackPower,
					id:          id,
					team:        elfTeam,
				}
				id++

				b.units = append(b.units, u)
				b.allUnits[s] = u
			} else if ch == 'G' {
				u := &unit{
					location:    &s,
					hp:          200,
					attackPower: 3,
					id:          id,
					team:        goblinTeam,
				}
				id++

				b.units = append(b.units, u)
				b.allUnits[s] = u
			}
		}
	}

	return b
}
