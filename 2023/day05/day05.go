package day05

import (
	"regexp"
	"strconv"
	"strings"
)

type almanacMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func solvePartOne(input []string) int {
	seeds := []int{}
	for _, str := range strings.Split(input[0][len("seeds: "):], " ") {
		seed, _ := strconv.Atoi(str)
		seeds = append(seeds, seed)
	}

	// updateable slice of values to apply maps to
	values := []int{}
	values = append(values, seeds...)

	lineIndex := 3
	for {
		if lineIndex >= len(input) {
			break
		}

		// parse next group of lines into almanac maps
		maps := []almanacMap{}
		for ; lineIndex < len(input) && input[lineIndex] != ""; lineIndex++ {
			maps = append(maps, parseAlmanacMap(input[lineIndex]))
		}

		// skip blank line
		lineIndex++

		for i, val := range values {
			for _, m := range maps {
				if val >= m.sourceRangeStart && val <= m.sourceRangeStart+m.rangeLength {

					// update values[i] as per map and break
					d := val - m.sourceRangeStart
					values[i] = m.destinationRangeStart + d
					break
				}
			}
		}
	}

	lowest := values[0]
	for _, val := range values {
		lowest = min(lowest, val)
	}

	return lowest
}

func solvePartTwo(input []string) int {
	return -1
}

func parseAlmanacMap(line string) almanacMap {
	regex := `(?P<Dest>\d+) (?P<Src>\d+) (?P<Range>\d+)`
	compiled := regexp.MustCompile(regex)
	match := compiled.FindStringSubmatch(line)

	var dest, src, rng int

	for i, subexpName := range compiled.SubexpNames() {
		if i == 0 || subexpName == "" {
			continue
		}

		if i > len(match) {
			break
		}

		switch subexpName {
		case "Dest":
			dest, _ = strconv.Atoi(match[i])
		case "Src":
			src, _ = strconv.Atoi(match[i])
		case "Range":
			rng, _ = strconv.Atoi(match[i])
		}
	}

	return almanacMap{
		destinationRangeStart: dest,
		sourceRangeStart:      src,
		rangeLength:           rng,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
