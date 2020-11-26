package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println(partOne("input.txt"))
	fmt.Println(partTwo("input.txt"))
}

func partOne(input string) int {
	return readStream(input)
}

func partTwo(input string) int {
	file, _ := ioutil.ReadFile(input)

	var isEscaped, inGarbage bool
	var garbageCount int

	// fmt.Println(file)

	for _, c := range file {
		switch c {
		case '<':
			if isEscaped {
				isEscaped = false
			} else if !isEscaped && !inGarbage {
				inGarbage = true
			} else if !isEscaped && inGarbage {
				garbageCount++
			}
		case '>':
			if isEscaped {
				isEscaped = false
			} else if !isEscaped {
				inGarbage = false
			}
		case '!':
			if !isEscaped {
				isEscaped = true
			} else {
				isEscaped = false
			}
		default:
			if isEscaped {
				isEscaped = false
			} else if inGarbage {
				garbageCount++
			}
		}
	}
	return garbageCount
}

func readStream(input string) int {

	file, _ := ioutil.ReadFile(input)

	var isEscaped, inGarbage bool
	var level, score int

	for _, c := range file {

		switch c {
		case '{':
			if !inGarbage && !isEscaped {
				level++
			} else {
				isEscaped = false
			}
		case '}':
			if !inGarbage && !isEscaped {
				score += level
				level--
			} else {
				isEscaped = false
			}
		case '<':
			if !isEscaped {
				inGarbage = true
			} else {
				isEscaped = false
			}
		case '>':
			if !isEscaped {
				inGarbage = false
			} else {
				isEscaped = false
			}
		case '!':
			if !isEscaped {
				isEscaped = true
			} else {
				isEscaped = false
			}
		default:
			if isEscaped {
				isEscaped = false
			}
		}
	}

	return score
}
