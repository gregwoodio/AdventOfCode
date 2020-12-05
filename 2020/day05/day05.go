package day05

func solvePartOne(input []string) int {
	max := 0

	for _, boardingPass := range input {
		val := binaryValue(boardingPass)
		if val > max {
			max = val
		}
	}

	return max
}

func solvePartTwo(input []string) int {
	seats := make(map[int]bool)

	for _, boardingPass := range input {
		seats[binaryValue(boardingPass)] = true
	}

	for seat := 1; seat < 1023; seat++ {
		if _, ok := seats[seat]; !ok {
			_, prev := seats[seat-1]
			_, next := seats[seat+1]

			if next && prev {
				return seat
			}
		}
	}

	return -1
}

func binaryValue(str string) int {
	val := 0

	for _, ch := range str[:7] {
		val = val << 1

		if ch == 'B' {
			val |= 1
		}
	}

	for _, ch := range str[7:] {
		val = val << 1

		if ch == 'R' {
			val |= 1
		}
	}

	return val
}
