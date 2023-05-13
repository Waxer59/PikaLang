package parser

import (
	"fmt"
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/lexer"
	"pika/pkg/lexer/lexerTypes"
)

type Parser struct {
	tokens []lexerTypes.Token
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ProduceAST(input string) ast.Program {
	p.tokens = lexer.Tokenize(input)
	fmt.Println(p.tokens)

	program := ast.Program{
		Kind: astTypes.Program,
		Body: []ast.Stmt{},
	}

	for p.notEOF() {
		program.Body = append(program.Body, p.parseStmt())
	}

	return program
}
