package day01

func solvePartOne(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}

	return -1
}

func solvePartTwo(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}

	return -1
}
