package filterexpr

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	input := "((string,1,contain,banana)or(time,2,>,2025-03-20 00:00:00))and(int,3,==,1000)"

	expected := []Token{
		{Type: TokenLParen, Value: "("},
		{Type: TokenLParen, Value: "("},
		{Type: TokenValue, Value: "string"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "1"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "contain"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "banana"},
		{Type: TokenRParen, Value: ")"},
		{Type: TokenOp, Value: "or"},
		{Type: TokenLParen, Value: "("},
		{Type: TokenValue, Value: "time"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "2"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: ">"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "2025-03-20 00:00:00"},
		{Type: TokenRParen, Value: ")"},
		{Type: TokenRParen, Value: ")"},
		{Type: TokenOp, Value: "and"},
		{Type: TokenLParen, Value: "("},
		{Type: TokenValue, Value: "int"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "3"},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "=="},
		{Type: TokenComma, Value: ","},
		{Type: TokenValue, Value: "1000"},
		{Type: TokenRParen, Value: ")"},
	}

	tokens, err := tokenize(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("tokenize result mismatch\nGot: %#v\nWant: %#v", tokens, expected)
	}
}
