package eval

import (
	"fmt"
	"unicode"
	"math"
)

type Mathematical struct {
	Equation string
	Output   float64
}

func NewMathematical(equation string) *Mathematical {
	return &Mathematical{Equation: equation, Output: 0}
}

// Token Type Getter
func (math *Mathematical) GetOutput() float64 {
	// Remove all spaces from the expression
	result, err := evalMathExpr(math.Equation)

	if err != nil {
		fmt.Println(err)
	}

	math.Output = result

	return result
}

func evalMathExpr(expr string) (float64, error) {
	// Remove all spaces from the expression
	expr = removeSpaces(expr)

	// Initialize the stack to hold the operands
	stack := []float64{}

	// Initialize the operator variable
	var op rune

	// Initialize the flag variable to keep track of whether the current digit
	// is the first digit of a new operand or not
	firstDigit := true

	// Iterate over the characters in the expression
	for _, c := range expr {
		switch c {
		case '+', '-', '*', '/', '^':
			// Push the current operand onto the stack and clear it
			value, err := toFloat(stack)
			if err != nil {
				return 0, err
			}
			if op == '+' {
				stack[len(stack)-1] += value
			} else if op == '-' {
				stack[len(stack)-1] -= value
			} else if op == '*' {
				stack[len(stack)-1] *= value
			} else if op == '^' {
				stack[len(stack)-1] = math.Pow(stack[len(stack)-1], value)
			} else if op == '/' {
				if value == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				stack[len(stack)-1] /= value
			}
			op = c

			// Reset the firstDigit flag
			firstDigit = true
		default:
			// Append the current digit to the current operand
			// This if statement has an issue with parenthesis
			// Give this prompt to ChatGPT for potential Fix
			// Prompt: It seems evalMathExpr has issue understanding 
			//and working with parenthesis, because the equation: (6/2)^3 
			// doesn't work but the equation: 6/2^3 does. Can you fix this about evalMathExpr
			if len(stack) > 0 && op == ' ' && !firstDigit {
				continue
			}else if len(stack) == 0 || op == '+' || op == '-' {
				stack = append(stack, float64(toDigit(c)))
				fmt.Println(stack, float64(toDigit(c)))
			} else if op == '*' {
				stack[len(stack)-1] *= float64(toDigit(c))
			} else if op == '^' {
				stack[len(stack)-1] = math.Pow(stack[len(stack)-1], float64(toDigit(c)))
			} else if op == '/' {
				if toDigit(c) == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				stack[len(stack)-1] /= float64(toDigit(c))
			}

			// Set the firstDigit flag to false
			firstDigit = false
			op = ' '
		}
	}

	// The final value on the stack is the result of the expression
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression 2")
	}
	return stack[0], nil
}

// removeSpaces removes all spaces from a string.
func removeSpaces(str string) string {
	var result string
	for _, c := range str {
		if !unicode.IsSpace(c) {
			result += string(c)
		}
	}
	return result
}

// toDigit converts a rune representing a digit to an integer.
func toDigit(c rune) int {
	return int(c - '0')
}

// toFloat converts a slice of digits to a float64.
func toFloat(digits []float64) (float64, error) {
	var result float64
	for _, d := range digits {
		result = result*10 + d
	}
	return result, nil
}
