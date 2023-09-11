package lexer_test

import (
	"crest/lexer"
	"crest/token"
	"testing"
)

type TokenCase struct {
	InputString string
	Types       []token.TokenType
}

func TestLexer(t *testing.T) {
	cases := map[string]TokenCase{
		"basic string variable": {
			InputString: "var myName \"Tim\"",
			Types: []token.TokenType{
				token.VAR,
				token.IDENTIFIER,
				token.STRING,
				token.EOF,
			},
		},
		"basic list variable": {
			InputString: "var names {\"Tim\", \"Jane\"}",
			Types: []token.TokenType{
				token.VAR,
				token.IDENTIFIER,
				token.LEFT_BRACE,
				token.STRING,
				token.COMMA,
				token.STRING,
				token.RIGHT_BRACE,
				token.EOF,
			},
		},
		"function declaration": {
			InputString: "func sayHello (name) { print \"hello \" + name }",
			Types: []token.TokenType{
				token.FUNC,
				token.IDENTIFIER,
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.RIGHT_PAREN,
				token.LEFT_BRACE,
				token.IDENTIFIER,
				token.STRING,
				token.PLUS,
				token.IDENTIFIER,
				token.RIGHT_BRACE,
				token.EOF,
			},
		},
		"function call": {
			InputString: "sayHello \"Tim\"",
			Types: []token.TokenType{
				token.IDENTIFIER,
				token.STRING,
				token.EOF,
			},
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			if !slicesMatch(l.TokenTypes(), testcase.Types) {
				t.Fatal("types do not match", testcase.Types, l.TokenTypes())
			}
		})
	}
}

func slicesMatch(a []token.TokenType, b []token.TokenType) bool {
	if len(a) != len(b) {
		return false
	}

	for index, aType := range a {
		bType := b[index]
		if bType != aType {
			return false
		}
	}
	return true
}
