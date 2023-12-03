package day03

import (
	"strconv"
	"strings"
	"unicode"
)

func solvePartOne(input []string) int {
	sum := 0

	symbols := "&*#%-+=/@$"

	for row, line := range input {
		num := 0
		nextToSymbol := false

		for col, ch := range line {
			if !unicode.IsDigit(ch) {
				if nextToSymbol {
					sum += num
				}

				num = 0
				nextToSymbol = false
			} else if unicode.IsDigit(ch) {
				curr, _ := strconv.Atoi(string(ch))
				num *= 10
				num += curr

				// check for surrounding symbols
				if !nextToSymbol {
					if row-1 >= 0 {
						if (col-1 >= 0 && strings.Contains(symbols, string(input[row-1][col-1]))) ||
							strings.Contains(symbols, string(input[row-1][col])) ||
							(col+1 < len(line) && strings.Contains(symbols, string(input[row-1][col+1]))) {
							nextToSymbol = true
						}
					}
					if (col-1 >= 0 && strings.Contains(symbols, string(input[row][col-1]))) ||
						(col+1 < len(line) && strings.Contains(symbols, string(input[row][col+1]))) {
						nextToSymbol = true
					}

					if row+1 < len(input) {
						if (col-1 >= 0 && strings.Contains(symbols, string(input[row+1][col-1]))) ||
							strings.Contains(symbols, string(input[row+1][col])) ||
							(col+1 < len(line) && strings.Contains(symbols, string(input[row+1][col+1]))) {

							nextToSymbol = true
						}
					}
				}
			}
		}

		if num != 0 && nextToSymbol {
			sum += num
		}
	}

	return sum

}

func solvePartTwo(input []string) int {
	return -1
}
