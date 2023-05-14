package parser

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/lexer/token_type"
	"strconv"

	"golang.org/x/exp/slices"
)

func (p *Parser) parseAdditiveExpr() ast.Expr {
	var left = p.parseMultiplicativeExpr()

	for slices.Contains(ast_types.AdditiveExpr, p.at().Value) {
		var op = p.subtract().Value
		var right = p.parseMultiplicativeExpr()
		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left
}

func (p *Parser) parseMultiplicativeExpr() ast.Expr {
	var left = p.parseCallMemberExpr()

	for slices.Contains(ast_types.MultiplicativeExpr, p.at().Value) {
		var op = p.subtract().Value
		var right = p.parseCallMemberExpr()
		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left
}

func (p *Parser) parseCallMemberExpr() ast.Expr {
	member := p.parseMemberExpr()

	if p.at().Type == token_type.LeftParen {
		return p.parseCallExpr(member)
	}

	return member
}

func (p *Parser) parseCallExpr(caller ast.Expr) ast.Expr {
	var callExpr ast.Expr = ast.CallExpr{
		Kind:   ast_types.CallExpr,
		Caller: caller,
		Args:   p.parseArgs(),
	}

	if p.at().Type == token_type.LeftParen {
		callExpr = p.parseCallExpr(callExpr)
	}

	return callExpr
}

func (p *Parser) parseArgs() []ast.Expr {
	p.expect(token_type.LeftParen, "Expected '('")

	args := []ast.Expr{}

	if p.at().Type != token_type.RightParen {
		args = p.parseArgsList()
	}

	p.expect(token_type.RightParen, "Expected ')'")

	return args
}

func (p *Parser) parseArgsList() []ast.Expr {
	args := []ast.Expr{p.parseAssigmentExpr()}

	for p.at().Type == token_type.Comma && p.subtract().Type == token_type.Comma {
		args = append(args, p.parseAssigmentExpr())
	}

	return args
}

func (p *Parser) parseMemberExpr() ast.Expr {
	obj := p.parsePrimaryExpr()

	for p.at().Type == token_type.Dot || p.at().Type == token_type.LeftBracket {
		operator := p.subtract()
		var property ast.Expr
		var computed bool

		switch operator.Type {
		case token_type.Dot:
			computed = false
			property = p.parsePrimaryExpr()

			if property.GetKind() != ast_types.Identifier {
				panic("Expected identifier")
			}
		case token_type.LeftBracket:
			computed = true
			property = p.parseExpr()

			p.expect(token_type.RightBracket, "Expected ']'")
		}

		obj = ast.MemberExpr{
			Kind:     ast_types.MemberExpr,
			Object:   obj,
			Property: property,
			Computed: computed,
		}
	}

	return obj
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseAssigmentExpr()
}

func (p *Parser) parseObjectExpr() ast.Expr {
	if p.at().Type != token_type.LeftBrace {
		return p.parseAdditiveExpr()
	}

	p.subtract() // advance post open brace

	properties := []ast.Property{}

	for p.notEOF() && p.at().Type != token_type.RightBrace {
		key := p.expect(token_type.Identifier, "Expected a key").Value

		// Allows shorthand syntax: { key, } && { key }
		switch p.at().Type {
		case token_type.Comma:
			p.subtract()
			properties = append(properties, ast.Property{
				Kind:  ast_types.Property,
				Key:   key,
				Value: nil,
			})
			continue
		case token_type.RightBrace:
			properties = append(properties, ast.Property{
				Kind:  ast_types.Property,
				Key:   key,
				Value: nil,
			})
			continue
		}

		p.expect(token_type.Colon, "Expected ':'")
		value := p.parseExpr()

		properties = append(properties, ast.Property{
			Kind:  ast_types.Property,
			Key:   key,
			Value: value,
		})

		if p.at().Type != token_type.RightBrace {
			p.expect(token_type.Comma, "Expected ',' or '}'")
		}
	}

	p.expect(token_type.RightBrace, "Expected '}'")
	return ast.ObjectLiteral{
		Kind:       ast_types.ObjectLiteral,
		Properties: properties,
	}
}

func (p *Parser) parseAssigmentExpr() ast.Expr {
	var left = p.parseObjectExpr()

	if p.at().Type == token_type.Equals {
		p.subtract() // consume '='
		value := p.parseAssigmentExpr()
		return ast.AssigmentExpr{
			Kind:    ast_types.AssigmentExpr,
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
	case token_type.BooleanLiteral:
		b, err := strconv.ParseBool(p.subtract().Value)
		if err != nil {
			panic("Something went wrong with parsing: " + err.Error())
		}
		return ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: b}
	case token_type.Null:
		p.subtract() // consume 'null'
		return ast.NullLiteral{Kind: ast_types.NullLiteral, Value: nil}
	case token_type.Identifier:
		return ast.Identifier{Kind: ast_types.Identifier, Symbol: p.subtract().Value}
	case token_type.Number:
		n, err := strconv.Atoi(p.subtract().Value)
		if err != nil {
			panic("Something went wrong with parsing: " + err.Error())
		}
		return ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: n}
	case token_type.LeftParen:
		p.subtract() // consume '('
		value := p.parseExpr()
		p.expect(token_type.RightParen, "Expected ')'")
		return value
	default:
		errorMsg = "Expected an expression"
	}

	panic("Something went wrong with parsing: " + errorMsg)
}
