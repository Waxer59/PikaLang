package parser

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/lexer"
	"pika/pkg/lexer/token_type"
)

type Parser struct {
	tokens []token_type.Token
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ProduceAST(input string) (*ast.Program, error) {
	p.tokens = lexer.Tokenize(input)
	program := ast.Program{
		Kind: ast_types.Program,
		Body: []ast.Stmt{},
	}

	for p.notEOF() {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		program.Body = append(program.Body, stmt)
	}

	return &program, nil
}
