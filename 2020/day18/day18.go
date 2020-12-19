package day18

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePartOne(input []string) int {
	sum := 0
	for _, line := range input {
		val, _ := strconv.Atoi(evaluate(line))
		sum += val
	}

	return sum
}

func solvePartTwo(input []string) int {
	sum := 0
	for _, line := range input {
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")

		sum += evaluate2(line)
	}

	return sum
}

func evaluate(expression string) string {
	paren := strings.Index(expression, "(")
	if paren > -1 {
		closeParen := paren + 1
		depth := 1
		for depth > 0 {
			if expression[closeParen] == '(' {
				depth++
			}
			if expression[closeParen] == ')' {
				depth--
			}
			closeParen++
		}

		subexp := evaluate(expression[paren+1 : closeParen-1])
		expression = string(expression[:paren]) + subexp + expression[closeParen:]

		return evaluate(expression)
	}

	// no parentheses
	tokens := strings.Split(expression, " ")
	tokens = jankyShuntingYard(tokens)
	val, _ := strconv.Atoi(tokens[0])

	return strconv.Itoa(val)
}

// janky because all operators have the same precedence
func jankyShuntingYard(tokens []string) []string {
	s := []string{}
	q := []string{}

	for _, token := range tokens {
		if token == "+" || token == "*" {
			s = append([]string{token}, s...)
		} else {
			if len(s) > 0 {
				op := s[0]
				s = s[1:]

				operand := q[0]
				q = q[:len(q)-1]

				n1, _ := strconv.Atoi(operand)
				n2, _ := strconv.Atoi(token)

				switch op {
				case "+":
					q = append(q, fmt.Sprintf("%d", n1+n2))
					break
				case "*":
					q = append(q, fmt.Sprintf("%d", n1*n2))
					break
				}
			} else {
				q = append(q, token)
			}
		}
	}

	for len(s) > 0 {
		t := s[0]
		q = append(q, t)
		s = s[1:]
	}

	return q
}

func evaluate2(input string) int {
	tokens := strings.Split(input, " ")
	tokens = backwardShuntingYard(tokens)
	return evaluateRPN(tokens)
}

func backwardShuntingYard(tokens []string) []string {
	q := []string{}
	s := []string{}

	for _, token := range tokens {
		if token == "(" {
			s = append(s, token)
		} else if token == ")" {
			var op string
			for {
				op = s[len(s)-1]
				s = s[:len(s)-1]
				if op == "(" {
					break
				}
				q = append(q, op)
			}
		} else if token == "+" {
			for len(s) > 0 {
				op := s[len(s)-1]
				if op == "*" || op == "(" || op == "+" {
					break
				}

				s = s[:len(s)-1]
				q = append(q, op)
			}

			s = append(s, token)
		} else if token == "*" {
			for len(s) > 0 {
				op := s[len(s)-1]
				if op == "(" || op == "*" {
					break
				}

				s = s[:len(s)-1]
				q = append(q, op)
			}

			s = append(s, token)
		} else {
			q = append(q, token)
		}
	}

	for len(s) > 0 {
		q = append(q, s[len(s)-1])
		s = s[:len(s)-1]
	}

	return q
}

func evaluateRPN(tokens []string) int {
	s := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			s[len(s)-2] += s[len(s)-1]
			s = s[:len(s)-1]
			break
		case "*":
			s[len(s)-2] *= s[len(s)-1]
			s = s[:len(s)-1]
			break
		default:
			val, err := strconv.Atoi(token)
			if err != nil {
				fmt.Errorf("could not parse to int: '%s'", token)
			}

			s = append(s, val)
			break
		}
	}

	return s[0]
}
