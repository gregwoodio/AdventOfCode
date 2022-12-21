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
	return -1
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
