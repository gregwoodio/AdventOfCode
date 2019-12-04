package day04

func solvePartOne(from, to int) int {
	count := 0

	for i := from; i <= to; i++ {
		if isValid(i) {
			count++
		}
	}

	return count
}

func isValid(num int) bool {
	digits := getDigits(num)

	if len(digits) != 6 {
		return false
	}

	hasDouble := false
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}

		if digits[i] == digits[i+1] {
			hasDouble = true
		}
	}

	if !hasDouble {
		return false
	}

	return true
}

func getDigits(num int) []int {
	nums := []int{}
	for num > 0 {
		nums = append([]int{num % 10}, nums...)
		num /= 10
	}
	return nums
}
