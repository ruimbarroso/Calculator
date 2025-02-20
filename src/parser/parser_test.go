package parser_test

import (
	"calculator/src/parser"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/sanity-io/litter"
)

type EquationResult struct {
	eq             string
	expextedResult float64
}

var equations = []EquationResult{
	// Basic arithmetic
	{"1 + 1", 2},
	{"2 - 1", 1},
	{"2 * 3", 6},
	{"6 / 2", 3},
	{"2 + 3 * 4", 14},
	{"(2 + 3) * 4", 20},
	{"2 * (3 + 4)", 14},
	{"2 * 3 + 4", 10},

	// Unary negation
	{"-1", -1},
	{"-1 + 2", 1},
	{"2 + (-1)", 1},
	{"-2 * 3", -6},
	{"2 * (-3)", -6},
	{"-2 * (-3)", 6},
	{"-(2 + 3)", -5},

	// Division and floating-point results
	{"1 / 2", 0.5},
	{"1 / 4", 0.25},
	{"2 / 4", 0.5},
	{"3 / 2", 1.5},

	// Exponentiation
	{"2 ^ 3", 8},
	{"2 ^ 0", 1},
	{"2 ^ -1", 0.5},
	{"-(2 ^ 2)", -4},
	{"(-2) ^ 2", 4},
	{"9 r2", 3},
	{"8 r3 ", 2},
	{"100 l 10", 2},

	// Complex expressions
	{"2 + 3 * 4 - 5 / 2", 11.5},
	{"(2 + 3) * (4 - 5) / 2", -2.5},
	{"2 * (3 + 4) - 5 / (2 + 3)", 13},
	{"-2 * (3 + 4) - 5 / (2 + 3)", -15},

	// Edge cases
	{"0 + 0", 0},
	{"0 * 5", 0},
	{"5 * 0", 0},
	{"0 / 5", 0},
	{"5 / 0", math.Inf(1)}, // Division by zero (should handle gracefully)
	{"-0", 0},              // Negative zero
}

func TestEquationParser(t *testing.T) {
	for i, eq := range equations {
		fmt.Printf("Equation %d: %s\n", i+1, eq.eq)
		start := time.Now()
		ast := parser.Parse(eq.eq)
		duration := time.Since(start)

		fmt.Printf("Result Equation: %s\n", ast.ToString())
		fmt.Printf("Result: %g\n", ast.Eval())
		fmt.Printf("Duration: %v\n", duration)

		fmt.Println()

		if eq.expextedResult != ast.Eval() {
			t.Errorf("In Equation %s\n Expected result is %g but the result was %g", eq.eq, eq.expextedResult, ast.Eval())
			litter.Dump(ast)
		}
	}
}
