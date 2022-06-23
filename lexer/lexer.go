package lexer

import (
	"github.com/yonh/go-monkey-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // 当前字符位于字符串所在索引
	readPosition int  // 当前读取字符的下一个字符的位置
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// reachChar 的目的是读取下一个字符，并更新 position, readPosition，
// 如果读取到文件末尾，则将 l.ch = 0,这是 NUL 字符的ASCII编码，用来表示“尚未读取到任何内容”或“文件结尾”
// 注意：此方法近支持ASCII字符，不能完全支持Unicode字符，这么做是为了简单
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 循环读取字符，直至遇到非字母或下划线字符，返回读取到的字母或下划线字符串
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 循环读取字符，直至遇到非数字字符，返回读取到的数字字符串
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 判断字符是否是字母或下划线
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 此方法首先根据l.ch返回相应的Token类型，然后调用l.readChar()将读取索引指向下一个字符
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 跳过空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
