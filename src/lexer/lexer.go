package lexer

import (
	"crest/token"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func New(input string) Lexer {
	lexer := Lexer{
		Input:   input,
		Line:    1,
		Start:   0,
		Current: 0,
	}
	err := lexer.ReadInput()
	if err != nil {
		panic(err)
	}
	return lexer
}

type Lexer struct {
	Input   string
	Tokens  []token.Token
	Start   int
	Current int
	Line    int
}

func (l *Lexer) ReadInput() error {
	for !l.isAtEnd() {
		l.Start = l.Current
		err := l.readChar()
		if err != nil {
			return err
		}
	}
	l.Start++
	l.addToken(token.EOF, "", "")
	return nil
}

func (l *Lexer) readChar() error {
	char := l.nextChar()
	switch char {
	case "(":
		l.addToken(token.LEFT_PAREN, char, char)
	case ")":
		l.addToken(token.RIGHT_PAREN, char, char)
	case "{":
		l.addToken(token.LEFT_BRACE, char, char)
	case "}":
		l.addToken(token.RIGHT_BRACE, char, char)
	case ",":
		l.addToken(token.COMMA, char, char)
	case ".":
		l.addToken(token.DOT, char, char)
	case ":":
		l.addToken(token.COLON, char, char)
	case "+":
		if l.matchNext("+") {
			l.addToken(token.INCREMENT, "++", "++")
		} else {
			l.addToken(token.PLUS, char, char)
		}
	case "-":
		if l.matchNext("-") {
			l.addToken(token.DECREMENT, "--", "--")
		} else {
			l.addToken(token.MINUS, char, char)
		}
	case "*":
		l.addToken(token.STAR, char, char)
	case "/":
		l.addToken(token.SLASH, char, char)
	case "?":
		l.addToken(token.QUESTION, char, char)
	case "!":
		if l.matchNext("=") {
			l.addToken(token.BANG_EQUAL, "!=", "!=")
		} else {
			l.addToken(token.BANG, char, char)
		}
	case "=":
		if l.matchNext("=") {
			l.addToken(token.DOUBLE_EQUAL, "==", "==")
		} else {
			l.addToken(token.EQUAL, char, char)
		}
	case "<":
		if l.matchNext("=") {
			l.addToken(token.LESS_EQUAL, "<=", "<=")
		} else {
			l.addToken(token.LESS, char, char)
		}
	case ">":
		if l.matchNext("=") {
			l.addToken(token.GREATER_EQUAL, ">=", ">=")
		} else {
			l.addToken(token.GREATER, char, char)
		}
	case "\"":
		l.matchString()
	case "\n":
		l.Line++
	case " ", "\r", "\t":
		break
	default:
		if isDigit(char) {
			l.matchNumber()
		} else if isLetter(char) {
			l.matchIdentifier()
		} else {
			return fmt.Errorf("unsupported type: %s", char)
		}
	}
	return nil
}

func (l *Lexer) nextChar() string {
	char := string(l.Input[l.Current])
	l.Current++
	return char
}

func (l *Lexer) addToken(tokenType token.TokenType, text string, literal interface{}) {
	l.Tokens = append(l.Tokens, token.Token{
		Type:     tokenType,
		Text:     text,
		Literal:  literal,
		Position: l.Start,
		Line:     l.Line,
	})
}

func (l *Lexer) isAtEnd() bool {
	return l.Current >= len(l.Input)
}

func (l *Lexer) peek() string {
	if l.isAtEnd() {
		return ""
	}

	return string(l.Input[l.Current])
}

func (l *Lexer) peekNext() string {
	if l.Current+1 >= len(l.Input) {
		return ""
	}
	return string(l.Input[l.Current+1])
}

func (l *Lexer) matchNumber() {
	for isDigit(l.peek()) {
		l.nextChar()
	}

	if l.peek() == "." && isDigit(l.peekNext()) {
		l.nextChar()

		for isDigit(l.peek()) {
			l.nextChar()
		}
	}

	text := l.Input[l.Start:l.Current]

	var val interface{}
	if strings.Contains(text, ".") {
		val, _ = strconv.ParseFloat(text, 64)
		l.addToken(token.FLOAT, text, val)
	} else {
		intVal, _ := strconv.ParseInt(text, 10, 0)
		val = int(intVal)
		l.addToken(token.NUMBER, text, val)
	}
}

func (l *Lexer) matchString() {
	for l.peek() != "\"" && !l.isAtEnd() {
		l.nextChar()
	}

	l.nextChar()
	text := l.Input[l.Start+1 : l.Current-1]
	l.addToken(token.STRING, text, text)
}

func (l *Lexer) matchIdentifier() {
	for isAlphaNumeric(l.peek()) && !l.isAtEnd() {
		l.nextChar()
	}

	text := l.Input[l.Start:l.Current]
	if tokenType, ok := token.Keywords()[text]; ok {
		l.addToken(tokenType, text, text)
	} else {
		l.addToken(token.IDENTIFIER, text, text)
	}
}

func (l *Lexer) matchNext(expected string) bool {
	if string(l.Input[l.Current]) != expected {
		return false
	}
	l.nextChar()
	return true
}

func isDigit(ch string) bool {
	_, err := strconv.Atoi(ch)
	return err == nil
}

func isLetter(ch string) bool {
	for _, r := range ch {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isAlphaNumeric(ch string) bool {
	return isDigit(ch) || isLetter(ch)
}
