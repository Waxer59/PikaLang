package parser_test

import (
	"reflect"
	"testing"

	"github.com/Waxer59/PikaLang/pkg/ast"
	"github.com/Waxer59/PikaLang/pkg/ast/ast_types"
	"github.com/Waxer59/PikaLang/pkg/lexer"
	"github.com/Waxer59/PikaLang/pkg/parser"
)

type ParserTest struct {
	input        string
	expectedExpr []ast.Expr
	expectedErr  error
}

func testParseExpr(t *testing.T, tests []ParserTest, p *parser.Parser) {
	for _, test := range tests {
		tokens, err := lexer.Tokenize(test.input)

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		p.SetTokens(tokens)
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
			input: "a = 1",
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
			input: "a = 1 + 1",
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
			input: "a = 1 ? 1 : 1",
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
			input: "++x",
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
			input: "--y",
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
			input: "!true",
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
			input: "!(false)",
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
			input: "func(42)",
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
			input: "add(2, 3)",
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
			input: "+5",
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
			input: "-10",
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
			input: "2 ** 3",
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
			input: "5 + 3",
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
			input: "10 - 7",
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
			input: "(x,y)=>{x+y}",
			expectedExpr: []ast.Expr{
				ast.ArrowFunctionExpr{
					Kind: ast_types.ArrowFunctionExpr,
					Params: []ast.Identifier{
						{Kind: ast_types.Identifier, Symbol: "x"},
						{Kind: ast_types.Identifier, Symbol: "y"},
					},
					Body: []ast.Stmt{
						ast.BinaryExpr{
							Kind: ast_types.BinaryExpr,
							Left: ast.Identifier{
								Kind:   ast_types.Identifier,
								Symbol: "x",
							},
							Right: ast.Identifier{
								Kind:   ast_types.Identifier,
								Symbol: "y",
							},
							Operator: "+",
						},
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
			input: "{name: \"John\", age: 30 }",
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
			input: "5 < 10",
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
			input: "20 > 15",
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
			input: "5 == 5",
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
			input: "true != false",
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
			input: "true && false",
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
			input: "true && true && true",
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
			input: "true || false",
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
			input: "true || true || false",
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
			input: "true ? 10 : 5",
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
			input: "5 ? x : y ? 10 : 20",
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
