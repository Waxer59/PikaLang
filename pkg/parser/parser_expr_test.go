package parser_test

import (
	"reflect"
	"testing"

	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/ast/ast_types"
	"github.com/Waxer59/PikaLang/pkg/lexer/token_type"
	"github.com/Waxer59/PikaLang/pkg/parser"
)

type ParserTest struct {
	input        []token_type.Token
	expectedExpr []ast.Expr
	expectedErr  error
}

func testParseExpr(t *testing.T, tests []ParserTest, p *parser.Parser) {
	for _, test := range tests {
		p.SetTokens(test.input)
		stmt, err := p.ParseStmt()

		t.Log(test)

		// Check error
		if test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("Expected error: %v, but got: %v", test.expectedErr, err)
		}

		// Check tokens
		if !reflect.DeepEqual(stmt, test.expectedExpr[0]) {
			t.Errorf("Expected expr: %v, but got: %v", test.expectedExpr, stmt)
		}
	}
}

func TestParseAssigmentExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Identifier, Value: "a"},
				{Type: token_type.Equals, Value: "="},
				{Type: token_type.Number, Value: "1"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.AssigmentExpr{
					Kind: ast_types.AssigmentExpr,
					Assigne: ast.Identifier{
						Kind:   ast_types.Identifier,
						Symbol: "a",
					},
					Value: ast.NumericLiteral{
						Kind:  ast_types.NumericLiteral,
						Value: 1,
					},
					Operator: "=",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Identifier, Value: "a"},
				{Type: token_type.Equals, Value: "="},
				{Type: token_type.Number, Value: "1"},
				{Type: token_type.Plus, Value: "+"},
				{Type: token_type.Number, Value: "1"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.AssigmentExpr{
					Kind: ast_types.AssigmentExpr,
					Assigne: ast.Identifier{
						Kind:   ast_types.Identifier,
						Symbol: "a",
					},
					Value: ast.BinaryExpr{
						Kind: ast_types.BinaryExpr,
						Left: ast.NumericLiteral{
							Kind:  ast_types.NumericLiteral,
							Value: 1,
						},
						Right: ast.NumericLiteral{
							Kind:  ast_types.NumericLiteral,
							Value: 1,
						},
						Operator: "+",
					},
					Operator: "=",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Identifier, Value: "a"},
				{Type: token_type.Equals, Value: "="},
				{Type: token_type.Number, Value: "1"},
				{Type: token_type.QuestionMark, Value: "?"},
				{Type: token_type.Number, Value: "1"},
				{Type: token_type.Colon, Value: ":"},
				{Type: token_type.Number, Value: "1"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.AssigmentExpr{
					Kind: ast_types.AssigmentExpr,
					Assigne: ast.Identifier{
						Kind:   ast_types.Identifier,
						Symbol: "a",
					},
					Value: ast.ConditionalExpr{
						Kind: ast_types.ConditionalExpr,
						Condition: ast.NumericLiteral{
							Kind:  ast_types.NumericLiteral,
							Value: 1,
						},
						Consequent: ast.NumericLiteral{
							Kind:  ast_types.NumericLiteral,
							Value: 1,
						},
						Alternate: ast.NumericLiteral{
							Kind:  ast_types.NumericLiteral,
							Value: 1,
						},
					},
					Operator: "=",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParsePrefixUpdateExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Decrement, Value: "++"},
				{Type: token_type.Identifier, Value: "x"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.UpdateExpr{
					Kind:     ast_types.UpdateExpr,
					Operator: "++",
					Argument: ast.Identifier{Kind: ast_types.Identifier, Symbol: "x"},
					Prefix:   true,
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Decrement, Value: "--"},
				{Type: token_type.Identifier, Value: "y"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.UpdateExpr{
					Kind:     ast_types.UpdateExpr,
					Operator: "--",
					Argument: ast.Identifier{Kind: ast_types.Identifier, Symbol: "y"},
					Prefix:   true,
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseLogicalNotExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Bang, Value: "!"},
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.UnaryExpr{
					Kind:     ast_types.UnaryExpr,
					Operator: "!",
					Argument: ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Prefix:   false,
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Bang, Value: "!"},
				{Type: token_type.LeftParen, Value: "("},
				{Type: token_type.BooleanLiteral, Value: "false"},
				{Type: token_type.RightParen, Value: ")"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.UnaryExpr{
					Kind:     ast_types.UnaryExpr,
					Operator: "!",
					Argument: ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: false},
					Prefix:   false,
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseCallExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Identifier, Value: "func"},
				{Type: token_type.LeftParen, Value: "("},
				{Type: token_type.Number, Value: "42"},
				{Type: token_type.RightParen, Value: ")"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.CallExpr{
					Kind:   ast_types.CallExpr,
					Caller: ast.Identifier{Kind: ast_types.Identifier, Symbol: "func"},
					Args: []ast.Expr{
						ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 42},
					},
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Identifier, Value: "add"},
				{Type: token_type.LeftParen, Value: "("},
				{Type: token_type.Number, Value: "2"},
				{Type: token_type.Comma, Value: ","},
				{Type: token_type.Number, Value: "3"},
				{Type: token_type.RightParen, Value: ")"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.CallExpr{
					Kind:   ast_types.CallExpr,
					Caller: ast.Identifier{Kind: ast_types.Identifier, Symbol: "add"},
					Args: []ast.Expr{
						ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 2},
						ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 3},
					},
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseNegativeAndPositiveExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Plus, Value: "+"},
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.UnaryExpr{
					Kind:     ast_types.UnaryExpr,
					Operator: "+",
					Argument: ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
					Prefix:   true,
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Minus, Value: "-"},
				{Type: token_type.Number, Value: "10"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.UnaryExpr{
					Kind:     ast_types.UnaryExpr,
					Operator: "-",
					Argument: ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 10},
					Prefix:   true,
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseExponentialExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "2"},
				{Type: token_type.BinaryOperator, Value: "**"},
				{Type: token_type.Number, Value: "3"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 2},
					Right:    ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 3},
					Operator: "**",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseAdditiveExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.Plus, Value: "+"},
				{Type: token_type.Number, Value: "3"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
					Right:    ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 3},
					Operator: "+",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "10"},
				{Type: token_type.Minus, Value: "-"},
				{Type: token_type.Number, Value: "7"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 10},
					Right:    ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 7},
					Operator: "-",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseArrowFunctionExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.LeftParen, Value: "("},
				{Type: token_type.Identifier, Value: "x"},
				{Type: token_type.Comma, Value: ","},
				{Type: token_type.Identifier, Value: "y"},
				{Type: token_type.RightParen, Value: ")"},
				{Type: token_type.Arrow, Value: "=>"},
				{Type: token_type.LeftBrace, Value: "{"},
				{Type: token_type.Number, Value: "x"},
				{Type: token_type.Plus, Value: "+"},
				{Type: token_type.Number, Value: "y"},
				{Type: token_type.RightBrace, Value: "}"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.ArrowFunctionExpr{
					Kind: ast_types.ArrowFunctionExpr,
					Params: []ast.Identifier{
						{Kind: ast_types.Identifier, Symbol: "x"},
						{Kind: ast_types.Identifier, Symbol: "y"},
					},
					Body: []ast.Stmt{
						nil, nil,
					},
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseObjectExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.LeftBrace, Value: "{"},
				{Type: token_type.Identifier, Value: "name"},
				{Type: token_type.Colon, Value: ":"},
				{Type: token_type.DoubleQoute, Value: "\""},
				{Type: token_type.DoubleQoute, Value: "John"},
				{Type: token_type.DoubleQoute, Value: "\""},
				{Type: token_type.Comma, Value: ","},
				{Type: token_type.Identifier, Value: "age"},
				{Type: token_type.Colon, Value: ":"},
				{Type: token_type.Number, Value: "30"},
				{Type: token_type.RightBrace, Value: "}"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.ObjectLiteral{
					Kind: ast_types.ObjectLiteral,
					Properties: []ast.Property{
						{
							Kind:  ast_types.Property,
							Key:   "name",
							Value: ast.StringLiteral{Kind: ast_types.StringLiteral, Value: "John"},
						},
						{
							Kind:  ast_types.Property,
							Key:   "age",
							Value: ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 30},
						},
					},
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseComparisonExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.Less, Value: "<"},
				{Type: token_type.Number, Value: "10"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
					Right:    ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 10},
					Operator: "<",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "20"},
				{Type: token_type.Greater, Value: ">"},
				{Type: token_type.Number, Value: "15"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 20},
					Right:    ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 15},
					Operator: ">",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseEqualityExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.Equals, Value: "=="},
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
					Right:    ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
					Operator: "==",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.NotEqual, Value: "!="},
				{Type: token_type.BooleanLiteral, Value: "false"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.BinaryExpr{
					Kind:     ast_types.BinaryExpr,
					Left:     ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Right:    ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: false},
					Operator: "!=",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseLogicalAndExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.And, Value: "&&"},
				{Type: token_type.BooleanLiteral, Value: "false"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.LogicalExpr{
					Kind:     ast_types.LogicalExpr,
					Left:     ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Right:    ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: false},
					Operator: "&&",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.And, Value: "&&"},
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.And, Value: "&&"},
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.LogicalExpr{
					Kind: ast_types.LogicalExpr,
					Left: ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Right: ast.LogicalExpr{
						Kind:     ast_types.LogicalExpr,
						Left:     ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
						Right:    ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
						Operator: "&&",
					},
					Operator: "&&",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseLogicalOrExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.Or, Value: "||"},
				{Type: token_type.BooleanLiteral, Value: "false"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.LogicalExpr{
					Kind:     ast_types.LogicalExpr,
					Left:     ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Right:    ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: false},
					Operator: "||",
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.Or, Value: "||"},
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.Or, Value: "||"},
				{Type: token_type.BooleanLiteral, Value: "false"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.LogicalExpr{
					Kind: ast_types.LogicalExpr,
					Left: ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Right: ast.LogicalExpr{
						Kind:     ast_types.LogicalExpr,
						Left:     ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
						Right:    ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: false},
						Operator: "||",
					},
					Operator: "||",
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}

