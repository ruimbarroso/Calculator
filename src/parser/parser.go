package parser

import (
	"calculator/src/lexer"
	"fmt"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func (p *parser) current() lexer.Token {
	return p.tokens[p.pos]
}
func (p *parser) advance() lexer.Token {
	next := p.tokens[p.pos]
	p.pos++
	return next
}
func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	token := p.current()

	if token.Kind != expectedKind {

		err := fmt.Sprintf("Expected %s but recieved %s instead\n", lexer.TokenKindString(expectedKind), token.KindString())

		panic(err)
	}

	return p.advance()
}
func createParser(tokens []lexer.Token) *parser {
	/* createTokenLookups()
	createTypeTokenLookups() */
	createTokenLookups()
	p := &parser{
		tokens: tokens,
		pos:    0,
	}

	return p
}

func Parse(source string) Expr {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\n\nerror parsing equation\n\n")
			// return errors.New("error parsing equation")
		}
	}()
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)

	expr := parse_expr(p, default_bp)
	return expr
}

func parse_expr(p *parser, bp binding_power) Expr {
	tokenKind := p.current().Kind
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for bp_lu[p.current().Kind] > bp {
		tokenKind = p.current().Kind
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}
