package day04

func solvePartOne(from, to int) int {
	count := 0

	for i := from; i <= to; i++ {
		if isValid(i, false) {
			count++
		}
	}

	return count
}

func solvePartTwo(from, to int) int {
	count := 0

	for i := from; i <= to; i++ {
		if isValid(i, true) {
			count++
		}
	}

	return count
}

func isValid(num int, isPartTwo bool) bool {
	digits := getDigits(num)
	countDigits := [10]int{}

	if len(digits) != 6 {
		return false
	}

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}

		countDigits[digits[i]]++
	}

	countDigits[digits[5]]++

	if !isPartTwo {
		for _, val := range countDigits {
			if val >= 2 {
				return true
			}
		}
	} else {
		for _, val := range countDigits {
			if val == 2 {
				return true
			}
		}
	}

	return false
}

func getDigits(num int) []int {
	nums := []int{}
	for num > 0 {
		nums = append([]int{num % 10}, nums...)
		num /= 10
	}
	return nums
}
