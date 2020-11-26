package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// sample value should be 18
	values := process("input.txt")

	partOne(values)
	partTwo(values)
}

func partOne(numbers [][]float64) {
	var sum float64
	for _, line := range numbers {
		var low float64 = 10000
		var high float64 = -1
		for _, num := range line {
			high = math.Max(high, num)
			low = math.Min(low, num)
		}
		sum += high - low
	}
	fmt.Printf("Part 1: %.0f\n", sum)
}

func partTwo(numbers [][]float64) {
	var sum float64

	for _, line := range numbers {
		for _, num := range line {
			for _, comp := range line {
				if int(num)%int(comp) == 0 && num != comp {
					sum += num / comp
				}
			}
		}
	}

	fmt.Printf("Part 2: %.0f\n", sum)
}

func process(filepath string) [][]float64 {
	file, _ := os.Open(filepath)
	reader := bufio.NewReader(file)

	var inputNumbers [][]float64

	for line, _, _ := reader.ReadLine(); line != nil; {
		var numbers []float64

		split := strings.Split(string(line), "\t")

		for _, num := range split {
			float, _ := strconv.ParseFloat(num, 64)
			numbers = append(numbers, float)
		}
		inputNumbers = append(inputNumbers, numbers)
		line, _, _ = reader.ReadLine()
	}

	return inputNumbers
}
