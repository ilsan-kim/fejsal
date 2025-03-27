package filterexpr

import "strings"

type TokenType int

const (
	TokenLParen TokenType = iota // (
	TokenRParen                  // )
	TokenOp                      // and, or, &&, ||
	TokenComma                   // ,
	TokenValue                   // string, int, time - 1 (column of csv), email (key of json) - contain, equal - banana
)

type Token struct {
	Type  TokenType
	Value string
}

func tokenize(input string) ([]Token, error) {
	var tokens []Token
	var buf strings.Builder

	isKeyword := func(s string) bool {
		return s == "and" || s == "or" || s == "&&" || s == "||"
	}

	flushBuf := func() {
		if buf.Len() == 0 {
			return
		}
		word := buf.String()
		if isKeyword(word) {
			tokens = append(tokens, Token{Type: TokenOp, Value: word})
		} else {
			tokens = append(tokens, Token{Type: TokenValue, Value: word})
		}
		buf.Reset()
	}

	for i := 0; i < len(input); i++ {
		c := input[i]
		switch c {
		case '(':
			flushBuf()
			tokens = append(tokens, Token{Type: TokenLParen, Value: "("})
		case ')':
			flushBuf()
			tokens = append(tokens, Token{Type: TokenRParen, Value: ")"})
		case ',':
			flushBuf()
			tokens = append(tokens, Token{Type: TokenComma, Value: ","})
		default:
			buf.WriteByte(c)
		}
	}
	flushBuf()

	return tokens, nil
}
