package day08

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type row struct {
	values []int
}

type layer struct {
	rows []row
}

func solvePartOne(input string, width, height int) int {
	pixels := stringToIntSlice(input)
	layers := findLayers(width, height, pixels)

	// find layer with fewest 0s
	fewestZerosIndex := -1
	fewestZeros := math.MaxInt32

	for i, layer := range layers {
		zeros := 0
		for _, r := range layer.rows {
			for _, val := range r.values {
				if val == 0 {
					zeros++
				}
			}
		}

		if zeros < fewestZeros {
			fewestZeros = zeros
			fewestZerosIndex = i
		}
	}

	// count 1s and 2s
	var ones, twos int
	for _, r := range layers[fewestZerosIndex].rows {
		for _, val := range r.values {
			if val == 1 {
				ones++
			}

			if val == 2 {
				twos++
			}
		}
	}

	return ones * twos
}

func solvePartTwo(input string, width, height int) [][]int {
	pixels := stringToIntSlice(input)
	layers := findLayers(width, height, pixels)

	final := make([][]int, height)

	for col := 0; col < height; col++ {
		final[col] = make([]int, width)

		for row := 0; row < width; row++ {

			for l := 0; l < len(layers); l++ {
				if layers[l].rows[col].values[row] == 2 {
					continue
				}
				final[col][row] = layers[l].rows[col].values[row]
				break
			}

			fmt.Printf("%d", final[col][row])
		}
		fmt.Printf("\n")
	}

	return final
}

func findLayers(width, height int, pixels []int) []layer {
	rows := []row{}

	for i := 0; i < len(pixels); i += width {
		rows = append(rows, row{
			values: pixels[i : i+width],
		})
	}

	layers := []layer{}

	for i := 0; i < len(rows); i += height {
		layers = append(layers, layer{
			rows: rows[i : i+height],
		})
	}

	return layers
}

func stringToIntSlice(input string) []int {
	ints := []int{}
	for _, val := range input {
		num, err := strconv.Atoi(fmt.Sprintf("%c", val))
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, num)
	}
	return ints
}
