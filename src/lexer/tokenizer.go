package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}
type regexHandler func(lex *lexer, regex *regexp.Regexp)

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

func (l *lexer) advanceN(i int) {
	l.pos += i
}

func (l *lexer) push(token *Token) {
	l.Tokens = append(l.Tokens, *token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) advance() {
	lex.pos += 1
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_end() bool {
	return lex.pos >= len(lex.source)
}

func Tokenize(source string) []Token {
	lex := createNewLexer(strings.ReplaceAll(source, " ", ""))

	for !lex.at_end() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("lexer error: unrecognized token near '%v'", lex.remainder()))
		}
	}

	lex.push(newToken(END, ";"))
	return lex.Tokens
}
func EquationIsValid(tokens []Token) bool {
	parentesesCount := 0
	operator := false
	for _, t := range tokens {
		if t.Kind == OPEN_PAREN {
			parentesesCount++
		} else if t.Kind == CLOSE_PAREN {
			parentesesCount--
		}
		if parentesesCount < 0 {
			return false
		}

		if !IsOneOf(t.Kind, []TokenKind{OPEN_PAREN, CLOSE_PAREN, NUMBER}) {
			if !operator {
				return false
			}
			operator = false
		} else {
			operator = true
		}
	}

	return parentesesCount == 0
}

func createNewLexer(source string) *lexer {
	return &lexer{
		patterns: []regexPattern{
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
			{regexp.MustCompile(`\^`), defaultHandler(HAT, "^")},
			{regexp.MustCompile(`r`), defaultHandler(ROOT, "r")},
			{regexp.MustCompile(`l`), defaultHandler(LOG, "l")},
			{regexp.MustCompile(`l`), defaultHandler(SLASH, "-")},
		},
		Tokens: make([]Token, 0),
		source: source,
		pos:    0,
	}
}
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, _ *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(newToken(kind, value))
	}
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(newToken(NUMBER, match))
	lex.advanceN(len(match))
}
