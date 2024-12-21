// Copyright (c) 2024 Rui Barroso
// This code is licensed under the MIT License.

package model

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Sum, Subtraction, Multiplication, Division, Exponent, Root, Logarithm, OpenParen, CloseParen define the string constants used for representing basic mathematical operations and parentheses in the equation.
const (
	Sum            = "+" // Represents the addition operation.
	Subtraction    = "-" // Represents the subtraction operation.
	Multiplication = "*" // Represents the multiplication operation.
	Division       = "/" // Represents the division operation.
	Exponent       = "^" // Represents the exponentiation operation.
	Root           = "r" // Represents the root operation (e.g., square root).
	Logarithm      = "l" // Represents the logarithm operation.
	OpenParen      = "(" // Represents an opening parenthesis.
	CloseParen     = ")" // Represents a closing parenthesis.
)

// Equation represents a mathematical equation as a string.
type Equation struct {
	Equation string // The equation string to be parsed and evaluated.
}

// ParseEquation parses the input equation into individual components (numbers and operators) based on a given toIgnore.
// The toIgnore determines which characters to ignore during parsing.
// It returns a slice of strings representing the equation members.
func (eq *Equation) ParseEquation(toIgnore string) []string {
	members := make([]string, 0, len(eq.Equation))
	member := ""
	for _, c := range eq.Equation {
		if strings.ContainsRune(toIgnore, c) {
			continue
		}
		if strings.ContainsRune("+-*/^rl()", c) {
			if member != "" {
				members = append(members, member, string(c))
			} else {
				members = append(members, string(c))
			}

			member = ""

		} else {
			member += string(c)
		}
	}
	if member != "" {
		members = append(members, member)
	}

	return members
}

// Evaluate evaluates the parsed equation by processing each operation according to the order of operations (PEMDAS).
// It first processes parentheses, then exponentiation, followed by multiplication, and finally addition/subtraction.
// The result of each step is passed to the next, and the final result is returned.
func Evaluate(members []string) (float64, error) {

	if !CheckEquation(members) {
		return 0.0, errors.New("equation format error")
	}
	var err error
	members, err = EvaluateParentesis(members)
	if err != nil {
		return 0.0, err
	}

	members, err = EvaluateExponentiation(members)
	if err != nil {
		return 0.0, err
	}

	members, err = EvaluateMultiplication(members)
	if err != nil {
		return 0.0, err
	}

	return EvaluateAddition(members)
}

// EvaluateParentesis handles the evaluation of expressions inside parentheses by recursively resolving
// the sub-expressions within them. Parentheses are processed from innermost to outermost.
func EvaluateParentesis(members []string) ([]string, error) {
	membersAux := make([]string, 0, len(members))
	for i := len(members) - 1; i >= 0; i-- {

		if members[i] == CloseParen {
			aux := 0
			for j := i - 1; j >= 0; j-- {
				if members[j] == CloseParen {
					aux++
				} else if members[j] == OpenParen {
					if aux > 0 {
						aux--
						continue
					}
					res, err := Evaluate(members[j+1 : i])
					if err != nil {
						return nil, errors.New("evaluating parentises error")
					}
					membersAux = append(membersAux, fmt.Sprintf("%f", res))
					i = j
					break
				}
			}
		} else {
			membersAux = append(membersAux, members[i])
		}

	}

	// Reverse the order to preserve original sequence after processing.
	slices.Reverse(membersAux)
	return membersAux, nil
}

