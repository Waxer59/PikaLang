package parser

import (
	"errors"
	compilerErrors "github.com/Waxer59/PikaLang/internal/errors"
	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/lexer/token_type"
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

func (p *Parser) expect(typeExpected token_type.TokenType, errMsg string) (token_type.Token, error) {
	prev := p.subtract()

	if (prev == token_type.Token{} || prev.Type != typeExpected) {
		return prev, errors.New(errMsg)
	}

	return prev, nil
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
	_, err := p.expect(token_type.LeftParen, compilerErrors.ErrSyntaxExpectedLeftParen)
	if err != nil {
		return nil, err
	}

	var args []ast.Identifier

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

	_, err = p.expect(token_type.RightParen, compilerErrors.ErrSyntaxExpectedRightParen)
	if err != nil {
		return nil, err
	}

	return args, nil
}

func (p *Parser) parseCallExprArgs() ([]ast.Expr, error) {
	_, err := p.expect(token_type.LeftParen, compilerErrors.ErrSyntaxExpectedLeftParen)
	if err != nil {
		return nil, err
	}

	if p.at().Type == token_type.RightParen {
		p.subtract() // Remove the closing paren
		return []ast.Expr{}, nil
	}

	argsList, err := p.parseArgsList()

	args, ok := argsList.([]ast.Expr)

	if err != nil || !ok {
		return nil, err
	}

	_, err = p.expect(token_type.RightParen, compilerErrors.ErrSyntaxExpectedRightParen)
	if err != nil {
		return nil, err
	}

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

	_, err = p.expect(token_type.Colon, compilerErrors.ErrSyntaxExpectedColon)
	if err != nil {
		return nil, err
	}

	return args, nil
}

func (p *Parser) parseArgsList() (any, error) {
	var args []ast.Expr

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

	_, err := p.expect(token_type.LeftBrace, compilerErrors.ErrSyntaxExpectedLeftBrace)
	if err != nil {
		return nil, err
	}

	for p.at().Type != token_type.RightBrace && p.notEOF() {
		stmt, err := p.ParseStmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}

	_, err = p.expect(token_type.RightBrace, compilerErrors.ErrSyntaxExpectedRightBrace)
	if err != nil {
		return nil, err
	}

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
