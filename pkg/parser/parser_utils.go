package parser

import (
	"os"
	"pika/pkg/lexer/token_type"

	"github.com/fatih/color"
)

func (p *Parser) at() token_type.Token {
	return p.tokens[0]
}

func (p *Parser) next() token_type.Token {
	if len(p.tokens) > 1 {
		return p.tokens[1]
	}
	return token_type.Token{}
}

func (p *Parser) subtract() token_type.Token {
	prev := p.at()
	p.tokens = p.tokens[1:]
	return prev
}

func (p *Parser) expect(typeExpected token_type.TokenType, errMsg string) token_type.Token {
	prev := p.subtract()
	if (prev == token_type.Token{} || prev.Type != typeExpected) {
		color.Red(errMsg)
		os.Exit(0)
	}

	return prev
}

func (p *Parser) notEOF() bool {
	return p.at().Type != token_type.EOF
}