// EvaluateExponentiation processes exponentiation (^), roots (r), and logarithms (l) in the equation. It applies the corresponding operations
// in the order of their occurrence and returns the updated equation.
func EvaluateExponentiation(members []string) ([]string, error) {
	membersAux := make([]string, 0, len(members))
	for i := 0; i < len(members); i++ {
		switch members[i] {
		case Exponent:
			var err error
			membersAux, err = Operation(i, members, membersAux, func(a float64, b float64) float64 { return math.Pow(a, b) })
			if err != nil {
				return nil, err
			}
			i++

		case Root:
			var err error
			membersAux, err = Operation(i, members, membersAux, func(a float64, b float64) float64 { return math.Pow(a, 1/b) })
			if err != nil {
				return nil, err
			}
			i++

		case Logarithm:
			var err error
			membersAux, err = Operation(i, members, membersAux, func(a float64, b float64) float64 { return math.Log(a) / math.Log(b) })
			if err != nil {
				return nil, err
			}
			i++

		default:
			membersAux = append(membersAux, members[i])
		}

	}

	return membersAux, nil
}

// EvaluateMultiplication processes multiplication (*) and division (/) operations in the equation.
// It applies the corresponding operations in the order of their occurrence and returns the updated equation.
func EvaluateMultiplication(members []string) ([]string, error) {
	membersAux := make([]string, 0, len(members))
	for i := 0; i < len(members); i++ {
		switch members[i] {
		case Multiplication:
			var err error
			membersAux, err = Operation(i, members, membersAux, func(a float64, b float64) float64 { return a * b })
			if err != nil {
				return nil, err
			}
			i++

		case Division:
			var err error
			membersAux, err = Operation(i, members, membersAux, func(a float64, b float64) float64 { return a / b })
			if err != nil {
				return nil, err
			}
			i++

		default:
			membersAux = append(membersAux, members[i])
		}

	}

	return membersAux, nil
}

// EvaluateAddition processes addition (+) and subtraction (-) operations in the equation.
// It applies the corresponding operations in the order they occur and returns the final result of the equation.
func EvaluateAddition(members []string) (float64, error) {
	sum, err := ParseFloat(members[0])
	if err != nil {
		return 0.0, err
	}
	for i := 1; i < len(members); i += 2 {
		if members[i] == Subtraction {
			res, err := ParseFloat(members[i+1])
			if err != nil {
				return 0.0, err
			}
			sum -= res
		} else {
			res, err := ParseFloat(members[i+1])
			if err != nil {
				return 0.0, err
			}
			sum += res
		}

	}
	return sum, nil
}

// CheckEquation checks the validity of the equation. It ensures:
// 1. Proper placement of operators.
// 2. Correct number of operands and operators.
// 3. Balanced parentheses.
func CheckEquation(members []string) bool {
	parentsesCount := 0
	for i, c := range members {

		if len(c) == 1 && strings.Contains("+-*/^rl", c) {
			if i == 0 || i == len(members)-1 {
				return false
			}
			_, err := ParseFloat(members[i-1])
			if err != nil && !IsParenteses(members[i-1]) {
				return false
			}
			_, err = ParseFloat(members[i+1])
			if err != nil && !IsParenteses(members[i+1]) {
				return false
			}

		} else if c == OpenParen {
			parentsesCount += 1
		} else if c == CloseParen {
			parentsesCount -= 1
		}

		if parentsesCount < 0 {
			return false
		}
	}
	return parentsesCount == 0
}

// ParseFloat converts a string to a float64, returning an error if the conversion fails.
func ParseFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, errors.New("number not readable")
	}
	return res, nil
}

// IsParenteses checks if the given string is an opening or closing parenthesis.
func IsParenteses(str string) bool {
	return str == OpenParen || str == CloseParen
}

// Operation performs a binary operation (e.g., addition, subtraction) between two operands
// located at the specified index in the equation members. It uses the provided callback function
// to perform the operation and returns the updated list of equation members.
func Operation(index int, members []string, membersResult []string, callback func(float64, float64) float64) ([]string, error) {
	a, err := ParseFloat(members[index-1])
	if err != nil {
		return nil, err
	}
	b, err := ParseFloat(members[index+1])
	if err != nil {
		return nil, err
	}
	membersResult = append(membersResult[:index-1], fmt.Sprintf("%f", callback(a, b)))

	return membersResult, nil
}
