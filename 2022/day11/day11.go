package day11

import (
	"sort"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

type monkey struct {
	id           int
	items        []int
	operation    func(int, int) int
	oper1        *int
	oper2        *int
	divisor      int
	trueIndex    int
	falseIndex   int
	inspectCount int
}

func solvePartOne(input []string) int {
	monkeys := parseInput(input)
	simulate(monkeys, false)

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})

	monkeyBusiness := monkeys[0].inspectCount * monkeys[1].inspectCount

	return monkeyBusiness
}

func solvePartTwo(input []string) int {
	monkeys := parseInput(input)
	simulate(monkeys, true)

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})

	monkeyBusiness := monkeys[0].inspectCount * monkeys[1].inspectCount

	return monkeyBusiness
}

func parseInput(input []string) []*monkey {
	var m *monkey
	monkeys := []*monkey{}

	for _, line := range input {
		line := strings.Trim(line, " ")
		if strings.HasPrefix(line, "Monkey") {
			if m != nil {
				monkeys = append(monkeys, m)
			}

			// cheat a bit; I know that the number of monkeys is a single digit
			id, err := strconv.Atoi(line[7:8])
			if err != nil {
				panic(err)
			}

			m = &monkey{
				id:           id,
				items:        []int{},
				inspectCount: 0,
			}
		} else if strings.HasPrefix(line, "Starting items: ") {
			itemStrs := strings.Split(line[16:], ", ")
			for _, itemStr := range itemStrs {
				item, err := strconv.Atoi(itemStr)
				if err != nil {
					panic(err)
				}
				m.items = append(m.items, item)
			}
		} else if strings.HasPrefix(line, "Operation: ") {
			operStrs := strings.Split(line[17:], " ")
			if operStrs[0] != "old" {
				val, err := strconv.Atoi(operStrs[0])
				if err != nil {
					panic(err)
				}

				m.oper1 = &val
			}
			if operStrs[1] == "+" {
				m.operation = add
			} else {
				m.operation = mul
			}
			if operStrs[2] != "old" {
				val, err := strconv.Atoi(operStrs[2])
				if err != nil {
					panic(err)
				}

				m.oper2 = &val
			}
		} else if strings.HasPrefix(line, "Test: divisible by") {
			parts := strings.Split(line, " ")
			val, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				panic(err)
			}

			m.divisor = val
		} else if strings.HasPrefix(line, "If true:") {
			parts := strings.Split(line, " ")
			val, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				panic(err)
			}

			m.trueIndex = val
		} else if strings.HasPrefix(line, "If false:") {
			parts := strings.Split(line, " ")
			val, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				panic(err)
			}

			m.falseIndex = val
		}
	}

	if m != nil {
		monkeys = append(monkeys, m)
	}

	return monkeys
}

func simulate(monkeys []*monkey, isPartTwo bool) {
	var prod int = 1
	for _, m := range monkeys {
		prod *= m.divisor
	}

	var rounds int
	if isPartTwo {
		rounds = 10000
	} else {
		rounds = 20
	}
	for round := 0; round < rounds; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.inspectCount++

				var oper1, oper2 int
				if m.oper1 == nil {
					oper1 = item
				} else {
					oper1 = *m.oper1
				}
				if m.oper2 == nil {
					oper2 = item
				} else {
					oper2 = *m.oper2
				}
				newItem := m.operation(oper1, oper2)

				// divide by 3
				if !isPartTwo {
					newItem /= 3
				}

				newItem %= prod
				if newItem%m.divisor == 0 {
					monkeys[m.trueIndex].items = append(monkeys[m.trueIndex].items, newItem)
				} else {
					monkeys[m.falseIndex].items = append(monkeys[m.falseIndex].items, newItem)
				}
			}
			m.items = []int{}
		}
	}
}
