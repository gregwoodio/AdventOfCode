package day25

func solvePartOne(cardPublicKey, doorPublicKey int) int {

	cardLoopSize := findLoopSize(7, cardPublicKey)
	doorLoopSize := findLoopSize(7, doorPublicKey)

	cardEncryptionKey := transform(doorPublicKey, cardLoopSize)
	doorEncryptionKey := transform(cardPublicKey, doorLoopSize)

	if cardEncryptionKey == doorEncryptionKey {
		return cardEncryptionKey
	}

	return -1
}

func findLoopSize(subject, target int) int {
	val, loopSize := 1, 0

	for ; val != target; loopSize++ {
		val *= subject
		val %= 20201227
	}

	return loopSize
}

func transform(subject, loopSize int) int {
	val := 1

	for i := 0; i < loopSize; i++ {
		val *= subject
		val %= 20201227
	}

	return val
}
