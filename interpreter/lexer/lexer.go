// lexer/lexer.go
package lexer

import "monkey-go/interpreter/token"

type Lexer struct {
  input string
  position  int
  readPosition  int
  ch  byte
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}

func (l *Lexer) NextToken() token.Token {
  var t token.Token

  switch l.ch {
    case '=':
      t = newToken(token.ASSIGN, l.ch)
    case ';':
      t = newToken(token.SEMICOLON, l.ch)
    case '(':
      t = newToken(token.LPAREN, l.ch)
    case ')':
      t = newToken(token.RPAREN, l.ch)
    case ',':
      t = newToken(token.COMMA, l.ch)
    case '+':
      t = newToken(token.PLUS, l.ch)
    case '{':
      t = newToken(token.LBRACE, l.ch)
    case '}':
      t = newToken(token.RBRACE, l.ch)
    case '0':
      t.Literal = ""
      t.Type = token.EOF
  }

  l.readChar()
  return t
}

func newToken(t token.TokenType, ch byte) token.Token {
  return token.Token{Type: t, Literal: string(ch)}
}

func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }

  l.position = l.readPosition
  l.readPosition += 1
}

