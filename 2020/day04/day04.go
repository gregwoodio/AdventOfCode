package day04

import (
	"regexp"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	return solve(input, false)
}

func solvePartTwo(input []string) int {
	return solve(input, true)
}

func solve(input []string, isPartTwo bool) int {
	passport := make(map[string]string)
	validPassports := 0

	for _, line := range input {
		if line == "" {
			if !isPartTwo && checkPassport(passport) || checkPassportPart2(passport) {
				validPassports++
			}

			passport = make(map[string]string)
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			parts := strings.Split(field, ":")
			passport[parts[0]] = parts[1]
		}
	}

	if !isPartTwo && checkPassport(passport) || checkPassportPart2(passport) {
		validPassports++
	}

	return validPassports
}

func checkPassport(p map[string]string) bool {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, field := range required {
		if _, ok := p[field]; !ok {
			return false
		}
	}

	return true
}

func checkPassportPart2(p map[string]string) bool {
	// birth year
	val, ok := p["byr"]
	if !ok {
		return false
	}
	byr, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// issue year
	val, ok = p["iyr"]
	if !ok {
		return false
	}
	iyr, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// expiration year
	val, ok = p["eyr"]
	if !ok {
		return false
	}
	eyr, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// height
	val, ok = p["hgt"]
	if !ok {
		return false
	}
	heightRegex := regexp.MustCompile(`(?P<value>[0-9]+)(?P<unit>(cm|in))`)
	heightParts := heightRegex.FindStringSubmatch(val)
	if len(heightParts) < 2 {
		return false
	}
	heightValue, err := strconv.Atoi(heightParts[1])
	if err != nil {
		return false
	}
	if heightParts[2] == "cm" {
		if heightValue < 150 || heightValue > 193 {
			return false
		}
	} else if heightParts[2] == "in" {
		if heightValue < 59 || heightValue > 76 {
			return false
		}
	} else {
		// wrong unit
		return false
	}

	// hair color
	val, ok = p["hcl"]
	if !ok {
		return false
	}
	hclRegex := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	if !hclRegex.Match([]byte(val)) {
		return false
	}

	// eye colour
	val, ok = p["ecl"]
	if !ok {
		return false
	}
	eclRegex := regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
	if !eclRegex.Match([]byte(val)) {
		return false
	}

	// passport ID
	val, ok = p["pid"]
	if !ok {
		return false
	}
	pidRegex := regexp.MustCompile(`^[0-9]{9}$`)
	if !pidRegex.Match([]byte(val)) {
		return false
	}

	// Country ID is not required

	return true
}
