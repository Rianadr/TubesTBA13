package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenFor TokenType = iota
	TokenIdentifier
	TokenIllegal
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf(t.Value)
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) scanIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.ch {
	case 'f':
		if strings.HasPrefix(l.input[l.position:], "for") && !unicode.IsLetter(l.peekChar()) && !unicode.IsDigit(l.peekChar()) && l.peekChar() != '_' {
			token.Type = TokenFor
			token.Value = "for"
			l.readChar()
		} else {
			token.Type = TokenIdentifier
			token.Value = l.scanIdentifier()
		}
	case 0:
		token.Type = TokenEOF
	default:
		token.Type = TokenIllegal
		token.Value = string(l.ch)
	}

	l.readChar()
	return token
}

func main() {
	input := `for i := 0; i < 10; i++ {
		fmt.Println(i)
	}`

	lexer := NewLexer(input)

	for {
		token := lexer.NextToken()
		if token.Type == TokenEOF {
			break
		}
		fmt.Println(token)
	}
}
