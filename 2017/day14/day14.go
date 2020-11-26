package main

import (
	"fmt"
	"strconv"
	"strings"
)

var bytes = [...]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

func main() {
	input := "oundnydw"
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input string) int {
	hashes := solve(input)

	count := 0
	for _, h := range hashes {
		for _, l := range h {
			num, _ := strconv.ParseInt(string(l), 16, 32)
			count += bytes[num]
		}
	}

	return count
}

func partTwo(input string) int {
	hashes := solve(input)

	binary := make([][]bool, 128)
	for i, hash := range hashes {
		for _, l := range hash {
			num, _ := strconv.ParseInt(string(l), 16, 32)
			binary[i] = append(binary[i], num&8 > 0)
			binary[i] = append(binary[i], num&4 > 0)
			binary[i] = append(binary[i], num&2 > 0)
			binary[i] = append(binary[i], num&1 > 0)
		}
	}

	var zones [][]int
	for i := 0; i < 128; i++ {
		zones = append(zones, make([]int, 128))
	}

	counter := 0

	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if binary[x][y] && zones[x][y] == 0 {
				counter++
				checkSurroundingCells(x, y, counter, binary, zones)
			}
		}
	}

	return counter
}

func checkSurroundingCells(x int, y int, counter int, binary [][]bool, zones [][]int) {
	if x < 0 || x > 127 || y < 0 || y > 127 {
		return
	}

	if !binary[x][y] {
		return
	}

	if zones[x][y] > 0 {
		return
	}

	zones[x][y] = counter

	checkSurroundingCells(x+1, y, counter, binary, zones)
	checkSurroundingCells(x-1, y, counter, binary, zones)
	checkSurroundingCells(x, y+1, counter, binary, zones)
	checkSurroundingCells(x, y-1, counter, binary, zones)
}

func solve(input string) []string {
	var hashes = make([]string, 0)

	for i := 0; i < 128; i++ {
		hashRound := input + "-" + strconv.Itoa(i)
		h := hash([]byte(hashRound), 256)
		hashes = append(hashes, h)
	}

	return hashes
}

func hash(input []byte, listSize int) string {
	listEnding := []byte{17, 31, 73, 47, 23}
	input = append(input, listEnding...)

	list := makeList(listSize)
	currPos, skipSize := 0, 0

	for round := 0; round < 64; round++ {

		for _, length := range input {
			list = reverseSubslice(list, currPos, int(length))
			currPos = (currPos + int(length) + skipSize) % len(list)
			skipSize++
		}
	}

	var denseHash []string
	for i := 0; i < len(list); i += 16 {
		var val int
		for j := i; j < i+16; j++ {
			val ^= list[j]
		}
		denseHash = append(denseHash, fmt.Sprintf("%02x", val))
	}

	return strings.Join(denseHash, "")
}

func reverseSubslice(input []int, startingIndex int, length int) []int {
	target := (startingIndex + length - 1) % len(input)
	for i := 0; i < length/2; i++ {
		tmp := input[(startingIndex+i)%len(input)]
		input[(startingIndex+i)%len(input)] = input[target]
		input[target] = tmp
		target = (target - 1) % len(input)
		if target < 0 {
			target = len(input) - 1
		}
	}
	return input
}

func makeList(listSize int) []int {
	slice := make([]int, 0)
	for i := 0; i < listSize; i++ {
		slice = append(slice, i)
	}
	return slice
}
