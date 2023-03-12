package lab2

import (
	"fmt"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts

func PrefixToInfix(input string) (string, error) {
	var stack []string
	const operators = "+-*/^"
	const error = "invalid input expression"
	if len(input) == 0 {
		return "", fmt.Errorf(error)
	}

	var inputs = strings.Split(input, " ")
	symbol, expression1, expression2 := "", "", ""

	for i := len(inputs) - 1; i >= 0; i-- {
		symbol = inputs[i]

		if !strings.Contains(operators, symbol) {
			stack = append(stack, symbol)
		} else {
			if len(stack) < 2 {
				return "", fmt.Errorf(error)
			}

			expression1, stack = Pop(stack)
			expression2, stack = Pop(stack)

			var new = Join(expression1, expression2, symbol)
			stack = append(stack, new)
		}
	}

	var res = ""
	res, stack = Pop(stack)

	if len(stack) > 0 {
		return "", fmt.Errorf(error)
	}

	return res, nil
}

func Pop(str []string) (string, []string) {
	if len(str) == 0 {
	  return "", []string{}
	}

	return str[len(str)-1], str[:len(str)-1]
  }

  func Join(expression1 string, expression2 string, symbol string) (string){
	if strings.Contains(expression1, " ") && strings.Contains("*/^", symbol) {
	  expression1 = "(" + expression1 + ")"
	}
	if strings.Contains(expression2, " ") && strings.Contains("*/^", symbol) {
	  expression2 = "(" + expression2 + ")"
	}

	var newExpression = expression1 + " " + symbol + " " + expression2
	return newExpression
  }