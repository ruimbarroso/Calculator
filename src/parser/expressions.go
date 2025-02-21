// Copyright (c) 2025 Rui Barroso
// This code is licensed under the MIT License.
package parser

import (
	"calculator/src/lexer"
	"fmt"
	"math"
)

// Expr represents an equation expression that can be converted to a string and evaluated.
type Expr interface {
	// ToString returns a string representation of the expression.
	ToString() string
	// Eval computes and returns the value of the expression.
	Eval() float64
}

// NumberExpr represents a numeric expression.
type NumberExpr struct {
	// Value holds the numeric value of the expression.
	Value float64
}

func (n NumberExpr) ToString() string {
	return fmt.Sprintf("%g", n.Value)
}
func (n NumberExpr) Eval() float64 {
	return n.Value
}

// BinaryExpr represents an expression with a binary operator.
// It contains a left-hand expression, an operator token, and a right-hand expression.
type BinaryExpr struct {
	// Left is the expression on the left side of the operator.
	Left Expr
	// Operator is the token representing the binary operator.
	Operator lexer.Token
	// Right is the expression on the right side of the operator.
	Right Expr
}

func (n BinaryExpr) ToString() string {
	return fmt.Sprintf("(%s %s %s)", n.Left.ToString(), n.Operator.Value, n.Right.ToString())
}
func (n BinaryExpr) Eval() float64 {
	a := n.Left.Eval()
	b := n.Right.Eval()

	switch n.Operator.Kind {
	case lexer.PLUS:
		return a + b
	case lexer.DASH:
		return a - b
	case lexer.STAR:
		return round(a*b, 10)
	case lexer.SLASH:
		return round(a/b, 10)
	case lexer.PERCENT:
		return math.Remainder(a, b)
	case lexer.ROOT:
		return round(math.Pow(a, 1/b), 10)
	case lexer.HAT:
		return round(math.Pow(a, b), 10)
	case lexer.LOG:
		return round(math.Log(a)/math.Log(b), 10)
	default:
		panic(fmt.Sprintf("Operator %s not recognized", n.Operator.KindString()))
	}
}
func round(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

// UnaryExpr represents an expression with a unary operator.
// It contains an operator token and a member expression upon which the operator is applied.
type UnaryExpr struct {
	// Operator is the token representing the unary operator.
	Operator lexer.Token
	// Member is the expression on which the operator is applied.
	Member Expr
}

func (n UnaryExpr) ToString() string {
	return fmt.Sprintf("(%s%s)", n.Operator.Value, n.Member.ToString())
}
func (n UnaryExpr) Eval() float64 {
	switch n.Operator.Kind {
	case lexer.DASH:
		return -1 * n.Member.Eval()
	default:
		panic(fmt.Sprintf("Operator %s not recognized", n.Operator.KindString()))
	}
}
