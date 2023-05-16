package lexer

import (
	"errors"
	compilerErrors "pika/internal/errors"
	"pika/pkg/lexer/token_type"
	"strings"
)

func Tokenize(input string) ([]token_type.Token, error) {
	tokens := []token_type.Token{}
	src := strings.Split(input, "")

	substract := func(i int) string {
		str := src[0]
		src = src[i:]
		return str
	}

	for len(src) > 0 {
		tokenStr := src[0]
		// Check if token is skippable
		if IsSkippable(tokenStr) {
			substract(1)
			continue
		}

		// Check for keywords
		if keyword, ok := token_type.KEYWORDS[tokenStr]; ok {
			tokens = append(tokens, token_type.Token{Type: keyword, Value: tokenStr})
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
		case "+", "-", "%":
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: tokenStr})
		case "*":
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: tokenStr})
		case "/":
			nextChar := src[1]
			switch nextChar {
			case "/":
				substract(2) // consume ' // '
				for src[0] != "\n" {
					substract(1)
					if len(src) <= 0 {
						break
					}
				}
			case "*":
				substract(2) // consume /*
				for src[0]+src[1] != "*/" {
					substract(1)
					if len(src) <= 1 { // if the comment is not terminated
						return nil, errors.New(string(compilerErrors.ErrSyntaxUnterminatedMultilineComment))
					}
				}
				substract(2) // consume */
			default:
				tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: tokenStr})
			}
		case "=":
			tokens = append(tokens, token_type.Token{Type: token_type.Equals, Value: tokenStr})
		case ";":
			tokens = append(tokens, token_type.Token{Type: token_type.SemiColon, Value: tokenStr})
		case "(":
			tokens = append(tokens, token_type.Token{Type: token_type.LeftParen, Value: tokenStr})
		case ")":
			tokens = append(tokens, token_type.Token{Type: token_type.RightParen, Value: tokenStr})
		case "{":
			tokens = append(tokens, token_type.Token{Type: token_type.LeftBrace, Value: tokenStr})
		case "}":
			tokens = append(tokens, token_type.Token{Type: token_type.RightBrace, Value: tokenStr})
		case "[":
			tokens = append(tokens, token_type.Token{Type: token_type.LeftBracket, Value: tokenStr})
		case "]":
			tokens = append(tokens, token_type.Token{Type: token_type.RightBracket, Value: tokenStr})
		case ",":
			tokens = append(tokens, token_type.Token{Type: token_type.Comma, Value: tokenStr})
		case ":":
			tokens = append(tokens, token_type.Token{Type: token_type.Colon, Value: tokenStr})
		case ".":
			tokens = append(tokens, token_type.Token{Type: token_type.Dot, Value: tokenStr})
		case "\"":
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: tokenStr}) // Append double qoute
			substract(1)

			var str string
			for src[0] != "\"" {
				str += substract(1)
			}

			tokens = append(tokens, token_type.Token{Type: token_type.StringLiteral, Value: str})
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: tokenStr})
		case "'":
			tokens = append(tokens, token_type.Token{Type: token_type.SingleQoute, Value: tokenStr})
		default:
			tokens = append(tokens, token_type.Token{Type: token_type.Identifier, Value: tokenStr})
		}

		// Remove token
		if len(src) > 0 {
			substract(1)
		}
	}
	tokens = append(tokens, token_type.Token{Type: token_type.EOF, Value: "EndOfFile"})
	return tokens, nil
}
