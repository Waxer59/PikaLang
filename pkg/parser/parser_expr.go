package parser

import (
	"errors"
	"os"
	compilerErrors "pika/internal/errors"
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/lexer/token_type"
	"strconv"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

func (p *Parser) parseExpr() (ast.Expr, error) {
	expr, err := p.parseAssigmentExpr()
	return expr, err
}

func (p *Parser) parseAdditiveExpr() (ast.Expr, error) {
	left, err := p.parseExponentialExpr()

	if err != nil {
		return nil, err
	}

	for slices.Contains(ast_types.AdditiveExpr, p.at().Value) {
		var op = p.subtract().Value
		right, err := p.parseExponentialExpr()

		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left, nil
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	tk := p.at().Type
	errorMsg := ""

	switch tk {
	case token_type.BooleanLiteral:
		b, err := strconv.ParseBool(p.subtract().Value)
		if err != nil {
			errorMsg = err.Error()
			break
		}
		return ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: b}
	case token_type.Null:
		p.subtract() // consume 'null'
		return ast.NullLiteral{Kind: ast_types.NullLiteral, Value: nil}
	case token_type.NaN:
		p.subtract() // consume 'NaN'
		return ast.NaNLiteral{Kind: ast_types.NaNLiteral, Value: nil}
	case token_type.Identifier:
		return ast.Identifier{Kind: ast_types.Identifier, Symbol: p.subtract().Value}
	case token_type.Number:
		n, err := strconv.ParseFloat(p.subtract().Value, 64)

		if err != nil {
			errorMsg = err.Error()
			break
		}

		return ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: n}
	case token_type.DoubleQoute:
		p.subtract() // consume '"'
		value := p.subtract().Value
		p.expect(token_type.DoubleQoute, compilerErrors.ErrSyntaxExpectedDoubleQoute)
		return ast.StringLiteral{
			Kind:  ast_types.StringLiteral,
			Value: value,
		}
	case token_type.LeftParen:
		p.subtract() // consume '('
		value, err := p.parseExpr()

		if err != nil {
			errorMsg = err.Error()
			break
		}

		p.expect(token_type.RightParen, compilerErrors.ErrSyntaxExpectedRightParen)
		return value
	default:
		errorMsg = "Expected an expression"
	}

	color.Red("Something went wrong with parsing: " + errorMsg)
	os.Exit(0)
	return nil
}

func (p *Parser) parseMemberExpr() (ast.Expr, error) {
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
				return nil, errors.New(compilerErrors.ErrFuncExpectedIdentifer)
			}
		case token_type.LeftBracket:
			var err error

			computed = true
			property, err = p.parseExpr()

			if err != nil {
				return nil, err
			}

			p.expect(token_type.RightBracket, compilerErrors.ErrSyntaxExpectedRightBracket)
		}

		obj = ast.MemberExpr{
			Kind:     ast_types.MemberExpr,
			Object:   obj,
			Property: property,
			Computed: computed,
		}
	}

	return obj, nil
}

func (p *Parser) parseCallMemberExpr() (ast.Expr, error) {
	member, err := p.parseMemberExpr()
	if err != nil {
		return nil, err
	}

	if p.at().Type == token_type.LeftParen {
		callExpr, err := p.parseCallExpr(member)
		return callExpr, err
	}

	return member, nil
}

