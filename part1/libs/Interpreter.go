package calc1

import (
	"fmt"
	"strconv"
	"unicode"
)

// Interpreter 为解析器
type Interpreter struct {
	Text         string
	Pos          int
	currentToken *Token
}

func (interpreter *Interpreter) getNextToken() *Token {
	text := interpreter.Text
	if interpreter.Pos > len(text)-1 { // 解析到句末
		return &Token{EOF, ""}
	}
	currentChar := text[interpreter.Pos]

	if unicode.IsDigit(rune(currentChar)) {
		token := Token{INTEGER, text[interpreter.Pos : interpreter.Pos+1]}
		interpreter.Pos++
		return &token
	}

	if currentChar == '+' {
		token := Token{PLUS, text[interpreter.Pos : interpreter.Pos+1]}
		interpreter.Pos++
		return &token
	}

	return nil
}

func (interpreter *Interpreter) eat(tokenType TokenType) {
	if interpreter.currentToken.GetType() == tokenType {
		interpreter.currentToken = interpreter.getNextToken()
	}
}

// Expr 解析表达式
func (interpreter *Interpreter) Expr() (int, error) {
	interpreter.currentToken = interpreter.getNextToken()

	if interpreter.currentToken.GetType() != INTEGER {
		return 0, fmt.Errorf("Expect INTEGER, actual: %v", interpreter.currentToken.GetType())
	}
	num, err := strconv.ParseInt(interpreter.currentToken.GetLexeme(), 10, 0)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse %v to INTEGER", interpreter.currentToken.GetLexeme())
	}
	left := NumToken{*(interpreter.currentToken), int(num)}

	interpreter.currentToken = interpreter.getNextToken()
	if interpreter.currentToken.GetType() != PLUS {
		return 0, fmt.Errorf("Expect PLUS, actual: %v", interpreter.currentToken.GetType())
	}
	// op := PlusToken{*interpreter.currentToken}

	interpreter.currentToken = interpreter.getNextToken()
	if interpreter.currentToken.GetType() != INTEGER {
		return 0, fmt.Errorf("Expect INTEGER, actual %v", interpreter.currentToken.GetType())
	}
	num, err = strconv.ParseInt(interpreter.currentToken.GetLexeme(), 10, 0)
	right := NumToken{*(interpreter.currentToken), int(num)}

	interpreter.currentToken = interpreter.getNextToken()
	if interpreter.currentToken.GetType() != EOF {
		return 0, fmt.Errorf("Expect EOF, actual %v", interpreter.currentToken.GetType())
	}

	result := left.GetValue() + right.GetValue()
	return result, nil
}
