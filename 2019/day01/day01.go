package day01

func solvePartOne(input []int) int {
	var sum int

	for _, mass := range input {
		sum += calculateFuel(mass)
	}

	return sum
}

func solvePartTwo(input []int) int {
	var sum int

	for _, mass := range input {
		fuel := calculateFuel(mass)

		for fuel > 0 {
			sum += fuel
			fuel = calculateFuel(fuel)
		}
	}

	return sum
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}
