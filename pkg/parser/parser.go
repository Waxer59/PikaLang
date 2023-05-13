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

func (p *Parser) next() lexerTypes.Token {
	prev := p.at()
	p.tokens = p.tokens[1:]
	return prev
}

func (p *Parser) expect(typeExpected lexerTypes.TokenType, errMsg string) lexerTypes.Token {
	prev := p.next()
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
	parseExpr := p.parseExpr()
	return parseExpr
}

func (p *Parser) parseAdditiveExpr() ast.Expr {
	var left = p.parseMultiplicativeExpr()

	for slices.Contains(astTypes.AdditiveExpr, p.at().Value) {
		var op = p.next().Value
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
		var op = p.next().Value
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
		return ast.Identifier{Kind: astTypes.Identifier, Symbol: p.next().Value}
	case lexerTypes.Null:
		p.next()
		return ast.NullLiteral{Kind: astTypes.NullLiteral, Value: "null"}
	case lexerTypes.Number:
		n, err := strconv.Atoi(p.next().Value)
		if err != nil {
			panic("Something went wrong with parsing: " + err.Error())
		}
		return ast.NumericLiteral{Kind: astTypes.NumericLiteral, Value: n}
	case lexerTypes.RightParen:
		p.next()
		value := p.parseExpr()
		p.expect(lexerTypes.LeftParen, "Expected ')'")
		return value
	default:
		errorMsg = "Expected an expression"
	}

	panic("Something went wrong with parsing: " + errorMsg)
}
