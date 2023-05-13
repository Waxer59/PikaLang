package parser

import "pika/pkg/lexer/lexerTypes"

func (p *Parser) at() lexerTypes.Token {
	return p.tokens[0]
}

func (p *Parser) subtract() lexerTypes.Token {
	prev := p.at()
	p.tokens = p.tokens[1:]
	return prev
}

func (p *Parser) expect(typeExpected lexerTypes.TokenType, errMsg string) lexerTypes.Token {
	prev := p.subtract()
	if (prev == lexerTypes.Token{} || prev.Type != typeExpected) {
		panic("Parser Error:\n" + errMsg)
	}

	return prev
}

func (p *Parser) notEOF() bool {
	return p.at().Type != lexerTypes.EOF
}