func TestParseTernaryExpr(t *testing.T) {
	p := parser.New()

	tests := []ParserTest{
		{
			input: []token_type.Token{
				{Type: token_type.BooleanLiteral, Value: "true"},
				{Type: token_type.QuestionMark, Value: "?"},
				{Type: token_type.Number, Value: "10"},
				{Type: token_type.Colon, Value: ":"},
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.ConditionalExpr{
					Kind:       ast_types.ConditionalExpr,
					Condition:  ast.BooleanLiteral{Kind: ast_types.BooleanLiteral, Value: true},
					Consequent: ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 10},
					Alternate:  ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
				},
			},
			expectedErr: nil,
		},
		{
			input: []token_type.Token{
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.QuestionMark, Value: "?"},
				{Type: token_type.Identifier, Value: "x"},
				{Type: token_type.Colon, Value: ":"},
				{Type: token_type.Identifier, Value: "y"},
				{Type: token_type.QuestionMark, Value: "?"},
				{Type: token_type.Number, Value: "10"},
				{Type: token_type.Colon, Value: ":"},
				{Type: token_type.Number, Value: "20"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedExpr: []ast.Expr{
				ast.ConditionalExpr{
					Kind:      ast_types.ConditionalExpr,
					Condition: ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 5},
					Consequent: ast.Identifier{
						Kind:   ast_types.Identifier,
						Symbol: "x",
					},
					Alternate: ast.ConditionalExpr{
						Kind:       ast_types.ConditionalExpr,
						Condition:  ast.Identifier{Kind: ast_types.Identifier, Symbol: "y"},
						Consequent: ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 10},
						Alternate:  ast.NumericLiteral{Kind: ast_types.NumericLiteral, Value: 20},
					},
				},
			},
			expectedErr: nil,
		},
	}

	testParseExpr(t, tests, p)
}
