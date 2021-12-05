package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	input := utils.Load("real").ToStringSlice()

	fmt.Println("-----Part 1-----")
	fmt.Println("Sum of expression: ", evaluateSum(input, part1))
	fmt.Println("-----Part 2-----")
	fmt.Println("Sum of expression: ", evaluateSum(input, part2))
}

type tokenType int

const (
	Operand tokenType = iota
	Operator
)

type Token struct {
	Type  tokenType
	Value interface{}
}

type OperatorPrecedence map[string]int

var part1 = OperatorPrecedence{"+": 0, "*": 0}
var part2 = OperatorPrecedence{"+": 1, "*": 0}

func evaluate(expression string, precedence OperatorPrecedence) int {
	return evaluateParsed(parse(scan(expression), precedence))
}

func evaluateSum(expressions []string, precedence OperatorPrecedence) int {
	sum := 0
	for _, line := range expressions {
		sum += evaluate(line, precedence)
	}

	return sum
}

func scan(input string) []string {
	retVal := make([]string, 0)

	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		retVal = append(retVal, s.TokenText())
	}
	return retVal
}

// Implements Dijkstra's shunting-yard algorithm to parse the expression into RPN tokens
func parse(tokens []string, operators map[string]int) []Token {
	var result []Token
	var operatorStack []string

	for _, token := range tokens {
		// Is this token an operand (aka number)?
		if num, err := strconv.Atoi(token); err == nil {
			result = append(result, Token{Type: Operand, Value: num})
			continue
		}

		if token == "(" {
			operatorStack = append(operatorStack, token)
			continue
		}

		if token == ")" {
			// Pop operators off the stack until we find a matching open paren
			for len(operatorStack) > 0 {
				op := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if op == "(" {
					break
				}

				result = append(result, Token{Type: Operator, Value: op})
			}
			continue
		}

		// We have an operator
		// Determine its priority
		currentTokenPriority, ok := operators[token]
		if !ok {
			panic("unknown operator: " + token)
		}

		for len(operatorStack) > 0 {
			// Check top of the operator stack
			op := operatorStack[len(operatorStack)-1]
			if op == "(" {
				break
			}

			poppedOpPriority := operators[op]
			if currentTokenPriority <= poppedOpPriority {
				// Pop it off the stack and onto our result token set
				operatorStack = operatorStack[:len(operatorStack)-1]
				result = append(result, Token{Type: Operator, Value: op})
			} else {
				break
			}
		}

		operatorStack = append(operatorStack, token)
	}

	// Handle any remaining operators
	for len(operatorStack) > 0 {
		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		result = append(result, Token{Type: Operator, Value: op})
	}

	return result
}

func evaluateParsed(tokens []Token) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		// Push operands (numbers) onto the stack as we find them
		if token.Type == Operand {
			stack = append(stack, token.Value.(int))
			continue
		}

		// Pop two operands off the stack and apply them
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		// Perform the operation and push the result back onto the stack
		stack = append(stack, doOperation(token.Value.(string), a, b))
	}

	return stack[0]
}

func doOperation(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	}

	panic("unknown operator: " + op)
}
