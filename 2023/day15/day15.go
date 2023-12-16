package day15

import (
	"strconv"
	"strings"
)

type lens struct {
	label       string
	focalLength int
	remove      bool
	hash        int
}

func solvePartOne(input []string) int {
	sum := 0

	parts := strings.Split(input[0], ",")
	for _, part := range parts {
		sum += hash(part)
	}

	return sum
}

func solvePartTwo(input []string) int {
	sum := 0

	boxes := make([][]lens, 256)

	instructions := strings.Split(input[0], ",")
	for _, inst := range instructions {
		lens := makeLens(inst)

		if lens.remove {
			box := boxes[lens.hash]
			for i, l := range box {
				if l.label == lens.label {
					boxes[lens.hash] = append(box[:i], box[i+1:]...)
					break
				}
			}
		} else {
			found := false
			for i, l := range boxes[lens.hash] {
				if l.label == lens.label {
					found = true
					boxes[lens.hash][i].focalLength = lens.focalLength
					break
				}
			}

			if !found {
				boxes[lens.hash] = append(boxes[lens.hash], lens)
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focalLength
		}
	}

	return sum
}

func hash(input string) int {
	curr := 0

	for _, ch := range input {
		curr += int(ch)
		curr *= 17
		curr %= 256
	}
	return curr
}

func makeLens(input string) lens {
	index := 0
	remove := false
	for i, ch := range input {
		if ch == '=' {
			index = i
			break
		}
		if ch == '-' {
			remove = true
			index = i
			break
		}
	}

	label := input[:index]
	h := hash(label)

	if !remove {
		f := input[index+1:]
		focalLength, _ := strconv.Atoi(f)

		return lens{
			label:       label,
			focalLength: focalLength,
			remove:      remove,
			hash:        h,
		}
	}

	return lens{
		label:  label,
		remove: remove,
		hash:   h,
	}
}
