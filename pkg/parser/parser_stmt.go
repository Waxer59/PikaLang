package parser

import (
	"errors"
	compilerErrors "pika/internal/errors"
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/lexer/token_type"
)

func (p *Parser) parseStmt() (ast.Stmt, error) {
	switch p.at().Type {
	case token_type.Const, token_type.Var:
		return p.parseVarDeclaration()
	case token_type.Fn:
		fn, err := p.parseFnDeclaration()
		return fn, err
	case token_type.If:
		cond, err := p.parseIfStatement()
		return cond, err
	default:
		expr, err := p.parseExpr()
		return expr, err
	}
}

func (p *Parser) parseIfStatement() (ast.Stmt, error) {
	var elseBody []ast.Stmt = nil
	var elseIfStmt []ast.ElseIfStatement = nil

	p.subtract() // consume 'if', 'else', or 'else if'

	p.expect(token_type.LeftParen, string(compilerErrors.ErrSyntaxExpectedLeftParen))

	condition, err := p.parseExpr()

	if err != nil {
		return nil, err
	}

	p.expect(token_type.RightParen, string(compilerErrors.ErrSyntaxExpectedRightParen))

	p.expect(token_type.LeftBrace, string(compilerErrors.ErrSyntaxExpectedLeftBrace))

	body, err := p.parseBodyStmt()

	if err != nil {
		return nil, err
	}

	p.expect(token_type.RightBrace, string(compilerErrors.ErrSyntaxExpectedRightBrace))

	for p.at().Type == token_type.Else && p.next().Type == token_type.If && p.at().Type != token_type.EOF {
		p.subtract() // consume 'else'
		p.subtract() // consume 'if'
		p.expect(token_type.LeftParen, string(compilerErrors.ErrSyntaxExpectedLeftParen))
		elseIfCondition, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		p.expect(token_type.RightParen, string(compilerErrors.ErrSyntaxExpectedRightParen))
		p.expect(token_type.LeftBrace, string(compilerErrors.ErrSyntaxExpectedLeftBrace))
		elseIfBody, err := p.parseBodyStmt()
		if err != nil {
			return nil, err
		}
		p.expect(token_type.RightBrace, string(compilerErrors.ErrSyntaxExpectedRightBrace))
		elseIfStmt = append(elseIfStmt, ast.ElseIfStatement{
			Condition: elseIfCondition,
			Body:      elseIfBody,
		})
	}

	if p.at().Type == token_type.Else {
		p.subtract() // consume 'else'
		p.expect(token_type.LeftBrace, string(compilerErrors.ErrSyntaxExpectedLeftBrace))
		elseBody, err = p.parseBodyStmt()
		if err != nil {
			return nil, err
		}
		p.expect(token_type.RightBrace, string(compilerErrors.ErrSyntaxExpectedRightBrace))
	}

	return ast.IfStatement{
		Kind:       ast_types.IfStatement,
		Condition:  condition,
		Body:       body,
		ElseBody:   elseBody,
		ElseIfStmt: elseIfStmt,
	}, nil
}

func (p *Parser) parseFnDeclaration() (ast.Stmt, error) {
	p.subtract() // consume 'fn'

	name := p.expect(token_type.Identifier, string(compilerErrors.ErrFuncExpectedIdentifer))

	args, err := p.parseArgs()

	if err != nil {
		return nil, err
	}

	var params = make([]string, len(args))

	for i, arg := range args {
		if arg.GetKind() != ast_types.Identifier {
			return nil, errors.New(string(compilerErrors.ErrFuncExpectedIdentifer))
		}
		params[i] = arg.(ast.Identifier).Symbol
	}

	p.expect(token_type.LeftBrace, string(compilerErrors.ErrSyntaxExpectedLeftBrace))

	body, err := p.parseBodyStmt()

	if err != nil {
		return nil, err
	}

	p.expect(token_type.RightBrace, string(compilerErrors.ErrSyntaxExpectedRightBrace))

	return ast.FunctionDeclaration{
		Kind:   ast_types.FunctionDeclaration,
		Name:   name.Value,
		Params: params,
		Body:   body,
	}, nil
}

func (p *Parser) parseBodyStmt() ([]ast.Stmt, error) {
	var body []ast.Stmt

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.EOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}

	return body, nil
}

func (p *Parser) parseVarDeclaration() (ast.Stmt, error) {
	isConstant := p.subtract().Type == token_type.Const
	identifier := p.expect(token_type.Identifier, string(compilerErrors.ErrVariableExpectedIdentifierNameFollowingConstOrVar)).Value

	if p.at().Type != token_type.Equals {
		if isConstant {
			return nil, errors.New(string(compilerErrors.ErrVariableIsConstant))
		}

		return ast.VariableDeclaration{
			Kind:       ast_types.VariableDeclaration,
			Constant:   false,
			Identifier: identifier,
			Value:      nil,
		}, nil
	}

	p.expect(token_type.Equals, string(compilerErrors.ErrSyntaxExpectedAsignation))

	expr, err := p.parseExpr()

	if err != nil {
		return nil, err
	}

	declaration := ast.VariableDeclaration{
		Kind:       ast_types.VariableDeclaration,
		Constant:   isConstant,
		Identifier: identifier,
		Value:      expr,
	}

	return declaration, nil
}
