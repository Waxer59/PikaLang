package lexer

import (
	"errors"
	compilerErrors "pika/internal/errors"
	"pika/pkg/lexer/token_type"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input          string
		expectedTokens []token_type.Token
		expectedError  error
	}{
		{
			input: "123",
			expectedTokens: []token_type.Token{
				{Type: token_type.Number, Value: "123"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedError: nil,
		},
		{
			input: "x = 5 + y",
			expectedTokens: []token_type.Token{
				{Type: token_type.Identifier, Value: "x"},
				{Type: token_type.Equals, Value: "="},
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.BinaryOperator, Value: "+"},
				{Type: token_type.Identifier, Value: "y"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedError: nil,
		},
		{
			input: "if x == 10 {\n\ty = 5\n}",
			expectedTokens: []token_type.Token{
				{Type: token_type.If, Value: "if"},
				{Type: token_type.Identifier, Value: "x"},
				{Type: token_type.EqualEqual, Value: "=="},
				{Type: token_type.Number, Value: "10"},
				{Type: token_type.LeftBrace, Value: "{"},
				{Type: token_type.Identifier, Value: "y"},
				{Type: token_type.Equals, Value: "="},
				{Type: token_type.Number, Value: "5"},
				{Type: token_type.RightBrace, Value: "}"},
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedError: nil,
		},
		{
			input: "/* This is a comment */",
			expectedTokens: []token_type.Token{
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedError: nil,
		},
		{
			input: "// This is a comment",
			expectedTokens: []token_type.Token{
				{Type: token_type.EOF, Value: "EndOfFile"},
			},
			expectedError: nil,
		},
		{
			input:          "/* Unterminated comment",
			expectedTokens: nil,
			expectedError:  errors.New(string(compilerErrors.ErrSyntaxUnterminatedMultilineComment)),
		},
	}

	for _, test := range tests {
		tokens, err := Tokenize(test.input)

		// Check error
		if !reflect.DeepEqual(err, test.expectedError) {
			t.Errorf("Expected error: %v, but got: %v", test.expectedError, err)
		}

		// Check tokens
		if !reflect.DeepEqual(tokens, test.expectedTokens) {
			t.Errorf("Expected tokens: %v, but got: %v", test.expectedTokens, tokens)
		}
	}
}
