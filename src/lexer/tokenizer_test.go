package lexer_test

import (
	"calculator/src/lexer"
	"fmt"
	"testing"
)

type EquationResult struct {
	eq           string
	expextedSize int
}

var equations = []EquationResult{
	{"-2", 3},
	{"-2+2", 5},
	{"2-2", 4},
	{"2+2", 4},
	{"5*2", 4},
	{"3*2+3", 6},
	{"(2/6)*5", 8},
	{"4r7-2^5", 8},
	{"45.2+81", 4},
	{"42^(3+2)", 8},
	{"45.2++81", -1},
	{"45.2+-81", -1},
	{"+45.2+81", -1},
	{"45.2+81+", -1},
	{"(45.2+81))", -1},
	{")45.2+81(", -1},
}

// Tokenize(equations[i]) []Token
// fmt.Println(tokens[i].ToString())
func TestTokenizeEquations(t *testing.T) {
	for i, res := range equations {
		tokens := lexer.Tokenize(res.eq)
		fmt.Printf("Equation %d: %s\n", i+1, res.eq)
		for _, token := range tokens {
			fmt.Printf("%s\n", token.ToString())
		}

		if res.expextedSize == -1 {
			if lexer.EquationIsValid(tokens) {
				t.Errorf("Expected Invalid Equation")
			}
		} else if res.expextedSize != len(tokens) {
			t.Errorf("Expected %d Tokens but had %d", res.expextedSize, len(tokens))
		}
		fmt.Println()
	}
}
