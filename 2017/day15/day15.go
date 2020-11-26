package main

import "fmt"

func main() {
	// fmt.Println(partOne(65, 8921))
	fmt.Println(partOne(618, 814))

	// fmt.Println(partTwo(65, 8921))
	fmt.Println(partTwo(618, 814))
}

func partOne(genAStart, genBStart int) int {
	genA, genB := genAStart, genBStart
	factorA, factorB := 16807, 48271
	divisor := 2147483647
	count := 0

	for round := 0; round < 40000000; round++ {
		genA = (genA * factorA) % divisor
		genB = (genB * factorB) % divisor

		// fmt.Printf("%10d\t%10d\n", genA, genB)
		if (genA & 0xffff) == (genB & 0xffff) {
			count++
		}
	}
	return count
}

func partTwo(genAStart, genBStart int) int {

	var judgeAList []int
	var judgeBList []int

	genA, genB := genAStart, genBStart
	factorA, factorB := 16807, 48271
	divisor := 2147483647

	for len(judgeAList) < 5000000 {
		genA = (genA * factorA) % divisor
		if genA%4 == 0 {
			judgeAList = append(judgeAList, genA)
		}
	}

	for len(judgeBList) < 5000000 {
		genB = (genB * factorB) % divisor
		if genB%8 == 0 {
			judgeBList = append(judgeBList, genB)
		}
	}

	count := 0
	for round := 0; round < 5000000; round++ {
		if (judgeAList[round] & 0xffff) == (judgeBList[round] & 0xffff) {
			count++
		}
	}

	return count
}
