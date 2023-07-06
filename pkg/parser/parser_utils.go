package parser

import (
	"errors"
	"os"
	compilerErrors "pika/internal/errors"
	"pika/pkg/ast"
	"pika/pkg/lexer/token_type"

	"github.com/fatih/color"
)

func (p *Parser) at() token_type.Token {
	return p.tokens[0]
}

func (p *Parser) atNext() token_type.Token {
	if len(p.tokens) > 1 {
		return p.tokens[1]
	}
	return token_type.Token{}
}

func (p *Parser) subtract(params ...int) token_type.Token {
	n := 1
	if len(params) > 0 {
		n = params[0]
	}

	prev := p.at()
	p.tokens = p.tokens[n:]
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

func (p *Parser) parseArgs(argType token_type.TokenType) ([]ast.Expr, error) {

	switch argType {
	case token_type.Fn, token_type.Case:
		return p.parseMultipleArgs(argType)
	case token_type.If, token_type.Switch:
		return p.parseSingleArg()
	}

	return nil, errors.New(compilerErrors.ErrSyntaxStatementNotFound)
}

func (p *Parser) parseSingleArg() ([]ast.Expr, error) {

	if p.at().Type == token_type.LeftParen { // Optional parens
		p.subtract() // Remove the opening paren
	}

	condition, err := p.parseExpr()

	if err != nil {
		return nil, err
	}

	if condition == nil {
		return nil, errors.New(compilerErrors.ErrConditionCannotBeEmpty)
	}

	if p.at().Type == token_type.RightParen { // Optional parens
		p.subtract() // Remove the closing paren
	}

	return []ast.Expr{condition}, nil
}

func (p *Parser) parseMultipleArgs(argType token_type.TokenType) ([]ast.Expr, error) {

	if argType == token_type.Fn {
		p.expect(token_type.LeftParen, compilerErrors.ErrSyntaxExpectedLeftParen)
	}

	if argType == token_type.Case && p.at().Type == token_type.Colon {
		return nil, errors.New(compilerErrors.ErrSyntaxCaseCannotBeEmpty)
	}

	args := []ast.Expr{}

	if p.at().Type != token_type.RightParen {
		var err error
		args, err = p.parseArgsList()

		if err != nil {
			return nil, err
		}
	}

	if argType == token_type.Fn {
		p.expect(token_type.RightParen, compilerErrors.ErrSyntaxExpectedRightParen)
	}

	return args, nil
}

func (p *Parser) parseArgsList() ([]ast.Expr, error) {
	assigmentExpr, err := p.parseAssigmentExpr()
	if err != nil {
		return nil, err
	}
	args := []ast.Expr{assigmentExpr}

	for p.at().Type == token_type.Comma && p.subtract().Type == token_type.Comma {
		assigmentExpr, err := p.parseAssigmentExpr()
		if err != nil {
			return nil, err
		}
		args = append(args, assigmentExpr)
	}

	return args, nil
}

func (p *Parser) parseBlockBodyStmt() ([]ast.Stmt, error) {
	var body []ast.Stmt

	p.expect(token_type.LeftBrace, compilerErrors.ErrSyntaxExpectedLeftBrace)

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.EOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}

	p.expect(token_type.RightBrace, compilerErrors.ErrSyntaxExpectedRightBrace)

	return body, nil
}

func (p *Parser) parseSwitchBodyStmt() ([]ast.Stmt, error) {
	var body []ast.Stmt

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.Case && p.at().Type != token_type.Default && p.at().Type != token_type.EOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}
	return body, nil
}
