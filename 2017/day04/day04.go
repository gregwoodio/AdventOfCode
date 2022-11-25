package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gregwoodio/adventofcode/m/aocutil"
)

func main() {
	input := aocutil.ReadStringsFromFile("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input []string) int {
	return checkWords(input, false)
}

func partTwo(input []string) int {
	return checkWords(input, true)
}

func checkWords(input []string, isPartTwo bool) int {
	count := 0

	for _, passphrase := range input {
		words := make(map[string]bool)
		parts := strings.Split(passphrase, " ")
		valid := true

		for _, word := range parts {

			if isPartTwo {
				word = sortWord(word)
			}

			if words[word] {
				valid = false
				break
			}
			words[word] = true
		}

		if valid {
			count++
		}
	}

	return count
}

func sortWord(word string) string {
	w := strings.Split(word, "")
	sort.Strings(w)
	return strings.Join(w, "")
}
