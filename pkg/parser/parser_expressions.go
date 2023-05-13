package parser

import (
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/lexer/lexerTypes"
	"strconv"

	"golang.org/x/exp/slices"
)

func (p *Parser) parseAdditiveExpr() ast.Expr {
	var left = p.parseMultiplicativeExpr()

	for slices.Contains(astTypes.AdditiveExpr, p.at().Value) {
		var op = p.subtract().Value
		var right = p.parseMultiplicativeExpr()
		left = ast.BinaryExpr{
			Kind:     astTypes.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left
}

func (p *Parser) parseMultiplicativeExpr() ast.Expr {
	var left = p.parsePrimaryExpr()

	for slices.Contains(astTypes.MultiplicativeExpr, p.at().Value) {
		var op = p.subtract().Value
		var right = p.parsePrimaryExpr()
		left = ast.BinaryExpr{
			Kind:     astTypes.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseAssigmentExpr()
}

func (p *Parser) parseAssigmentExpr() ast.Expr {
	var left = p.parseAdditiveExpr()

	if p.at().Type == lexerTypes.Equals {
		p.subtract() // consume '='
		value := p.parseAssigmentExpr()
		return ast.AssigmentExpr{
			Kind:    astTypes.AssigmentExpr,
			Assigne: left,
			Value:   value,
		}
	}

	return left
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	tk := p.at().Type
	errorMsg := ""

	switch tk {
	case lexerTypes.Identifier:
		return ast.Identifier{Kind: astTypes.Identifier, Symbol: p.subtract().Value}
	case lexerTypes.Number:
		n, err := strconv.Atoi(p.subtract().Value)
		if err != nil {
			panic("Something went wrong with parsing: " + err.Error())
		}
		return ast.NumericLiteral{Kind: astTypes.NumericLiteral, Value: n}
	case lexerTypes.RightParen:
		p.subtract()
		value := p.parseExpr()
		p.expect(lexerTypes.LeftParen, "Expected ')'")
		return value
	default:
		errorMsg = "Expected an expression"
	}

	panic("Something went wrong with parsing: " + errorMsg)
}
