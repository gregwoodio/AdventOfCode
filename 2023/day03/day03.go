package day03

import (
	"fmt"
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
	sum := 0

	gear := "*"
	gearNums := map[string][]int{}

	for row, line := range input {
		num := 0
		gearLocation := ""

		for col, ch := range line {
			if !unicode.IsDigit(ch) {
				if gearLocation != "" {
					if _, ok := gearNums[gearLocation]; !ok {
						gearNums[gearLocation] = []int{}
					}
					gearNums[gearLocation] = append(gearNums[gearLocation], num)
				}

				num = 0
				gearLocation = ""
			} else if unicode.IsDigit(ch) {
				curr, _ := strconv.Atoi(string(ch))
				num *= 10
				num += curr

				// check for surrounding symbols
				if gearLocation == "" {
					if row-1 >= 0 {
						if col-1 >= 0 && strings.Contains(gear, string(input[row-1][col-1])) {
							gearLocation = xyKey(col-1, row-1)
						}
						if strings.Contains(gear, string(input[row-1][col])) {
							gearLocation = xyKey(col, row-1)
						}
						if col+1 < len(line) && strings.Contains(gear, string(input[row-1][col+1])) {
							gearLocation = xyKey(col+1, row-1)
						}
					}
					if col-1 >= 0 && strings.Contains(gear, string(input[row][col-1])) {
						gearLocation = xyKey(col-1, row)
					}
					if col+1 < len(line) && strings.Contains(gear, string(input[row][col+1])) {
						gearLocation = xyKey(col+1, row)
					}

					if row+1 < len(input) {
						if col-1 >= 0 && strings.Contains(gear, string(input[row+1][col-1])) {
							gearLocation = xyKey(col-1, row+1)
						}
						if strings.Contains(gear, string(input[row+1][col])) {
							gearLocation = xyKey(col, row+1)
						}
						if col+1 < len(line) && strings.Contains(gear, string(input[row+1][col+1])) {
							gearLocation = xyKey(col+1, row+1)
						}
					}
				}
			}
		}

		if gearLocation != "" {
			if _, ok := gearNums[gearLocation]; !ok {
				gearNums[gearLocation] = []int{}
			}
			gearNums[gearLocation] = append(gearNums[gearLocation], num)
		}
		num = 0
		gearLocation = ""
	}

	for _, list := range gearNums {
		if len(list) == 2 {
			sum += list[0] * list[1]
		}
	}

	return sum
}

func xyKey(x, y int) string {
	return fmt.Sprintf("x:%d,y:%d", x, y)
}
