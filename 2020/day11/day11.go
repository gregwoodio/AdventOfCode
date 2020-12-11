package day11

type coord struct {
	x, y       int
	isChair    bool
	isOccupied bool
}

type waitingRoom struct {
	coords        [][]*coord
	rounds        int
	seatsOccupied int
}

func (wr waitingRoom) checkDirectionForNeighbour(x, y, dx, dy int, isPartTwo bool) bool {
	// are we still in the grid?
	if x+dx < 0 || x+dx >= len(wr.coords[0]) || y+dy < 0 || y+dy >= len(wr.coords) {
		return false
	}

	if !isPartTwo {
		if wr.coords[y+dy][x+dx].isChair && wr.coords[y+dy][x+dx].isOccupied {
			return true
		}
		return false
	}

	// empty floor, keep looking
	if !wr.coords[y+dy][x+dx].isChair {
		return wr.checkDirectionForNeighbour(x+dx, y+dy, dx, dy, isPartTwo)
	}

	return wr.coords[y+dy][x+dx].isOccupied
}

func (wr waitingRoom) simulate(isPartTwo bool) int {
	wr.rounds = 1
	for {
		// fmt.Println()

		new := [][]*coord{}
		wr.seatsOccupied = 0
		for y, row := range wr.coords {
			new = append(new, []*coord{})

			for x, c := range row {
				if !c.isChair {
					new[y] = append(new[y], &coord{
						x:       x,
						y:       y,
						isChair: false,
					})
					// fmt.Print(".")
					continue
				}

				neighbours := 0
				if wr.checkDirectionForNeighbour(x, y, -1, -1, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, -1, 0, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, -1, 1, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, 0, -1, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, 0, 1, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, 1, -1, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, 1, 0, isPartTwo) {
					neighbours++
				}
				if wr.checkDirectionForNeighbour(x, y, 1, 1, isPartTwo) {
					neighbours++
				}

				newChair := &coord{
					x:       x,
					y:       y,
					isChair: true,
				}
				if neighbours == 0 && !c.isOccupied {
					newChair.isOccupied = true
				} else if !isPartTwo && c.isOccupied && neighbours >= 4 || c.isOccupied && neighbours >= 5 {
					newChair.isOccupied = false
				} else {
					newChair.isOccupied = c.isOccupied
				}

				if newChair.isOccupied {
					// fmt.Print("#")
					wr.seatsOccupied++
					// } else {
					// 	fmt.Print("L")
				}

				new[y] = append(new[y], newChair)
			}

			// fmt.Println()
		}

		isSame := true

	check:
		for y, row := range wr.coords {
			for x, c := range row {
				if new[y][x].isOccupied != c.isOccupied {
					isSame = false
					break check
				}
			}
		}

		if isSame {
			return wr.seatsOccupied
		}

		wr.coords = new
		wr.rounds++
	}
}

func solvePartOne(input []string) int {
	wr := makeWaitingRoom(input)
	seatsOccupied := wr.simulate(false)
	return seatsOccupied
}

func solvePartTwo(input []string) int {
	wr := makeWaitingRoom(input)
	seatsOccupied := wr.simulate(true)
	return seatsOccupied
}

func makeWaitingRoom(input []string) *waitingRoom {
	wr := waitingRoom{}
	wr.coords = [][]*coord{}

	//fmt.Println("Initial:")

	for y, line := range input {
		wr.coords = append(wr.coords, []*coord{})
		for x, ch := range line {
			isChair := true
			if ch == '.' {
				isChair = false
			}

			c := &coord{
				x:       x,
				y:       y,
				isChair: isChair,
			}

			wr.coords[y] = append(wr.coords[y], c)

			// if c.isChair {
			// 	fmt.Print("L")
			// } else {
			// 	fmt.Print(".")
			// }
		}
		// fmt.Println()
	}

	return &wr
}
