package day11

import "fmt"

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

func (wr waitingRoom) simulate() int {
	wr.rounds = 1
	for {
		fmt.Println()

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
					fmt.Print(".")
					continue
				}

				neighbours := 0
				if x-1 >= 0 && y-1 >= 0 && wr.coords[y-1][x-1].isChair && wr.coords[y-1][x-1].isOccupied {
					neighbours++
				}
				if x-1 >= 0 && wr.coords[y][x-1].isChair && wr.coords[y][x-1].isOccupied {
					neighbours++
				}
				if x-1 >= 0 && y+1 < len(wr.coords) && wr.coords[y+1][x-1].isChair && wr.coords[y+1][x-1].isOccupied {
					neighbours++
				}
				if y-1 >= 0 && wr.coords[y-1][x].isChair && wr.coords[y-1][x].isOccupied {
					neighbours++
				}
				if y+1 < len(wr.coords) && wr.coords[y+1][x].isChair && wr.coords[y+1][x].isOccupied {
					neighbours++
				}
				if x+1 < len(row) && y-1 >= 0 && wr.coords[y-1][x+1].isChair && wr.coords[y-1][x+1].isOccupied {
					neighbours++
				}
				if x+1 < len(row) && wr.coords[y][x+1].isChair && wr.coords[y][x+1].isOccupied {
					neighbours++
				}
				if x+1 < len(row) && y+1 < len(wr.coords) && wr.coords[y+1][x+1].isChair && wr.coords[y+1][x+1].isOccupied {
					neighbours++
				}

				newChair := &coord{
					x:       x,
					y:       y,
					isChair: true,
				}
				if neighbours == 0 && !c.isOccupied {
					newChair.isOccupied = true
				} else if c.isOccupied && neighbours >= 4 {
					newChair.isOccupied = false
				} else {
					newChair.isOccupied = c.isOccupied
				}

				if newChair.isOccupied {
					fmt.Print("#")
					wr.seatsOccupied++
				} else {
					fmt.Print("L")
				}

				new[y] = append(new[y], newChair)
			}

			fmt.Println()
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
	seatsOccupied := wr.simulate()
	return seatsOccupied
}

func solvePartTwo(input []string) int {
	return -1
}

func makeWaitingRoom(input []string) *waitingRoom {
	wr := waitingRoom{}
	wr.coords = [][]*coord{}

	fmt.Println("Initial:")

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

			if c.isChair {
				fmt.Print("L")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	return &wr
}
