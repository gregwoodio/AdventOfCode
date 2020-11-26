package day04

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

func solvePartOne(input string) int {
	return solve(input, "00000")
}

func solvePartTwo(input string) int {
	return solve(input, "000000")
}

// solve checks that the hash of a given input and some number starts with the given prefix.
// This function splits the work across a number of worker goroutines.
func solve(input, prefix string) int {
	workers := 25
	result := make(chan int, 1)
	done := make(chan bool, workers-1)
	var solution int

	for i := 0; i < workers; i++ {
		go solveWorker(i, workers, input, prefix, result, done)
	}

	idx := 0
workloop:
	for {
		select {
		case val := <-result:
			solution = val

			for w := 0; w < workers-1; w++ {
				done <- true
			}
			break workloop
		default:
		}

		idx = (idx + 1) % workers
	}

	return solution
}

// solveWorker checks the hash and returns the value that starts with the prefix if found
func solveWorker(start, increment int, key, prefix string, out chan int, done chan bool) {
	for i := start; true; i += increment {
		if checkHash(i, key, prefix) {
			out <- i
		}

		select {
		case <-done:
			return
		default:
		}
	}
}

func checkHash(num int, key, prefix string) bool {
	// append key number with current value
	value := fmt.Sprintf("%s%d", key, num)

	// find MD5 hash
	h := md5.New()
	io.WriteString(h, value)
	hashed := fmt.Sprintf("%x", h.Sum(nil))

	// check if it starts with the prefix
	regex := fmt.Sprintf("^%s.*", prefix)
	re := regexp.MustCompile(regex)
	return re.Match([]byte(hashed))
}
