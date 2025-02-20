package lexer

import "fmt"

type TokenKind int

const (
	END TokenKind = iota
	NUMBER

	// Parenteses
	OPEN_PAREN
	CLOSE_PAREN

	//Maths
	PLUS
	DASH
	SLASH
	STAR
	PERCENT
	ROOT
	HAT
	LOG
)

func TokenKindString(kind TokenKind) string {
	switch kind {
	case END:
		return "END"
	case NUMBER:
		return "NUMBER"
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case PLUS:
		return "PLUS"
	case DASH:
		return "DASH"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case PERCENT:
		return "PERCENT"
	case ROOT:
		return "ROOT"
	case HAT:
		return "HAT"
	case LOG:
		return "LOG"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", kind)
	}
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (t Token) ToString() string {
	return fmt.Sprintf("%s \"%s\"", t.KindString(), t.Value)

}
func (t Token) KindString() string {
	return TokenKindString(t.Kind)

}

func newToken(kind TokenKind, value string) *Token {
	return &Token{
		Kind: kind, Value: value,
	}
}

func IsOneOf(kind TokenKind, list []TokenKind) bool {
	for _, item := range list {
		if item == kind {
			return true
		}
	}

	return false
}
