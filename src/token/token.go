package token

const (
	// single char tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	COLON
	PLUS
	MINUS
	STAR
	SLASH
	QUESTION

	// one or two-char tokens
	DOUBLE_EQUAL
	BANG
	BANG_EQUAL
	EQUAL
	GREATER
	LESS
	GREATER_EQUAL
	LESS_EQUAL
	INCREMENT // ++
	DECREMENT // --

	// literals
	IDENTIFIER
	STRING
	NUMBER
	FLOAT

	// keywords
	IF
	FUNC
	VAR
	RETURN
	TRUE
	FALSE
	NIL

	NEWLINE
	EOF
)

type TokenType int

func (tt TokenType) String() string {
	switch tt {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case COLON:
		return "COLON"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case QUESTION:
		return "QUESTION"
	case DOUBLE_EQUAL:
		return "DOUBLE_EQUAL"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL:
		return "EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case INCREMENT:
		return "INCREMENT"
	case DECREMENT:
		return "DECREMENT"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case FLOAT:
		return "FLOAT"
	case IDENTIFIER:
		return "IDENTIFIER"
	case IF:
		return "IF"
	case FUNC:
		return "FUNC"
	case VAR:
		return "VAR"
	case RETURN:
		return "RETURN"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case NEWLINE:
		return "NEWLINE"
	case EOF:
		return "EOF"
	default:
		return ""
	}
}

type Token struct {
	Type     TokenType
	Text     string
	Literal  interface{}
	Position int
	Line     int
}

func Keywords() map[string]TokenType {
	return map[string]TokenType{
		"if":     IF,
		"func":   FUNC,
		"var":    VAR,
		"return": RETURN,
		"true":   TRUE,
		"false":  FALSE,
		"nil":    NIL,
	}
}

func IsKeyword(tt TokenType) bool {
	for _, tokenType := range Keywords() {
		if tt == tokenType {
			return true
		}
	}
	return false
}
