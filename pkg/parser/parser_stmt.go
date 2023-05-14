package parser

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/lexer/token_type"
)

func (p *Parser) parseStmt() ast.Stmt {
	switch p.at().Type {
	case token_type.Const, token_type.Var:
		return p.parseVarDeclaration()
	case token_type.Fn:
		return p.parseFnDeclaration()
	default:
		return p.parseExpr()
	}
}

func (p *Parser) parseFnDeclaration() ast.Stmt {
	p.subtract() // consume 'fn'

	name := p.expect(token_type.Identifier, "Expected identifier name following 'fn'")

	args := p.parseArgs()

	var params = make([]string, len(args))

	for i, arg := range args {
		if arg.GetKind() != ast_types.Identifier {
			panic("Expected identifier name")
		}
		params[i] = arg.(ast.Identifier).Symbol
	}

	p.expect(token_type.LeftBrace, "Expected '{'")

	var body []ast.Stmt

	for p.at().Type != token_type.RightBrace && p.at().Type != token_type.EOF {
		body = append(body, p.parseStmt())
	}

	p.expect(token_type.RightBrace, "Expected '}'")

	return ast.FunctionDeclaration{
		Kind:   ast_types.FunctionDeclaration,
		Name:   name.Value,
		Params: params,
		Body:   body,
	}
}

func (p *Parser) parseVarDeclaration() ast.Stmt {
	isConstant := p.subtract().Type == token_type.Const
	identifier := p.expect(token_type.Identifier, "Expected identifier name following 'const' or 'var'").Value

	if p.at().Type != token_type.Equals {
		if isConstant {
			panic("Must declare constant value")
		}

		return ast.VariableDeclaration{
			Kind:       ast_types.VariableDeclaration,
			Constant:   false,
			Identifier: identifier,
			Value:      nil,
		}
	}

	p.expect(token_type.Equals, "Expected '='")
	declaration := ast.VariableDeclaration{
		Kind:       ast_types.VariableDeclaration,
		Constant:   isConstant,
		Identifier: identifier,
		Value:      p.parseExpr(),
	}

	return declaration
}