func (p *Parser) parseMultiplicativeExpr() (ast.Expr, error) {
	left, err := p.parseCallMemberExpr()

	if err != nil {
		return nil, err
	}

	for slices.Contains(ast_types.MultiplicativeExpr, p.at().Value) {
		var op = p.subtract().Value
		right, err := p.parseCallMemberExpr()

		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left, nil
}

func (p *Parser) parseSufixUpdateExpr() (ast.Expr, error) {
	if p.at().Type == token_type.Identifier && p.atNext().Value == "++" || p.atNext().Value == "--" {
		argument := p.parsePrimaryExpr()
		ident, ok := argument.(ast.Identifier)

		if !ok {
			return nil, errors.New(compilerErrors.ErrSyntaxExpectedIdentifier)
		}

		op := p.subtract().Value // consume '++' or '--'

		return ast.UpdateExpr{
			Kind:     ast_types.UpdateExpr,
			Operator: op,
			Argument: ident,
			Prefix:   false,
		}, nil
	}

	return p.parseMultiplicativeExpr()
}

func (p *Parser) parsePrefixUpdateExpr() (ast.Expr, error) {
	if p.at().Value == "++" || p.at().Value == "--" {
		op := p.subtract().Value // consume '++' or '--'
		argument := p.parsePrimaryExpr()
		ident, ok := argument.(ast.Identifier)
		if !ok {
			return nil, errors.New(compilerErrors.ErrSyntaxExpectedIdentifier)
		}

		return ast.UpdateExpr{
			Kind:     ast_types.UpdateExpr,
			Operator: op,
			Argument: ident,
			Prefix:   true,
		}, nil
	}
	return p.parseSufixUpdateExpr()
}

func (p *Parser) parseNegativeAndPositiveExpr() (ast.Expr, error) {
	if p.at().Value == "+" || p.at().Value == "-" {
		op := p.subtract().Value // consume '-' or '+'
		argument, err := p.parseNegativeAndPositiveExpr()
		if err != nil {
			return nil, err
		}
		return ast.UnaryExpr{
			Kind:     ast_types.UnaryExpr,
			Operator: op,
			Argument: argument,
			Prefix:   true,
		}, nil
	}

	return p.parsePrefixUpdateExpr()
}

func (p *Parser) parseLogicalNotExpr() (ast.Expr, error) {
	if p.at().Type == token_type.Bang {
		op := p.subtract().Value // consume '!'
		argument, err := p.parseLogicalNotExpr()
		if err != nil {
			return nil, err
		}
		return ast.UnaryExpr{
			Kind:     ast_types.UnaryExpr,
			Operator: op,
			Argument: argument,
			Prefix:   false,
		}, nil
	}

	return p.parseNegativeAndPositiveExpr()
}

func (p *Parser) parseExponentialExpr() (ast.Expr, error) {
	left, err := p.parseLogicalNotExpr()

	if err != nil {
		return nil, err
	}

	for p.at().Value == "**" {
		op := p.subtract().Value
		right, err := p.parseLogicalNotExpr()

		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left, nil
}

func (p *Parser) parseCallExpr(caller ast.Expr) (ast.Expr, error) {
	args, err := p.parseArgs(token_type.Fn)
	if err != nil {
		return nil, err
	}
	var callExpr ast.Expr = ast.CallExpr{
		Kind:   ast_types.CallExpr,
		Caller: caller,
		Args:   args,
	}

	if p.at().Type == token_type.LeftParen {
		var err error

		callExpr, err = p.parseCallExpr(callExpr)
		if err != nil {
			return nil, err
		}
	}

	return callExpr, nil
}

func (p *Parser) parseObjectExpr() (ast.Expr, error) {
	if p.at().Type != token_type.LeftBrace {
		additiveExpr, err := p.parseAdditiveExpr()
		return additiveExpr, err
	}

	p.subtract() // advance post open brace

	properties := []ast.Property{}

	for p.notEOF() && p.at().Type != token_type.RightBrace {
		key := p.expect(token_type.Identifier, compilerErrors.ErrSyntaxExpectedKey).Value

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

		p.expect(token_type.Colon, compilerErrors.ErrSyntaxExpectedColon)
		value, err := p.parseExpr()

		if err != nil {
			return nil, err
		}

		properties = append(properties, ast.Property{
			Kind:  ast_types.Property,
			Key:   key,
			Value: value,
		})

		if p.at().Type != token_type.RightBrace {
			p.expect(token_type.Comma, compilerErrors.ErrSyntaxExpectedComma)
		}
	}

	p.expect(token_type.RightBrace, compilerErrors.ErrSyntaxExpectedRightBrace)
	return ast.ObjectLiteral{
		Kind:       ast_types.ObjectLiteral,
		Properties: properties,
	}, nil
}

func (p *Parser) parseArrayExpr() (ast.Expr, error) {
	if p.at().Type != token_type.LeftBracket {
		additiveExpr, err := p.parseObjectExpr()
		return additiveExpr, err
	}

	p.subtract() // advance post open bracket

	elements := []ast.Expr{}

	for p.notEOF() && p.at().Type != token_type.RightBracket {
		val, err := p.parseExpr()

		if err != nil {
			return nil, err
		}

		if p.at().Type != token_type.RightBracket {
			p.expect(token_type.Comma, compilerErrors.ErrSyntaxExpectedComma)
		}

		elements = append(elements, val)
	}

	p.expect(token_type.RightBracket, compilerErrors.ErrSyntaxExpectedRightBracket)
	return ast.ArrayLiteral{
		Kind:     ast_types.ArrayLiteral,
		Elements: elements,
	}, nil
}

func (p *Parser) parseComparisonExpr() (ast.Expr, error) {
	left, err := p.parseArrayExpr()

	if err != nil {
		return nil, err
	}

	for slices.Contains(ast_types.ComparisonExpr, p.at().Value) && p.notEOF() {
		op := p.subtract().Value // consume operator
		right, err := p.parseObjectExpr()

		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left, nil
}

func (p *Parser) parseEqualityExpr() (ast.Expr, error) {
	left, err := p.parseComparisonExpr()
	if err != nil {
		return nil, err
	}

	for slices.Contains(ast_types.EqualityExpr, p.at().Value) && p.notEOF() {
		op := p.subtract().Value // consume operator
		right, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		left = ast.BinaryExpr{
			Kind:     ast_types.BinaryExpr,
			Left:     left,
			Right:    right,
			Operator: op,
		}
	}

	return left, nil
}

func (p *Parser) parseLogicalAndExpr() (ast.Expr, error) {
	left, err := p.parseEqualityExpr()

	if err != nil {
		return nil, nil
	}

	for p.at().Type == token_type.And && p.notEOF() {
		p.subtract() // consume '&&'
		right, err := p.parseExpr()
		if err != nil {
			return nil, nil
		}
		left = ast.LogicalExpr{
			Kind:     ast_types.LogicalExpr,
			Left:     left,
			Right:    right,
			Operator: "&&",
		}
	}

	return left, nil
}

func (p *Parser) parseLogicalOrExpr() (ast.Expr, error) {
	left, err := p.parseLogicalAndExpr()

	if err != nil {
		return nil, nil
	}

	for p.at().Type == token_type.Or && p.notEOF() {
		p.subtract() // consume '||'
		right, err := p.parseExpr()
		if err != nil {
			return nil, nil
		}
		left = ast.LogicalExpr{
			Kind:     ast_types.LogicalExpr,
			Left:     left,
			Right:    right,
			Operator: "||",
		}
	}

	return left, nil
}

func (p *Parser) parseTernaryExpr() (ast.Expr, error) {
	condition, err := p.parseLogicalOrExpr()

	if err != nil {
		return nil, err
	}

	if p.at().Type == token_type.QuestionMark {
		p.subtract() // consume '?'
		consequent, err := p.parseExpr()

		if err != nil {
			return nil, err
		}

		p.expect(token_type.Colon, compilerErrors.ErrSyntaxExpectedColon)

		alternate, err := p.parseExpr()

		if err != nil {
			return nil, err
		}

		return ast.ConditionalExpr{
			Kind:       ast_types.ConditionalExpr,
			Condition:  condition,
			Consequent: consequent,
			Alternate:  alternate,
		}, nil
	}

	return condition, nil
}

func (p *Parser) parseAssigmentExpr() (ast.Expr, error) {
	left, err := p.parseTernaryExpr()

	if err != nil {
		return nil, err
	}

	if slices.Contains(token_type.AssigmentOperators, p.at().Type) {
		op := p.subtract().Value // consume assigment operator
		value, err := p.parseTernaryExpr()
		if err != nil {
			return nil, err
		}
		return ast.AssigmentExpr{
			Kind:     ast_types.AssigmentExpr,
			Assigne:  left,
			Value:    value,
			Operator: op,
		}, nil
	}

	return left, nil
}
