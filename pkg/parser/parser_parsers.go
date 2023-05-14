package parser

import (
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/lexer/lexerTypes"
)

func (p *Parser) parseStmt() ast.Stmt {
	switch p.at().Type {
	case lexerTypes.Const, lexerTypes.Var:
		return p.parseVarDeclaration()
	default:
		return p.parseExpr()
	}
}

func (p *Parser) parseVarDeclaration() ast.Stmt {
	isConstant := p.subtract().Type == lexerTypes.Const
	identifier := p.expect(lexerTypes.Identifier, "Expected identifier name following 'const' or 'var'").Value

	if p.at().Type != lexerTypes.Equals {
		if isConstant {
			panic("Must declare constant value")
		}

		return ast.VariableDeclaration{
			Kind:       astTypes.VariableDeclaration,
			Constant:   false,
			Identifier: identifier,
			Value:      nil,
		}
	}

	p.expect(lexerTypes.Equals, "Expected '='")
	declaration := ast.VariableDeclaration{
		Kind:       astTypes.VariableDeclaration,
		Constant:   isConstant,
		Identifier: identifier,
		Value:      p.parseExpr(),
	}

	return declaration
}
