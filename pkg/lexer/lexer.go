package lexer

//! FIX THIS LEXER

/*
112 [112 114 105 110 116 40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
40 [40 34 72 101 108 108 111 32 87 111 114 108 100 34 41] print("Hello World")
[{print 1} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {( 10} {EndOfFile 22}]
*/

import (
	"errors"
	"fmt"
	compilerErrors "pika/internal/errors"
	"pika/pkg/lexer/token_type"
)

func Tokenize(input string) ([]token_type.Token, error) {
	tokens := []token_type.Token{}
	src := []rune(input)

	pos := 0

	substract := func(i int) rune {
		if pos+i >= len(src) {
			return 0
		}

		char := src[pos]
		pos += i
		return char
	}

	for ; pos < len(src)-1; substract(1) {
		fmt.Println(src[0], src, input)
		tokenStr := src[0]
		// Check if token is skippable
		if IsSkippable(tokenStr) {
			substract(1)
			continue
		}

		// Check for number
		if IsInt(tokenStr) {
			num, rest := ExtractInt(src)
			tokens = append(tokens, token_type.Token{Type: token_type.Number, Value: num})
			src = rest
			continue
		}

		// Check for alpha
		if IsAlpha(tokenStr) {
			alpha, rest := ExtractAlpha(src)
			alphaType := token_type.Identifier

			if keyword, ok := IsKeyword(alpha); ok {
				alphaType = keyword
			}

			tokens = append(tokens, token_type.Token{Type: alphaType, Value: alpha})

			src = rest
			continue
		}

		// Check for operators
		switch tokenStr {
		case '+', '-', '%':
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenStr)})
		case '*':
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenStr)})
		case '/':
			nextChar := src[1]
			switch nextChar {
			case '/':
				substract(2) // consume ' // '
				for src[0] != '\n' {
					substract(1)
					if len(src) <= 0 {
						break
					}
				}
			case '*':
				substract(2) // consume /*
				for src[0] != '*' && src[1] == '/' {
					substract(1)
					if len(src) <= 1 { // if the comment is not terminated
						return nil, errors.New(string(compilerErrors.ErrSyntaxUnterminatedMultilineComment))
					}
				}
				substract(2) // consume */
			default:
				tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: string(tokenStr)})
			}
		case '=':
			tokens = append(tokens, token_type.Token{Type: token_type.Equals, Value: string(tokenStr)})
		case ';':
			tokens = append(tokens, token_type.Token{Type: token_type.SemiColon, Value: string(tokenStr)})
		case '(':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftParen, Value: string(tokenStr)})
		case ')':
			tokens = append(tokens, token_type.Token{Type: token_type.RightParen, Value: string(tokenStr)})
		case '{':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftBrace, Value: string(tokenStr)})
		case '}':
			tokens = append(tokens, token_type.Token{Type: token_type.RightBrace, Value: string(tokenStr)})
		case '[':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftBracket, Value: string(tokenStr)})
		case ']':
			tokens = append(tokens, token_type.Token{Type: token_type.RightBracket, Value: string(tokenStr)})
		case ',':
			tokens = append(tokens, token_type.Token{Type: token_type.Comma, Value: string(tokenStr)})
		case ':':
			tokens = append(tokens, token_type.Token{Type: token_type.Colon, Value: string(tokenStr)})
		case '.':
			tokens = append(tokens, token_type.Token{Type: token_type.Dot, Value: string(tokenStr)})
		case '"':
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: string(tokenStr)}) // Append double qoute
			substract(1)

			var str string
			for src[0] != '\'' {
				str += string(substract(1))
			}

			tokens = append(tokens, token_type.Token{Type: token_type.StringLiteral, Value: str})
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: string(tokenStr)})
		case '\'':
			tokens = append(tokens, token_type.Token{Type: token_type.SingleQoute, Value: string(tokenStr)})
		default:
			tokens = append(tokens, token_type.Token{Type: token_type.Identifier, Value: string(tokenStr)})
		}
	}
	tokens = append(tokens, token_type.Token{Type: token_type.EOF, Value: "EndOfFile"})
	return tokens, nil
}
