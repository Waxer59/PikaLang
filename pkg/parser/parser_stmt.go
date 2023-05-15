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
	default:
		expr, err := p.parseExpr()
		return expr, err
	}
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

	var body []ast.Stmt

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.EOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}

	p.expect(token_type.RightBrace, string(compilerErrors.ErrSyntaxExpectedRightBrace))

	return ast.FunctionDeclaration{
		Kind:   ast_types.FunctionDeclaration,
		Name:   name.Value,
		Params: params,
		Body:   body,
	}, nil
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
