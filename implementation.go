package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func EvaluatePrefix(expression string) (int, error) {
	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, errors.New("empty expression")
	}

	stack := []int{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("insufficient operands for operator " + token)
			}

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res, err := applyOperator(token, a, b)
			if err != nil {
				return 0, err
			}

			stack = append(stack, res)
		} else {
			return 0, errors.New("invalid token: " + token)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression: stack has more than one element")
	}

	return stack[0], nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func applyOperator(op string, a, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	case "^":
		if b < 0 {
			return 0, errors.New("negative exponent")
		}
		result := 1
		for i := 0; i < b; i++ {
			result *= a
		}
		return result, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func main() {
	expr := "+ 5 * - 4 2 ^ 3 2"
	result, err := EvaluatePrefix(expr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
