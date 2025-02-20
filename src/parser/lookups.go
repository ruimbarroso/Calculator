package parser

import (
	"calculator/src/lexer"
	"strconv"
)

type binding_power int

const (
	default_bp binding_power = iota
	primary
	additive
	multiplicative
	exponential
	unary
)

type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}

type nud_handler func(p *parser) Expr
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

func parse_primary_expr(p *parser) Expr {
	number, _ := strconv.ParseFloat(p.advance().Value, 64)
	return NumberExpr{
		Value: number,
	}
}

func parse_unary_expr(p *parser) Expr {
	token := p.advance()
	Member := parse_expr(p, unary)
	return UnaryExpr{
		Operator: token,
		Member:   Member,
	}
}

func parse_binary_expr(p *parser, left Expr, bp binding_power) Expr {
	operatorToken := p.advance()
	right := parse_expr(p, bp_lu[operatorToken.Kind])

	return BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parse_grouping_expr(p *parser) Expr {
	p.expect(lexer.OPEN_PAREN)
	expr := parse_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)

	return expr
}
