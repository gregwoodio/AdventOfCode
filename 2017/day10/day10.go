package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gregwoodio/aocutil"
)

func main() {
	input := aocutil.ReadAndSplitInts("input.txt", ",")[0]
	fmt.Println(partOne(input, 256))

	line, _ := ioutil.ReadFile("input.txt")
	fmt.Println(partTwo(line, 256))
}

func partOne(lengths []int, listSize int) int {
	list := makeList(listSize)
	currPos, skipSize := 0, 0

	for _, length := range lengths {
		list = reverseSubslice(list, currPos, length)
		currPos = (currPos + length + skipSize) % len(list)
		skipSize++
	}

	return list[0] * list[1]
}

func partTwo(input []byte, listSize int) string {
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
		denseHash = append(denseHash, fmt.Sprintf("%x", val))
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
