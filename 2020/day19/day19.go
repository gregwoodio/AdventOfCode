package day19

import (
	"strings"
)

func solvePartOne(input []string) int {
	rules := make(map[string]string)
	index := 0

	for _, line := range input {
		if line == "" {
			break
		}

		rule := strings.Split(line, ": ")
		rules[rule[0]] = rule[1]

		index++
	}

	values := expand("0", rules)
	valid := make(map[string]bool)
	for _, val := range values {
		valid[val] = true
	}

	sum := 0
	for ; index < len(input); index++ {
		if valid[input[index]] {
			sum++
		}
	}

	return sum
}

func solvePartTwo(input []string) int {
	return -1
}

func expand(ruleID string, rules map[string]string) []string {
	if strings.HasPrefix(rules[ruleID], `"`) {
		// leaf node, return value
		return []string{
			strings.ReplaceAll(rules[ruleID], `"`, ``),
		}
	}

	parts := strings.Split(rules[ruleID], "|")
	ret := []string{}

	for _, part := range parts {
		ruleIDs := strings.Split(strings.Trim(part, " "), " ")
		ruleRet := []string{}

		for _, rID := range ruleIDs {
			newValues := expand(rID, rules)

			newRet := []string{}

			if len(ruleRet) > 0 {
				for _, r := range ruleRet {
					for _, nv := range newValues {
						newRet = append(newRet, r+nv)
					}
				}
			} else {
				newRet = newValues
			}
			// ruleRet = append(ruleRet, strings.Join(newValues, ""))

			ruleRet = newRet
		}

		ret = append(ret, ruleRet...)

		// newRet := []string{}
		// for _, r := range ret {
		// 	for _, rr := range ruleRet {
		// 		newRet = append(newRet, r+rr)
		// 	}
		// }

		// ret = newRet
	}

	return ret
}
