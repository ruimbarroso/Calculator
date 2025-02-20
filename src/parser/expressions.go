package parser

import (
	"calculator/src/lexer"
	"fmt"
	"math"
)

type Expr interface {
	ToString() string
	Eval() float64
}

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) ToString() string {
	return fmt.Sprintf("%g", n.Value)
}
func (n NumberExpr) Eval() float64 {
	return n.Value
}

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
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

type UnaryExpr struct {
	Operator lexer.Token
	Member   Expr
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
