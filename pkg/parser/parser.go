package parser

import (
	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/ast/ast_types"
	"github.com/Waxer59/PikaLang/pkg/lexer"
	"github.com/Waxer59/PikaLang/pkg/lexer/token_type"
)

type Parser struct {
	tokens []token_type.Token
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ProduceAST(input string) (*ast.Program, error) {
	var err error
	p.tokens, err = lexer.Tokenize(input)

	if err != nil {
		return nil, err
	}

	program := ast.Program{
		Kind: ast_types.Program,
		Body: []ast.Stmt{},
	}

	for p.notEOF() {
		stmt, err := p.ParseStmt()
		if err != nil {
			return nil, err
		}
		program.Body = append(program.Body, stmt)
	}

	return &program, nil
}

func (p *Parser) SetTokens(tokens []token_type.Token) {
	p.tokens = tokens
}
