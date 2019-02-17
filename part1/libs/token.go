package calc1

import (
	"fmt"
)

// TokenType 为Token类型
type TokenType int

// 定义Token类型枚举值
const (
	INTEGER TokenType = iota
	PLUS
	EOF
)

// Token 接口定义
type Token struct {
	tokenType TokenType
	lexeme    string
}

// GetType 返回Token类型
func (t *Token) GetType() TokenType {
	return t.tokenType
}

// GetLexeme 返回lexeme
func (t *Token) GetLexeme() string {
	return t.lexeme
}

// NumToken 为数字型Token
type NumToken struct {
	token      Token
	tokenValue int
}

// GetValue 返回NumToken的数值
func (numToken *NumToken) GetValue() int {
	return numToken.tokenValue
}

func (numToken *NumToken) String() string {
	return fmt.Sprintf("Token(INTEGER, %d)", numToken.tokenValue)
}

// PlusToken 表示加号Token
type PlusToken struct {
	Token
}

func (plusToken *PlusToken) String() string {
	return "Token(PLUS, +)"
}

// EOFToken 表示终止符
type EOFToken struct {
	Token
}

func (eofToken *EOFToken) String() string {
	return "Token(EOF, EOF)"
}
