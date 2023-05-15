package lexer

import (
	"pika/pkg/lexer/token_type"
	"strings"
)

func Tokenize(line string) []token_type.Token {
	tokens := []token_type.Token{}
	src := strings.Split(line, "")

	for len(src) > 0 {
		tokenStr := src[0]
		// Check if token is skippable
		if IsSkippable(tokenStr) {
			src = src[1:]
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
		case "+", "-", "*", "/", "%":
			tokens = append(tokens, token_type.Token{Type: token_type.BinaryOperator, Value: tokenStr})
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
			src = src[1:]

			var str string
			for src[0] != "\"" {
				str += src[0]
				src = src[1:]
			}

			tokens = append(tokens, token_type.Token{Type: token_type.StringLiteral, Value: str})
			tokens = append(tokens, token_type.Token{Type: token_type.DoubleQoute, Value: tokenStr})
		case "'":
			tokens = append(tokens, token_type.Token{Type: token_type.SingleQoute, Value: tokenStr})
		default:
			tokens = append(tokens, token_type.Token{Type: token_type.Identifier, Value: tokenStr})
		}

		// Remove token
		src = src[1:]
	}
	tokens = append(tokens, token_type.Token{Type: token_type.EOF, Value: "EndOfFile"})
	return tokens
}
