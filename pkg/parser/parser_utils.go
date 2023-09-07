package parser

import (
	"errors"
	"os"

	compilerErrors "github.com/Waxer59/PikaLang/internal/errors"
	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/lexer/token_type"

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

func (p *Parser) parseConditionalArg() (ast.Expr, error) {

	if p.at().Type == token_type.LeftParen { // Optional parens
		p.subtract() // Remove the opening paren
	}

	if p.at().Type == token_type.RightParen || p.at().Type == token_type.LeftBrace {
		return nil, errors.New(compilerErrors.ErrSyntaxConditionCantBeEmpty)
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

	return condition, nil
}

func (p *Parser) parseFunctionArgs() ([]ast.Identifier, error) {

	p.expect(token_type.LeftParen, compilerErrors.ErrSyntaxExpectedLeftParen)

	args := []ast.Identifier{}

	if p.at().Type == token_type.RightParen {
		p.subtract() // Remove the closing paren
		return args, nil
	}

	for {
		assigmentExpr, err := p.parseExpr()
		identifier, ok := assigmentExpr.(ast.Identifier)

		if err != nil || !ok {
			return nil, err
		}
		args = append(args, identifier)

		if p.at().Type != token_type.Comma {
			break
		}

		p.subtract() // consume ','
	}

	p.expect(token_type.RightParen, compilerErrors.ErrSyntaxExpectedRightParen)

	return args, nil
}

func (p *Parser) parseCallExprArgs() ([]ast.Expr, error) {
	p.expect(token_type.LeftParen, compilerErrors.ErrSyntaxExpectedLeftParen)

	if p.at().Type == token_type.RightParen {
		p.subtract() // Remove the closing paren
		return []ast.Expr{}, nil
	}

	argsList, err := p.parseArgsList()

	args, ok := argsList.([]ast.Expr)

	if err != nil || !ok {
		return nil, err
	}

	p.expect(token_type.RightParen, compilerErrors.ErrSyntaxExpectedRightParen)

	return args, nil
}

func (p *Parser) parseSwitchCaseArgs() ([]ast.Expr, error) {
	if p.at().Type == token_type.Colon {
		return nil, errors.New(compilerErrors.ErrSyntaxCaseCannotBeEmpty)
	}

	argsList, err := p.parseArgsList()

	args, ok := argsList.([]ast.Expr)

	if err != nil || !ok {
		return nil, err
	}

	p.expect(token_type.Colon, compilerErrors.ErrSyntaxExpectedColon)

	return args, nil
}

func (p *Parser) parseArgsList() (any, error) {
	args := []ast.Expr{}

	for {
		assigmentExpr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		args = append(args, assigmentExpr)

		if p.at().Type != token_type.Comma {
			break
		}

		p.subtract() // consume ','
	}

	return args, nil
}

func (p *Parser) parseBlockBodyStmt() ([]ast.Stmt, error) {
	var body []ast.Stmt

	p.expect(token_type.LeftBrace, compilerErrors.ErrSyntaxExpectedLeftBrace)

	for p.at().Type != token_type.RightBrace && p.notEOF() {
		stmt, err := p.ParseStmt()
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

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.Case && p.at().Type != token_type.Default && p.notEOF() {
		stmt, err := p.ParseStmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}
	return body, nil
}
