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
		return p.parseFnDeclaration()
	case token_type.If:
		return p.parseIfStatement()
	case token_type.Switch:
		return p.parseSwitchStatement()
	default:
		return p.parseExpr()
	}
}

func (p *Parser) parseSwitchStatement() (ast.Stmt, error) {
	p.subtract() // consume 'switch'

	args, err := p.parseArgs(token_type.Switch)

	if err != nil {
		return nil, err
	}

	var caseStmts []ast.CaseStatement
	var defaultStmt ast.CaseStatement
	condition := args[0]

	p.expect(token_type.LeftBrace, compilerErrors.ErrSyntaxExpectedLeftBrace)

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.EOF {
		if p.at().Type == token_type.Case {
			p.subtract() // consume 'case'

			caseCondition, err := p.parseArgs(token_type.Case)

			if err != nil {
				return nil, err
			}

			p.expect(token_type.Colon, compilerErrors.ErrSyntaxExpectedColon)

			body, err := p.parseSwitchBodyStmt()

			if err != nil {
				return nil, err
			}

			caseStmts = append(caseStmts, ast.CaseStatement{
				Condition: caseCondition,
				Body:      body,
			})
		}

		if p.at().Type == token_type.Default {
			p.subtract() // consume 'default'
			p.expect(token_type.Colon, compilerErrors.ErrSyntaxExpectedColon)
			body, err := p.parseSwitchBodyStmt()

			if err != nil {
				return nil, err
			}

			defaultStmt = ast.CaseStatement{
				Condition: nil,
				Body:      body,
			}
			break
		}
	}

	p.expect(token_type.RightBrace, compilerErrors.ErrSyntaxExpectedRightBrace)

	return ast.SwitchStatement{
		Kind:        ast_types.SwitchStatement,
		Condition:   condition,
		CaseStmts:   caseStmts,
		DefaultStmt: defaultStmt,
	}, nil
}

func (p *Parser) parseIfStatement() (ast.Stmt, error) {
	var elseBody []ast.Stmt = nil

	p.subtract() // consume 'if'

	args, err := p.parseArgs(token_type.If)

	if err != nil {
		return nil, err
	}

	condition := args[0]

	body, err := p.parseBlockBodyStmt()

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
		elseBody, err = p.parseBlockBodyStmt()
		if err != nil {
			return nil, err
		}
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

	for p.at().Type == token_type.Else && p.atNext().Type == token_type.If && p.at().Type != token_type.EOF {
		p.subtract(2) // consume 'else' & 'if'
		args, err := p.parseArgs(token_type.If)

		if err != nil {
			return nil, err
		}

		condition := args[0]

		body, err := p.parseBlockBodyStmt()

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

	body, err := p.parseBlockBodyStmt()

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
