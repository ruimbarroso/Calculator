// Copyright (c) 2025 Rui Barroso
// This code is licensed under the MIT License.
package parser

import (
	"calculator/src/lexer"
	"strconv"
)

// binding_power defines the precedence level for operators during parsing.
type binding_power int

const (
	default_bp binding_power = iota
	primary
	additive
	multiplicative
	exponential
	unary
)

// nud_lookup maps a lexer.TokenKind to its corresponding null denotation handler.
// A null denotation handler is responsible for parsing expressions that do not have a left-hand side.
type nud_lookup map[lexer.TokenKind]nud_handler

// led_lookup maps a lexer.TokenKind to its corresponding left denotation handler.
// A left denotation handler parses expressions that have a left-hand side.
type led_lookup map[lexer.TokenKind]led_handler

// bp_lookup maps a lexer.TokenKind to its binding power.
// This helps determine operator precedence during parsing.
type bp_lookup map[lexer.TokenKind]binding_power

// bp_lu is the lookup table for binding powers.
var bp_lu = bp_lookup{}

// nud_lu is the lookup table for null denotation (nud) handlers.
var nud_lu = nud_lookup{}

// led_lu is the lookup table for left denotation (led) handlers.
var led_lu = led_lookup{}

// nud_handler defines a function type for parsing expressions without a left-hand side.
// It takes a pointer to a parser and returns an expression.
type nud_handler func(p *parser) Expr

// led_handler defines a function type for parsing expressions with a left-hand side.
// It takes a pointer to a parser, the left-hand expression, and the current binding power,
// and returns a new expression.
type led_handler func(p *parser, left Expr, bp binding_power) Expr

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, bp binding_power, nud_fn nud_handler) {
	bp_lu[kind] = bp
	nud_lu[kind] = nud_fn
}

func createTokenLookups() {
	// Additive & Multiplicitave
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)
	led(lexer.ROOT, exponential, parse_binary_expr)
	led(lexer.HAT, exponential, parse_binary_expr)
	led(lexer.LOG, exponential, parse_binary_expr)

	// Literals & Symbols
	nud(lexer.NUMBER, primary, parse_primary_expr)

	// Unary Operators
	nud(lexer.DASH, additive, parse_unary_expr)

	// Grouping Expr
	nud(lexer.OPEN_PAREN, default_bp, parse_grouping_expr)
}

// parse_primary_expr parses a primary expression.
// A primary expression can be a literal, identifier, or any expression that doesn't require further operator precedence handling.
func parse_primary_expr(p *parser) Expr {
	number, _ := strconv.ParseFloat(p.advance().Value, 64)
	return NumberExpr{
		Value: number,
	}
}

// parse_unary_expr parses a unary expression.
// It handles expressions where a unary operator (such as '-') precedes an expression.
func parse_unary_expr(p *parser) Expr {
	token := p.advance()
	Member := parse_expr(p, unary)
	return UnaryExpr{
		Operator: token,
		Member:   Member,
	}
}

// parse_binary_expr parses a binary expression.
// It takes the left-hand side expression and the current binding power, and returns a new expression
// by handling the binary operator and the right-hand side expression.
func parse_binary_expr(p *parser, left Expr, bp binding_power) Expr {
	operatorToken := p.advance()
	right := parse_expr(p, bp_lu[operatorToken.Kind])

	return BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

// parse_grouping_expr parses a grouping expression.
// Grouping expressions, typically enclosed in parentheses, are used to explicitly specify the order of evaluation.
func parse_grouping_expr(p *parser) Expr {
	p.expect(lexer.OPEN_PAREN)
	expr := parse_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)

	return expr
}
