package parser

import (
	"fmt"
	"pika/pkg/ast"
	"pika/pkg/ast/astTypes"
	"pika/pkg/lexer"
	"pika/pkg/lexer/lexerTypes"
	"strconv"

	"golang.org/x/exp/slices"
)

type Parser struct {
	tokens []lexerTypes.Token
}

func (p *Parser) at() lexerTypes.Token {
	return p.tokens[0]
}

func (p *Parser) subtract() lexerTypes.Token {
	prev := p.at()
	p.tokens = p.tokens[1:]
	return prev
}

func (p *Parser) expect(typeExpected lexerTypes.TokenType, errMsg string) lexerTypes.Token {
	prev := p.subtract()
	if (prev == lexerTypes.Token{} || prev.Type != typeExpected) {
		panic("Parser Error:\n" + errMsg)
	}

	return prev
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ProduceAST(input string) ast.Program {
	p.tokens = lexer.Tokenize(input)
	fmt.Println(p.tokens)

	program := ast.Program{
		Kind: astTypes.Program,
		Body: []ast.Stmt{},
	}

	for p.notEOF() {
		program.Body = append(program.Body, p.parseStmt())
	}

	return program
}

func (p *Parser) notEOF() bool {
	return p.at().Type != lexerTypes.EOF
}

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

	fmt.Println(p)

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
	return p.parseAdditiveExpr()
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
