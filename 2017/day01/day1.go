package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var input, _ = ioutil.ReadFile("input.txt")

	partOne(string(input))
	partTwo(string(input))
}

func partOne(input string) {
	solve(input, 1)
}

func partTwo(input string) {
	solve(input, len(input)/2)
}

func solve(input string, offset int) {
	captcha := 0
	for i, c := range input {
		num := int(c - '0')
		num2 := int(input[(i+offset)%len(input)] - '0')
		if num == num2 {
			captcha += num
		}
	}
	fmt.Printf("%d\n", captcha)
}
