package day02

func solvePartOne(input []string) int {
	var (
		letterCounts map[rune]int
		twos         int
		threes       int
	)

	for _, letters := range input {
		letterCounts = make(map[rune]int)

		for _, letter := range letters {
			letterCounts[letter]++
		}

		var foundTwo, foundThree bool
		for _, val := range letterCounts {
			if val == 3 && !foundThree {
				threes++
				foundThree = true
			} else if val == 2 && !foundTwo {
				twos++
				foundTwo = true
			}
		}
	}

	return twos * threes
}

func solvePartTwo(input []string) string {
	for _, first := range input {

		for _, second := range input {

			if first == second {
				continue
			}

			diff := 0
			diffIndex := -1
			for i, ch := range first {
				if rune(second[i]) != ch {
					diff++
					diffIndex = i
					if diff > 1 {
						break
					}
				}
			}

			if diff > 1 {
				continue
			}

			return string(first[:diffIndex] + first[diffIndex+1:])
		}
	}

	return ""
}
