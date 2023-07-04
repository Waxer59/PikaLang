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
		return p.parseVarConstDeclaration()
	case token_type.Fn:
		fn, err := p.parseFnDeclaration()
		return fn, err
	case token_type.If:
		ifStmt, err := p.parseIfStatement()
		return ifStmt, err
	case token_type.Switch:
	}

	expr, err := p.parseExpr()
	return expr, err
}

func (p *Parser) parseSwitchStatement() (ast.Stmt, error) {
	p.subtract() // consume 'switch'

	args, err := p.parseArgs(token_type.If)

	if err != nil {
		return nil, err
	}

	condition := args[0]

	p.expect(token_type.LeftBrace, compilerErrors.ErrSyntaxExpectedLeftBrace)
}

func (p *Parser) parseIfStatement() (ast.Stmt, error) {
	var elseBody []ast.Stmt = nil

	p.subtract() // consume 'if'

	args, err := p.parseArgs(token_type.If)

	if err != nil {
		return nil, err
	}

	condition := args[0]

	body, err := p.parseBodyStmt()

	if err != nil {
		return nil, err
	}

	// Verify for else if statements
	elseIfStmt, err := p.parseElseIfStatement()

	if err != nil {
		return nil, err
	}

	// Verify for else statements
	if p.at().Type == token_type.Else {
		p.subtract() // consume 'else'
		p.expect(token_type.LeftBrace, compilerErrors.ErrSyntaxExpectedLeftBrace)
		elseBody, err = p.parseBodyStmt()
		if err != nil {
			return nil, err
		}
		p.expect(token_type.RightBrace, compilerErrors.ErrSyntaxExpectedRightBrace)
	}

	return ast.IfStatement{
		Kind:       ast_types.IfStatement,
		Condition:  condition,
		Body:       body,
		ElseBody:   elseBody,
		ElseIfStmt: elseIfStmt,
	}, nil
}

func (p *Parser) parseElseIfStatement() ([]ast.ElseIfStatement, error) {
	var elseIfStmt []ast.ElseIfStatement = nil

	for p.at().Type == token_type.Else && p.next().Type == token_type.If && p.at().Type != token_type.EOF {
		p.subtract(2) // consume 'else' & 'if'
		args, err := p.parseArgs(token_type.If)

		if err != nil {
			return nil, err
		}

		condition := args[0]

		body, err := p.parseBodyStmt()

		if err != nil {
			return nil, err
		}

		elseIfStmt = append(elseIfStmt, ast.ElseIfStatement{
			Condition: condition,
			Body:      body,
		})
	}

	return elseIfStmt, nil
}

func (p *Parser) parseFnDeclaration() (ast.Stmt, error) {
	p.subtract() // consume 'fn'

	name := p.expect(token_type.Identifier, compilerErrors.ErrFuncExpectedIdentifer)

	args, err := p.parseArgs(token_type.Fn)

	if err != nil {
		return nil, err
	}

	var params = make([]string, len(args))

	for i, arg := range args {
		if arg.GetKind() != ast_types.Identifier {
			return nil, errors.New(compilerErrors.ErrFuncExpectedIdentifer)
		}
		params[i] = arg.(ast.Identifier).Symbol
	}

	body, err := p.parseBodyStmt()

	if err != nil {
		return nil, err
	}

	return ast.FunctionDeclaration{
		Kind:   ast_types.FunctionDeclaration,
		Name:   name.Value,
		Params: params,
		Body:   body,
	}, nil
}

func (p *Parser) parseBodyStmt() ([]ast.Stmt, error) {
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

func (p *Parser) parseVarConstDeclaration() (ast.Stmt, error) {
	isConstant := p.subtract().Type == token_type.Const
	identifier := p.expect(token_type.Identifier, compilerErrors.ErrVariableExpectedIdentifierNameFollowingConstOrVar).Value

	if p.at().Type != token_type.Equals {
		if isConstant {
			return nil, errors.New(compilerErrors.ErrVariableIsConstant)
		}

		return ast.VariableDeclaration{
			Kind:       ast_types.VariableDeclaration,
			Constant:   false,
			Identifier: identifier,
			Value:      nil,
		}, nil
	}

	p.expect(token_type.Equals, compilerErrors.ErrSyntaxExpectedAsignation)

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
