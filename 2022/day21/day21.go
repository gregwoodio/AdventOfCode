package day21

import (
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	id        string
	num       *int
	operation func(a, b int) int
	op        string
	operand   []string
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func solvePartOne(input []string) int {
	monkeys := parse(input)

	return calc(monkeys["root"], monkeys)
}

func solvePartTwo(input []string) int {
	monkeys := parse(input)
	monkeys["humn"].num = nil
	root := monkeys["root"]
	// check equality of operands
	m1, ok := monkeys[root.operand[0]]
	if !ok {
		panic(fmt.Sprintf("Missing monkey %s", root.operand[0]))
	}
	m2, ok := monkeys[root.operand[1]]
	if !ok {
		panic(fmt.Sprintf("Missing monkey %s", root.operand[1]))
	}

	m1Path := findHumnPath(m1, monkeys)
	m2Path := findHumnPath(m2, monkeys)
	var result int
	var op string
	if m1Path {
		calc(m2, monkeys)
		result = *m2.num
		op = m1.op
		m1, m2 = monkeys[m1.operand[0]], monkeys[m1.operand[1]]
	} else {
		calc(m1, monkeys)
		result = *m1.num
		op = m2.op
		m1, m2 = monkeys[m2.operand[0]], monkeys[m2.operand[1]]
	}

	humn := monkeys["humn"]
	for {
		// find humn path
		m1Path = findHumnPath(m1, monkeys)
		m2Path = findHumnPath(m2, monkeys)

		// reverse the operation for the humn path
		if m1Path {
			calc(m2, monkeys)
			switch op {
			case "/":
				sum := mul(*m2.num, result)
				m1.num = &sum
			case "-":
				sum := add(*m2.num, result)
				m1.num = &sum
			case "*":
				sum := div(result, *m2.num)
				m1.num = &sum
			case "+":
				sum := sub(result, *m2.num)
				m1.num = &sum
			}
			result = *m1.num
			op = m1.op

			if humn.num != nil {
				return *humn.num
			}

			m1, m2 = monkeys[m1.operand[0]], monkeys[m1.operand[1]]

		} else if m2Path {
			calc(m1, monkeys)
			switch op {
			case "/":
				sum := div(*m1.num, result)
				m2.num = &sum
			case "-":
				sum := sub(*m1.num, result)
				m2.num = &sum
			case "*":
				sum := div(result, *m1.num)
				m2.num = &sum
			case "+":
				sum := sub(result, *m1.num)
				m2.num = &sum
			}
			result = *m2.num
			op = m2.op

			if humn.num != nil {
				return *humn.num
			}

			m1, m2 = monkeys[m2.operand[0]], monkeys[m2.operand[1]]
		}
	}
}

func findHumnPath(m *monkey, monkeys map[string]*monkey) bool {
	if m.id == "humn" {
		return true
	}

	if len(m.operand) > 0 {
		if m.operand[0] == "humn" || m.operand[1] == "humn" {
			return true
		} else {
			m1 := findHumnPath(monkeys[m.operand[0]], monkeys)
			m2 := findHumnPath(monkeys[m.operand[1]], monkeys)

			return m1 || m2
		}
	}

	return false
}

func calc(m *monkey, monkeys map[string]*monkey) int {
	if m.num != nil {
		return *m.num
	}

	m1, ok := monkeys[m.operand[0]]
	if !ok {
		panic(fmt.Sprintf("Missing monkey %s", m.operand[0]))
	}
	m2, ok := monkeys[m.operand[1]]
	if !ok {
		panic(fmt.Sprintf("Missing monkey %s", m.operand[1]))
	}

	num := m.operation(calc(m1, monkeys), calc(m2, monkeys))
	m.num = &num

	return num
}

func parse(input []string) map[string]*monkey {
	monkeys := map[string]*monkey{}

	for _, line := range input {
		parts := strings.Split(line, ": ")
		m := &monkey{
			id: parts[0],
		}

		if strings.Contains(parts[1], " ") {
			// operation
			ops := strings.Split(parts[1], " ")
			m.operand = []string{ops[0], ops[2]}
			m.op = ops[1]
			switch ops[1] {
			case "+":
				m.operation = add
			case "-":
				m.operation = sub
			case "*":
				m.operation = mul
			case "/":
				m.operation = div
			}
		} else {
			// number
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err.Error())
			}
			m.num = &val
		}

		monkeys[m.id] = m
	}

	return monkeys
}
